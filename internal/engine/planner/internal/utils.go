// Copyright 2021-2023 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package internal

import exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"

func WalkExpr(e *exprpb.Expr, pre func(e *exprpb.Expr)) {
	if e == nil {
		return
	}
	pre(e)
	switch e := e.ExprKind.(type) {
	case *exprpb.Expr_SelectExpr:
		WalkExpr(e.SelectExpr.Operand, pre)
	case *exprpb.Expr_CallExpr:
		WalkExpr(e.CallExpr.Target, pre)
		for _, arg := range e.CallExpr.Args {
			WalkExpr(arg, pre)
		}
	case *exprpb.Expr_StructExpr:
		for _, entry := range e.StructExpr.Entries {
			WalkExpr(entry.GetMapKey(), pre)
			WalkExpr(entry.GetValue(), pre)
		}
	case *exprpb.Expr_ComprehensionExpr:
		ce := e.ComprehensionExpr
		WalkExpr(ce.IterRange, pre)
		WalkExpr(ce.AccuInit, pre)
		WalkExpr(ce.LoopStep, pre)
		WalkExpr(ce.LoopCondition, pre)
		WalkExpr(ce.Result, pre)
	case *exprpb.Expr_ListExpr:
		for _, element := range e.ListExpr.Elements {
			WalkExpr(element, pre)
		}
	}
}

func UpdateIDs(e *exprpb.Expr) int64 {
	var n int64
	ids := make(map[*exprpb.Expr]int64)

	WalkExpr(e, func(e *exprpb.Expr) {
		if id, ok := ids[e]; ok {
			e.Id = id
		} else {
			n++
			ids[e] = n
			e.Id = n
		}
	})

	return n
}

func GetMaxID(e *exprpb.Expr) int64 {
	if e == nil {
		return 0
	}
	n := e.Id
	WalkExpr(e, func(e *exprpb.Expr) {
		if e.Id > n {
			n = e.Id
		}
	})

	return n
}

func MkCallExpr(op string, args ...*exprpb.Expr) *exprpb.Expr {
	e := &exprpb.Expr{
		ExprKind: &exprpb.Expr_CallExpr{CallExpr: &exprpb.Expr_Call{
			Function: op,
			Args:     args,
		}},
	}
	return e
}
