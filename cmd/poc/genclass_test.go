package main

import (
	"testing"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"path"
	"github.com/stretchr/testify/require"
)

func getCompileFunc(t *testing.T) func(s string) *jsonschema.Schema {
	return func(s string) *jsonschema.Schema {
		schema, err := jsonschema.Compile(path.Join("testdata", s))
		require.NoError(t, err)
		return schema
	}
}
func TestNewClass(t *testing.T) {
	is := require.New(t)
	compile := getCompileFunc(t)
	t.Run("address.json", func(t *testing.T) {
		sch := compile("address.json")
		cl, err := NewClass("Address", sch)
		is.NoError(err)
		is.True(cl.Equal(&Class{
			Name: "Address",
			Attrs: []ClassAttr{
				{"street_address", String, true},
				{"city", String, true},
				{"state", String, true},
			},
		}))
	})
}
