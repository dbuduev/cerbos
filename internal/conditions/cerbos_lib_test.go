// Copyright 2021 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package conditions_test

import (
	"fmt"
	"github.com/google/cel-go/common"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/parser"
	"google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"log"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/ghodss/yaml"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/ext"
	"github.com/stretchr/testify/require"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	"github.com/cerbos/cerbos/internal/conditions"
)

func TestCerbosLib(t *testing.T) {
	testCases := []struct {
		expr    string
		wantErr bool
	}{
		{expr: `"192.168.0.5".inIPAddrRange("192.168.0.0/24") == true`},
		{expr: `"192.169.1.5".inIPAddrRange("192.168.0.0/24") == false`},
		{expr: `"test".inIPAddrRange("192.168.0.0/24") == false`, wantErr: true},
		{expr: `"2001:0db8:0000:0000:0000:0000:1000:0000".inIPAddrRange("2001:db8::/48") == true`},
		{expr: `"3001:0fff:0000:0000:0000:0000:0000:0000".inIPAddrRange("2001:db8::/48") == false`},
		{expr: `timestamp("2021-05-01T00:00:00.000Z").timeSince() > duration("1h")`},
		{expr: `has_intersection([1,2,3],[3,5])`},
		{expr: `has_intersection([1,2,3],[4,5]) == false`},
		{expr: `has_intersection(['1','2','3'],['3','5'])`},
		{expr: `intersect([1,2,3],[2,3,5]) == [2,3]`},
		{expr: `intersect([1,2,3],[4,5]) == []`},
		{expr: `intersect(['1','2','3'],['3','5']) == ['3']`},
		{expr: `[1,2].is_subset([2,3]) == false`},
		{expr: `[[1],[2]].is_subset([[2],[3]]) == false`},
		{expr: `[1,2].is_subset([1,2])`},
		{expr: `[1,2].is_subset([1,2,3])`},
		{expr: `[[1],[2]].is_subset([[1],[2],[3]])`},
		{expr: `["1","2"].is_subset(["1","2","3"])`},
		{expr: `[1,1].is_subset([1])`},
		{expr: `[].is_subset([1])`},
		{expr: `[].except([1]) == []`},
		{expr: `[1].except([]) == [1]`},
		{expr: `[].except([]) == []`},
		{expr: `[1].except([1]) == []`},
		{expr: `[1].except([1,2,3]) == []`},
		{expr: `[1,3,5].except([2,4]) == [1,3,5]`},
		{expr: `[1,3,5].except([5,3]) == [1]`},
		{expr: `[1,2,3] + [3,5] == [1,2,3,3,5]`},
		{expr: `hierarchy("a.b.c.d") == hierarchy("a.b.c.d")`},
		{expr: `hierarchy("a.b.c.d") != hierarchy("a.b.c.d.e")`},
		{expr: `hierarchy("a:b:c:d", ":") == hierarchy("a.b.c.d")`},
		{expr: `hierarchy("aFOObFOOcFOOd", "FOO") == hierarchy("a.b.c.d")`},
		{expr: `hierarchy(["a","b","c","d"]) == hierarchy("a.b.c.d")`},
		{expr: `hierarchy("a.b.c.d").size() == 4`},
		{expr: `hierarchy("a.b.c.d")[2] == "c"`},
		{expr: `hierarchy("a.b").ancestorOf(hierarchy("a.b.c.d.e"))`},
		{expr: `hierarchy("a.b").ancestorOf(hierarchy("a.b")) == false`},
		{expr: `hierarchy("a.b.c.d").ancestorOf(hierarchy("a.b.c.d.e"))`},
		{expr: `hierarchy("x.y.c.d").ancestorOf(hierarchy("a.b.c.d.e")) == false`},
		{expr: `hierarchy("a.b.c.d").commonAncestors(hierarchy("a.b.c.d.e")) == hierarchy("a.b.c.d")`},
		{expr: `hierarchy("a.b.c.d").commonAncestors(hierarchy("a.b.c.d")) == hierarchy("a.b.c")`},
		{expr: `hierarchy("a.b.c.d").commonAncestors(hierarchy("x.y.z")).size() == 0`},
		{expr: `hierarchy("a.b.c.d.e").descendentOf(hierarchy("a.b"))`},
		{expr: `hierarchy("a.b").descendentOf(hierarchy("a.b")) == false`},
		{expr: `hierarchy("x.b").descendentOf(hierarchy("a.b")) == false`},
		{expr: `hierarchy("a.b.c.d.e").immediateChildOf(hierarchy("a.b.c.d"))`},
		{expr: `hierarchy("a.b.c.d.e").immediateChildOf(hierarchy("a.b.c")) == false`},
		{expr: `hierarchy("a.b.c.d").immediateParentOf(hierarchy("a.b.c.d.e"))`},
		{expr: `!hierarchy("a.b.c").immediateParentOf(hierarchy("a.b.c.d.e"))`},
		{expr: `hierarchy("a.b.c").siblingOf(hierarchy("a.b.d"))`},
		{expr: `hierarchy("a.b.c").siblingOf(hierarchy("x.b.d")) == false`},
		{expr: `hierarchy("a.b.c.d").siblingOf(hierarchy("a.b.d")) == false`},
		{expr: `hierarchy("a.b.c.d").overlaps(hierarchy("a.b.c.d")) == true`},
		{expr: `hierarchy("a.b.c.d").overlaps(hierarchy("a.b.c")) == true`},
		{expr: `hierarchy("a.b").overlaps(hierarchy("a.b.c.d")) == true`},
		{expr: `hierarchy("a.b.x").overlaps(hierarchy("a.b.c.d")) == false`},
	}
	env, err := cel.NewEnv(conditions.CerbosCELLib())
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.expr, func(t *testing.T) {
			is := require.New(t)
			ast, issues := env.Compile(tc.expr)
			is.NoError(issues.Err())

			prg, err := env.Program(ast)
			is.NoError(err)

			have, _, err := prg.Eval(cel.NoVars())
			if tc.wantErr {
				is.Error(err)
			} else {
				is.NoError(err)
				is.Equal(true, have.Value())
			}
		})
	}
}

