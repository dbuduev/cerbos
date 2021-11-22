package engine

import (
	"fmt"
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	requestv1 "github.com/cerbos/cerbos/api/genpb/cerbos/request/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/conditions"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/parser"
	"github.com/stretchr/testify/require"
	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

func Test_evaluateCondition(t *testing.T) {
	type args struct {
		expr      string
		condition *runtimev1.Condition
		input     *requestv1.ListResourcesRequest
	}

	unparse := func(t *testing.T, expr *expr.CheckedExpr) string {
		t.Helper()
		require.NotNil(t, expr)
		source, err := parser.Unparse(expr.Expr, expr.SourceInfo)
		require.NoError(t, err)
		return source
	}

	compile := func(expr string, input *requestv1.ListResourcesRequest) args {
		ast, iss := conditions.StdEnv.Compile(expr)
		require.Nil(t, iss, "Error is %s", iss.Err())
		checkedExpr, err := cel.AstToCheckedExpr(ast)
		require.NoError(t, err)
		c := &runtimev1.Condition{Op: &runtimev1.Condition_Expr{Expr: &runtimev1.Expr{
			Original: expr,
			Checked:  checkedExpr,
		}}}
		return args{
			expr:      expr,
			condition: c,
			input:     input,
		}
	}
	tests := []struct {
		args           args
		wantExpression string
	}{
		{
			args:           compile("false", &requestv1.ListResourcesRequest{}),
			wantExpression: "false",
		},
		{
			args: compile("P.attr.authenticated", &requestv1.ListResourcesRequest{
				Principal: &enginev1.Principal{
					Attr: map[string]*structpb.Value{"authenticated": {Kind: &structpb.Value_BoolValue{BoolValue: true}}},
				},
			}),
			wantExpression: "true",
		},
		{
			args: compile("R.attr.owner == P.attr.name", &requestv1.ListResourcesRequest{
				Principal: &enginev1.Principal{
					Attr: map[string]*structpb.Value{"name": {Kind: &structpb.Value_StringValue{StringValue: "harry"}}},
				},
			}),
			wantExpression: `R.attr.owner == "harry"`,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Expr:%q", tt.args.expr), func(t *testing.T) {
			is := require.New(t)
			got, err := evaluateCondition(tt.args.condition, tt.args.input)
			is.NoError(err)
			expression := got.GetExpression()
			is.Equal(tt.wantExpression, unparse(t, expression))
		})
	}
	for _, op := range []responsev1.ListResourcesResponse_LogicalOperation_Operator {responsev1.ListResourcesResponse_LogicalOperation_AND, responsev1.ListResourcesResponse_LogicalOperation_OR} {
		attr := make(map[string]*structpb.Value)
		conds := make([]*runtimev1.Condition, len(tests))

		exprList := &runtimev1.Condition_ExprList{}
		var c *runtimev1.Condition
		if op == responsev1.ListResourcesResponse_LogicalOperation_AND {
            c = &runtimev1.Condition{Op: &runtimev1.Condition_All{All: exprList}}
        } else {
            c = &runtimev1.Condition{Op: &runtimev1.Condition_Any{Any: exprList}}
        }
		t.Run(fmt.Sprintf("%s operation", responsev1.ListResourcesResponse_LogicalOperation_Operator_name[int32(op)]), func(t *testing.T) {
			is := require.New(t)
			for i := 0; i < len(tests); i++ {
				exprList.Expr = append(exprList.Expr, tests[i].args.condition)
				conds[i] = tests[i].args.condition
				input := tests[i].args.input
				if input.Principal != nil && input.Principal.Attr != nil {
					for k, v := range input.Principal.Attr {
						if _, ok := attr[k]; ok {
							t.Fatalf("Duplicate key %q", k)
						}
						attr[k] = v
					}
				}
			}
			got, err := evaluateCondition(c, &requestv1.ListResourcesRequest{Principal: &enginev1.Principal{Attr: attr}})
			is.NotNil(got)
			is.NoError(err)
			operation := got.GetLogicalOperation()
			is.NotNil(operation)
			is.Equal(op, operation.Operator)
			for i := 0; i < len(tests); i++ {
				expression := operation.Nodes[i].GetExpression()
				is.Equal(tests[i].wantExpression, unparse(t, expression))
			}
		})
	}
}