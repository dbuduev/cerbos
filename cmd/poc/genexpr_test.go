package main

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/cerbos/cerbos/internal/conditions"
	"strings"
)

func TestExpr(t *testing.T) {
	is := require.New(t)
	s := `request.resource.attr.status == "PENDING_APPROVAL"`

	ast, iss := conditions.StdEnv.Compile(s)
	is.Nil(iss)
	e := ast.Expr()
	is.NotNil(e)

	sb := new(strings.Builder)
	renderExpr(sb, e)
	is.Equal(s, sb.String())
}
