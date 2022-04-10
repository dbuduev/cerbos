package main

import (
	"testing"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"path"
	"github.com/stretchr/testify/require"
	"strings"
	"os"
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
	t.Run("address", func(t *testing.T) {
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

	t.Run("complex_object", func(t *testing.T) {
		sch := compile("complex_object.json")
		_, err := NewClass("ComplexObject", sch)
		is.NoError(err)
	})
}

func TestSaveAsClass(t *testing.T) {
	is := require.New(t)
	compile := getCompileFunc(t)
	sch := compile("address.json")
	sb := new(strings.Builder)
	SaveAsClass(sb, "Address", sch)

	dat, err := os.ReadFile("./testdata/address.golden.txt")
	is.NoError(err)
	is.Equal(string(dat), sb.String())
}
