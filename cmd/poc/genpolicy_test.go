package main

import (
	"github.com/cerbos/cerbos/internal/compile"
	"testing"
	"github.com/cerbos/cerbos/internal/storage/disk"
	"github.com/stretchr/testify/require"
	"github.com/cerbos/cerbos/internal/schema"
	"context"
	"path/filepath"
	"runtime"
	"github.com/cerbos/cerbos/internal/namer"
	"strings"
)

func pathToDir(tb testing.TB, dir string) string {
	tb.Helper()

	_, currFile, _, ok := runtime.Caller(0)
	if !ok {
		tb.Error("Failed to detect testdata directory")
		return ""
	}

	return filepath.Join(filepath.Dir(currFile), "testdata", dir)
}

func mkCompiler(t *testing.T) *compile.Manager {
	t.Helper()

	dir := pathToDir(t, "store")

	ctx, cancelFunc := context.WithCancel(context.Background())
	t.Cleanup(cancelFunc)

	store, err := disk.NewStore(ctx, &disk.Conf{Directory: dir})
	require.NoError(t, err)

	schemaConf := schema.NewConf(schema.EnforcementNone)
	schemaMgr := schema.NewFromConf(ctx, store, schemaConf)

	return compile.NewManagerFromDefaultConf(ctx, store, schemaMgr)
}

func TestNewCompiler(t *testing.T) {
	is := require.New(t)

	is.NotNil(mkCompiler(t))
}

func TestGetPolicy(t *testing.T) {
	resource, policyVer, scope := "leave_request", "staging", ""
	resourceModID := namer.ResourcePolicyModuleID(resource, policyVer, scope)
	compiler := mkCompiler(t)
	ctx := context.Background()
	rps, err := compiler.Get(ctx, resourceModID)
	is := require.New(t)
	is.NoError(err)
	is.NotNil(rps)
	is.NotNil(rps.GetResourcePolicy())

	sb := new(strings.Builder)
	Save(sb, rps.GetResourcePolicy())
	t.Log(sb.String())
}
