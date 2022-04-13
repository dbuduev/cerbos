package main

import (
	"io"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"fmt"
	"golang.org/x/exp/maps"
)

const (
	rolesMatched = `
function rolesMatched(a: Array<string>, b: Array<string>): boolean {
  for (let i = 0; i < a.length; i++) {
    for (let j = 0; j < b.length; j++) {
      if (a[i] == b[j]) {
        return true
      }
    }
  }
  return false
}
`
	actionMatched = `
function matchGlob(a: string, g: string): boolean {
  const ps = g.split(":");
  const as = a.split(":");
  if (ps.length != as.length) {
    return false;
  }
  for (let i = 0; i < ps.length; i++) {
    if (ps[i] != "*" && ps[i] != as[i]) {
      return false;
    }
  }
  return true;
}

export function actionMatched(a: string, globs: Array<string>): boolean {
  let result = false;

  for (let i = 0; i < globs.length && !result; i++) {
    result = matchGlob(a, globs[i])
  }

  return result;
}
`
)

func Save(w io.Writer, rps *runtimev1.RunnableResourcePolicySet) {
	var err error
	f0 := func(format string, a ...any) { // no indentation
		if err == nil {
			fmt.Fprintf(w, format, a...)
		}
	}
	indent := 0
	f := func(format string, a ...any) {
		if err == nil {
			for i := 0; i < indent; i++ {
				_, err = fmt.Fprintf(w, "  ")
			}
		}
		if err == nil {
			_, err = fmt.Fprintf(w, format, a...)
		}
	}
	f(rolesMatched)
	f(actionMatched)
	f("\nexport function check(request: Request): string {")
	f(`
  let actions : Array<string>;
  let roles : Array<string>;

  const P = request.principal;
  const R = request.resource;`)
	f("\n")
	indent = 1
	for _, r := range rps.Policies[0].Rules {
		// roles
		f("roles = [")
		csv(w, "%q", maps.Keys(r.Roles))
		f0("];\n")

		f("if(rolesMatched(P.roles, roles)) {\n")
		indent++
		// actions
		f("actions = [")
		csv(w, "%q", maps.Keys(r.Actions))
		f0("];\n")
		f("if(actionMatched(request.action, actions)) {\n")
		indent++
		if r.Condition == nil {
			f("return %q;\n", r.Effect.String())
		} else {
			f("if(")
			saveCondition(w, r.Condition)
			f0(") {\n")
			f("return %q;\n", r.Effect.String())
			f("}\n")
		}
		indent--
		f("}\n")
		indent--
		f("}\n")
	}

	f("return \"EFFECT_DENY\";\n")
	indent--
	f("}\n")
}

func saveCondition(w io.Writer, condition *runtimev1.Condition) {
	f0 := func(format string, a ...any) { // no indentation
		fmt.Fprintf(w, format, a...)
	}
	switch c := condition.Op.(type) {
	case *runtimev1.Condition_Expr:
		fmt.Fprintf(w, c.Expr.Original)
	case *runtimev1.Condition_All:
		n := len(c.All.Expr)
		for i := 0; i < n-1; i++ {
			f0("(")
			saveCondition(w, c.All.Expr[i])
			f0(") && ")
		}
		if n > 1 {
			f0("(")
		}
		saveCondition(w, c.All.Expr[n-1])
		if n > 1 {
			f0(")")
		}
	default:
		panic(fmt.Sprintf(""))
	}
}

func csv[T any](w io.Writer, s string, keys []T) {
	if len(keys) == 0 {
		return
	}
	for i := 0; i < len(keys)-1; i++ {
		fmt.Fprintf(w, s, keys[i])
		fmt.Fprintf(w, ", ")
	}
	fmt.Fprintf(w, s, keys[len(keys)-1])
}
