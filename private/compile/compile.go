// Copyright 2021-2022 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package compile

import (
	"context"
	"errors"
	"fmt"
	"io/fs"

	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	internalcompile "github.com/cerbos/cerbos/internal/compile"
	"github.com/cerbos/cerbos/internal/observability/logging"
	"github.com/cerbos/cerbos/internal/schema"
	"github.com/cerbos/cerbos/internal/storage/disk"
	"github.com/cerbos/cerbos/internal/storage/index"
	"go.uber.org/zap"
)

type Artefact struct {
	Error      error
	PolicySet  *runtimev1.RunnablePolicySet
	SourceFile string
}

type Errors struct {
	*runtimev1.Errors
}

func (e *Errors) Error() string {
	switch e.Kind.(type) {
	case *runtimev1.Errors_IndexBuildErrors:
		return "index build failed"
	case *runtimev1.Errors_CompileErrors:
		return "compilation failed"
	default:
		return fmt.Sprintf("unhandled error kind %T", e.Kind)
	}
}

func Files(ctx context.Context, fsys fs.FS) (<-chan Artefact, error) {
	idx, err := index.Build(ctx, fsys)
	if err != nil {
		idxErrs := new(index.BuildError)
		if errors.As(err, &idxErrs) {
			return nil, &Errors{
				Errors: &runtimev1.Errors{
					Kind: &runtimev1.Errors_IndexBuildErrors{IndexBuildErrors: idxErrs.IndexBuildErrors},
				},
			}
		}

		return nil, fmt.Errorf("failed to build index: %w", err)
	}

	outChan := make(chan Artefact, 1)

	go func() {
		defer close(outChan)

		store := disk.NewFromIndexWithConf(idx, &disk.Conf{})
		schemaMgr := schema.NewFromConf(ctx, store, schema.NewConf(schema.EnforcementReject))
		logger := logging.FromContext(ctx).Named("compile")

		inChan := idx.GetAllCompilationUnits(ctx)
		for unit := range inChan {
			srcFile := unit.MainSourceFile()
			log := logger.With(zap.String("source", srcFile))
			log.Debug("Compiling unit")

			artefact := Artefact{SourceFile: srcFile}
			artefact.PolicySet, artefact.Error = internalcompile.Compile(unit, schemaMgr)

			if artefact.Error != nil {
				log.Error("Compilation failed", zap.Error(artefact.Error))
				compErrs := new(internalcompile.ErrorList)
				if errors.As(artefact.Error, &compErrs) {
					artefact.Error = &Errors{
						Errors: &runtimev1.Errors{
							Kind: &runtimev1.Errors_CompileErrors{CompileErrors: compErrs.CompileErrors},
						},
					}
				}
			} else {
				log.Debug("Compilation succeeded")
			}

			log.Debug("Sending artefact")
			select {
			case outChan <- artefact:
				log.Debug("Artefact sent")
			case <-ctx.Done():
				log.Debug("Artefact send cancelled", zap.Error(ctx.Err()))
				return
			}
		}
	}()

	return outChan, nil
}
