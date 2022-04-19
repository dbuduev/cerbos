package main

import (
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/stretchr/testify/require"
	"strings"
	"fmt"
)

func main() {
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
