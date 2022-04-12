package main

import (
	"io"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"fmt"
	"golang.org/x/exp/maps"
)

func Save(w io.Writer, rps *runtimev1.RunnableResourcePolicySet) {
	indent := 0
	f0 := func(format string, a ...any) {
		fmt.Fprintf(w, format, a...)
	}
	f := func(format string, a ...any) {
		for i := 0; i < indent; i++ {
			fmt.Fprintf(w, "  ")
		}
		fmt.Fprintf(w, format, a...)
	}
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
		f("  actions = [")
		csv(w, "%q", maps.Keys(r.Actions))
		f0("];\n")
		f("  if(actionsMatched(request.action, actions)) {\n")
		indent++
		if r.Condition == nil {
			f("    return %q;\n", r.Effect.String())
		}
		indent--
		f("  }\n")
		indent--
		f("}\n")
	}

	f("return \"EFFECT_DENY\";\n")
	indent--
	f("}\n")
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