func prepareProgram(tb testing.TB, expr string) cel.Program {
	tb.Helper()
	is := require.New(tb)
	env, err := cel.NewEnv(conditions.CerbosCELLib())
	is.NoError(err)
	ast, issues := env.Compile(expr)
	is.NoError(issues.Err())

	prg, err := env.Program(ast)
	is.NoError(err)
	return prg
}

func generateExpr(size int) string {
	lhs := make([]string, size)
	for i := 0; i < size; i++ {
		lhs[i] = fmt.Sprintf("'%05d'", i)
	}
	rhs := make([]string, size)
	copy(rhs, lhs)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(rhs), func(i, j int) { rhs[i], rhs[j] = rhs[j], rhs[i] })
	return fmt.Sprintf("intersect([%s], [%s])", strings.Join(lhs, ","), strings.Join(rhs, ","))
}

func BenchmarkIntersect50(b *testing.B) {
	benchmarkIntersect(b, 50)
}

func BenchmarkIntersect25(b *testing.B) {
	benchmarkIntersect(b, 25)
}

func BenchmarkIntersect15(b *testing.B) {
	benchmarkIntersect(b, 15)
}

func BenchmarkIntersect5(b *testing.B) {
	benchmarkIntersect(b, 5)
}

func benchmarkIntersect(b *testing.B, size int) {
	b.Helper()
	expr := generateExpr(size)
	prg := prepareProgram(b, expr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, err := prg.Eval(cel.NoVars())
		require.NoError(b, err)
	}
}

