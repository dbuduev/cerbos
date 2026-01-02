// Copyright 2021-2025 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

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

	// Convert to absolute path if relative
	if !filepath.IsAbs(benchDir) {
		absDir, err := filepath.Abs(benchDir)
		if err != nil {
			b.Fatalf("failed to resolve absolute path: %v", err)
		}
		benchDir = absDir
	}

	ctx := context.Background()
	fsys := os.DirFS(benchDir)

	setupStart := time.Now()

	// Build index from the directory
	idx, err := index.Build(ctx, fsys, index.WithBuildFailureLogLevel(zap.DebugLevel))
	if err != nil {
		b.Fatalf("failed to build index: %v", err)
	}

	// Create disk store from index
	store := disk.NewFromIndexWithConf(idx, &disk.Conf{})

	// Create compiler manager
	compiler, err := compile.NewManager(ctx, store)
	if err != nil {
		b.Fatalf("failed to create compiler manager: %v", err)
	}

	// Create rule table
	ruleTable, err := ruletable.NewRuleTableFromLoader(ctx, compiler)
	if err != nil {
		b.Fatalf("failed to create rule table: %v", err)
	}

	// Create schema manager
	schemaMgr := schema.NewFromConf(ctx, store, schema.NewConf(schema.EnforcementReject))

	// Create rule table manager
	ruletableMgr, err := ruletable.NewRuleTableManager(ruleTable, compiler, schemaMgr)
	if err != nil {
		b.Fatalf("failed to create ruletable manager: %v", err)
	}

	// Create engine
	eng := engine.NewEphemeral(nil, ruletableMgr, schemaMgr)

	b.Logf("Setup took %s", time.Since(setupStart))

	for _, trace := range []bool{false, true} {
		b.Run(traceName(trace), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := Verify(ctx, fsys, eng, Config{Trace: trace})
				if err != nil {
					b.Fatalf("verify failed: %v", err)
				}
			}
		})
	}
}

func traceName(trace bool) string {
	if trace {
		return "Trace"
	}
	return "NoTrace"
}
