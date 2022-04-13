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
	"fmt"
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

func mkCompiler(t *testing.T) (*compile.Manager, schema.Manager) {
	t.Helper()

	dir := pathToDir(t, "store")

	ctx, cancelFunc := context.WithCancel(context.Background())
	t.Cleanup(cancelFunc)

	store, err := disk.NewStore(ctx, &disk.Conf{Directory: dir})
	require.NoError(t, err)

	schemaConf := schema.NewConf(schema.EnforcementReject)
	schemaMgr := schema.NewFromConf(ctx, store, schemaConf)

	return compile.NewManagerFromDefaultConf(ctx, store, schemaMgr), schemaMgr
}

func TestNewCompiler(t *testing.T) {
	is := require.New(t)

	compiler, _ := mkCompiler(t)
	is.NotNil(compiler)
}

func TestGetPolicy(t *testing.T) {
	resource, policyVer, scope := "leave_request", "staging", ""
	resourceModID := namer.ResourcePolicyModuleID(resource, policyVer, scope)
	compiler, schMgr := mkCompiler(t)
	ctx := context.Background()
	rps, err := compiler.Get(ctx, resourceModID)
	is := require.New(t)
	is.NoError(err)
	is.NotNil(rps)
	rp := rps.GetResourcePolicy()
	is.NotNil(rp)

	sb := new(strings.Builder)
	fmt.Fprintln(sb, baseClasses)
	schema, err := schMgr.LoadSchema(ctx, rp.Schemas.ResourceSchema.Ref)
	is.NoError(err)
	is.NotNil(schema)
	SaveAsClass(sb, "resourceAttr", schema)

	schema, err = schMgr.LoadSchema(ctx, rp.Schemas.PrincipalSchema.Ref)
	is.NoError(err)
	is.NotNil(schema)
	SaveAsClass(sb, "principalAttr", schema)

	Save(sb, rps.GetResourcePolicy())
	t.Log(sb.String())
}