func TestPartialEvaluationWithGlobalVars(t *testing.T) {
	expander := func(eh parser.ExprHelper, _ *expr.Expr, _ []*expr.Expr) (*expr.Expr, *common.Error) {
		return eh.Select(eh.Select(eh.Ident("R"), "attr"), "geo"), nil
	}
	geo := parser.NewReceiverMacro("geo", 0, expander)
	env, _ := cel.NewEnv(
		cel.Types(&enginev1.Resource{}),
		cel.Declarations(
			decls.NewVar("y", decls.NewListType(decls.String)),
			decls.NewVar(conditions.CELResourceAbbrev, decls.NewObjectType("cerbos.engine.v1.Resource")),
			decls.NewVar("z", decls.String),
			decls.NewVar("v", decls.NewMapType(decls.String, decls.Dyn))),
		ext.Strings(),
		cel.Macros(geo))

	vars, _ := cel.PartialVars(map[string]interface{}{
		"y": []string{"GB", "US"},
		"z": "ca",
		"v": map[string]interface{}{},
	}, cel.AttributePattern("R"))
	expr := "v.geo in (y + [z]).map(t, t.upperAscii())"
	want := `R.attr.geo in ["GB", "US", "CA"]`

	is := require.New(t)
	ast, issues := env.Compile(expr)
	if issues != nil {
		is.NoError(issues.Err())
	}
	prg, err := env.Program(ast, cel.EvalOptions(cel.OptTrackState, cel.OptPartialEval))
	is.NoError(err)

	out, det, err := prg.Eval(vars)
	is.NoError(err)
	is.True(types.IsUnknown(out))
	residual, err := env.ResidualAst(ast, det)
	astToString, err := cel.AstToString(residual)
	is.NoError(err)
	is.Equal(want, astToString)
}

func TestPartialEvaluation(t *testing.T) {
	env, _ := cel.NewEnv(
		cel.Types(&enginev1.Resource{}),
		cel.Declarations(
			decls.NewVar("y", decls.NewListType(decls.String)),
			decls.NewVar(conditions.CELResourceAbbrev, decls.NewObjectType("cerbos.engine.v1.Resource")),
			decls.NewVar("z", decls.String)), ext.Strings())

	vars, _ := cel.PartialVars(map[string]interface{}{
		"y": []string{"GB", "US"},
		"z": "ca",
	}, cel.AttributePattern("R"))
	tests := []struct {
		expr, result string
	}{
		{
			expr:   "R.attr.geo in (y + [z]).map(t, t.upperAscii())",
			result: `R.attr.geo in ["GB", "US", "CA"]`,
		},
		{
			expr:   "R.attr.geo in (y + [z]).map(t, t.upperAscii()) || 1 == 1",
			result: "true",
		},
		{
			expr:   `"CA" in (y + [z]).map(t, t.upperAscii())`,
			result: "true",
		},
		{
			expr:   `("CA" in (y + [z]).map(t, t.upperAscii())) && 1 == 2`,
			result: "false",
		},
		{
			expr:   `("NZ" in (y + [z]).map(t, t.upperAscii()))`,
			result: "false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.expr, func(t *testing.T) {
			is := require.New(t)
			expr := tt.expr
			ast, issues := env.Compile(expr)
			if issues != nil {
				is.NoError(issues.Err())
			}
			prg, err := env.Program(ast, cel.EvalOptions(cel.OptTrackState, cel.OptPartialEval))
			is.NoError(err)

			out, det, err := prg.Eval(vars)
			is.NotNil(det) // It is not nil if cel.OptTrackState is included in the cel.EvalOptions
			t.Log(out.Type())
			is.NoError(err)
			residual, err := env.ResidualAst(ast, det)
			is.NoError(err)
			bytes, err := yaml.Marshal(residual.Expr())
			log.Print("\n", string(bytes))
			is.NoError(err)
			astToString, err := cel.AstToString(residual)
			is.NoError(err)
			is.Equal(tt.result, astToString)
			log.Print(astToString)
		})
	}
}
