// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"go.uber.org/zap"

	"github.com/cerbos/cerbos/internal/compile"
	"github.com/cerbos/cerbos/internal/engine"
	"github.com/cerbos/cerbos/internal/ruletable"
	"github.com/cerbos/cerbos/internal/schema"
	"github.com/cerbos/cerbos/internal/storage/disk"
	"github.com/cerbos/cerbos/internal/storage/index"
)

// BenchmarkVerify benchmarks the Verify function against policies and tests in a directory.
// Usage: CERBOS_BENCH_DIR=/path/to/policies go test -bench=BenchmarkVerify -run='^$' ./internal/verify/
func BenchmarkVerify(b *testing.B) {
	benchDir := os.Getenv("CERBOS_BENCH_DIR")
	if benchDir == "" {
		b.Skip("CERBOS_BENCH_DIR environment variable not set")
	}

	if !filepath.IsAbs(benchDir) {
		absDir, err := filepath.Abs(benchDir)
		if err != nil {
			b.Fatalf("failed to resolve absolute path: %v", err)
		}
		benchDir = absDir
	}

	ctx := context.Background()
	fsys := os.DirFS(benchDir)

	idx, err := index.Build(ctx, fsys, index.WithBuildFailureLogLevel(zap.DebugLevel))
	if err != nil {
		b.Fatalf("failed to build index: %v", err)
	}

	store := disk.NewFromIndexWithConf(idx, &disk.Conf{})

	compiler, err := compile.NewManager(ctx, store)
	if err != nil {
		b.Fatalf("failed to create compiler manager: %v", err)
	}

	ruleTable, err := ruletable.NewRuleTableFromLoader(ctx, compiler)
	if err != nil {
		b.Fatalf("failed to create rule table: %v", err)
	}

	schemaMgr := schema.NewFromConf(ctx, store, schema.NewConf(schema.EnforcementReject))

	ruletableMgr, err := ruletable.NewRuleTableManager(ruleTable, compiler, schemaMgr)
	if err != nil {
		b.Fatalf("failed to create ruletable manager: %v", err)
	}

	eng := engine.NewEphemeral(nil, ruletableMgr, schemaMgr)

	for _, batching := range []bool{false, true} {
		for _, trace := range []bool{true} {
			name := modeName(trace, batching)
			b.Run(name, func(b *testing.B) {
				b.ReportAllocs()
				for b.Loop() {
					_, err := Verify(ctx, fsys, eng, Config{Trace: trace, Batching: batching})
					if err != nil {
						b.Fatalf("verify failed: %v", err)
					}
				}
			})
		}
	}
}

func modeName(trace, batching bool) string {
	name := "NoTrace"
	if trace {
		name = "Trace"
	}
	if batching {
		name += "_Batched"
	}
	return name
}
