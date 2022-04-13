package main

import (
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"fmt"
	"golang.org/x/exp/slices"
	"io"
)

func renderExpr(w io.Writer, e *exprpb.Expr) {
	f0 := func(format string, a ...any) {
		fmt.Fprintf(w, format, a...)
	}
	//	*Expr_ConstExpr
	//	*Expr_IdentExpr
	//	*Expr_SelectExpr
	//	*Expr_CallExpr
	switch expr := e.ExprKind.(type) {
	case *exprpb.Expr_CallExpr:
		if !isBinaryOp(expr) {
			panic("only binary operator implemented")
		}
		e := expr.CallExpr
		renderExpr(w, e.Args[0])
		f0(" == ")
		renderExpr(w, e.Args[1])
	case *exprpb.Expr_SelectExpr:
		renderExpr(w, expr.SelectExpr.Operand)
		f0(".%s", expr.SelectExpr.Field)
	case *exprpb.Expr_IdentExpr:
		f0(expr.IdentExpr.Name)
	case *exprpb.Expr_ConstExpr:
		renderConstExpr(w, expr.ConstExpr)
	default:
		panic(fmt.Errorf("unsupported expr kind: %T", expr))
	}
}

var binaryOps = []string{"_==_"}
var mapOps = map[string]string{
	"_==_": "==",
}

func isBinaryOp(e *exprpb.Expr_CallExpr) bool {
	if e != nil && e.CallExpr != nil {
		c := e.CallExpr
		return c.Target == nil && len(c.Args) == 2 && slices.Contains(binaryOps, c.Function)
	}
	return false
}

func renderConstExpr(w io.Writer, c *exprpb.Constant) {
	var a any
	switch v := c.ConstantKind.(type) {
	case *exprpb.Constant_BoolValue:
		a = v.BoolValue
	case *exprpb.Constant_DoubleValue:
		a = v.DoubleValue
	case *exprpb.Constant_Int64Value:
		a = v.Int64Value
	case *exprpb.Constant_NullValue:
		fmt.Fprintf(w, "null")
		return
	case *exprpb.Constant_StringValue:
		a = v.StringValue
	case *exprpb.Constant_Uint64Value:
		a = v.Uint64Value
	//case *exprpb.Constant_BytesValue:
	default:
		panic(fmt.Errorf("unsupported constant: %v", c))
	}
	fmt.Fprintf(w, "%#v", a)
}
