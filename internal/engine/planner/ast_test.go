// Copyright 2021-2023 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package planner

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"testing"

	"github.com/ghodss/yaml"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"google.golang.org/protobuf/testing/protocmp"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	"github.com/cerbos/cerbos/internal/conditions"
	"github.com/cerbos/cerbos/internal/util"
	"google.golang.org/protobuf/encoding/protojson"
)

//go:embed testdata/ast_build_expr.yaml
var astBuildExprBlob []byte

type (
	exOp = enginev1.PlanResourcesFilter_Expression_Operand
)

func getExpectedExpressions(t *testing.T) map[string]*exOp {
	t.Helper()

	var raw map[string]json.RawMessage
	err := yaml.Unmarshal(astBuildExprBlob, &raw)
	require.NoError(t, err)
	res := make(map[string]*exOp, len(raw))
	for k, v := range raw {
		expected := new(exOp)
		b, err := v.MarshalJSON()
		require.NoError(t, err)
		err = util.ReadJSONOrYAML(bytes.NewReader(b), expected)
		require.NoError(t, err, string(b))
		res[k] = expected
	}

	return res
}

func Test_buildExpr(t *testing.T) {
	env, err := conditions.StdPartialEnv.Extend(cel.Declarations(
		decls.NewVar("x", decls.NewListType(decls.Dyn)),
		decls.NewVar("a", decls.Dyn),
		decls.NewVar("b", decls.Dyn),
		decls.NewVar("c", decls.Dyn),
		decls.NewFunction("f", &exprpb.Decl_FunctionDecl_Overload{
			Params:             []*exprpb.Type{decls.Dyn, decls.Dyn, decls.Int},
			ResultType:         decls.Dyn,
			IsInstanceFunction: true,
		}, &exprpb.Decl_FunctionDecl_Overload{
			Params:     []*exprpb.Type{decls.Dyn, decls.Int},
			ResultType: decls.Dyn,
		}),
	))
	require.NoError(t, err)

	compile := func(t *testing.T, s string) *exprpb.Expr {
		ast, iss := env.Compile(s)
		require.Nil(t, iss, iss.Err())
		return ast.Expr()
	}

	for k, v := range getExpectedExpressions(t) {
		name := k
		want := v
		t.Run(name, func(t *testing.T) {
			is := require.New(t)
			acc := new(exOp)
			err := buildExpr(compile(t, name), acc)
			is.NoError(err)

			is.Empty(cmp.Diff(want, acc, protocmp.Transform()), "unexpected expression: %s", protojson.Format(acc))
		})
	}
}
