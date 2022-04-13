package main

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
	"errors"
	"fmt"
	"io"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"unicode"
)

type AttrType int

const (
	Unknown AttrType = iota
	Int
	Number
	String
	Object
	Array
	Boolean
)

type ClassAttr struct {
	Name     string
	Type     AttrType
	Required bool
}

type Class struct {
	Name  string
	Attrs []ClassAttr
}

var WrongSchemaTypeErr = errors.New("wrong schema type")

func NewClass(name string, s *jsonschema.Schema) (*Class, error) {
	if s.Types[0] != "object" {
		return nil, fmt.Errorf("expected %q, got %q: %w", "object", s.Types[0], WrongSchemaTypeErr)
	}
	cl := Class{Name: name}
	cl.Attrs = make([]ClassAttr, 0, len(s.Properties))

	for n, schema := range s.Properties {
		t, err := getAttrType(schema.Types[0])
		if err != nil {
			return nil, err
		}
		attr := ClassAttr{
			Name: n,
			Type: t,
		}
		for _, n2 := range s.Required {
			if n == n2 {
				attr.Required = true
				break
			}
		}
		cl.Attrs = append(cl.Attrs, attr)
	}

	return &cl, nil
}

func getAttrType(s string) (AttrType, error) {
	switch s {
	case "integer":
		return Int, nil
	case "string":
		return String, nil
	}
	return 0, fmt.Errorf("unexpected attribute schema type: %w", WrongSchemaTypeErr)
}

func (a *Class) Equal(b *Class) bool {
	if a == b {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Name != b.Name {
		return false
	}

	if len(a.Attrs) != len(b.Attrs) {
		return false
	}

	for i := range a.Attrs {
		if !a.Attrs[i].Equal(b.Attrs[i]) {
			return false
		}
	}

	return true
}

func (a ClassAttr) Equal(b ClassAttr) bool {
	if a.Name != b.Name {
		return false
	}

	if a.Required != b.Required {
		return false
	}

	if a.Type != b.Type {
		return false
	}

	return true
}

const baseClasses = `
class Request {
  requestId: string;
  action: string;
  principal: Principal;
  resource: Resource;
}

class Principal {
  id: string;
  roles: Array<string>;
  attr: principalAttr;
}

class Resource {
  id: string;
  attr: resourceAttr;
}`

func SaveAsClass(w io.Writer, name string, s *jsonschema.Schema) {
	fmt.Fprintf(w, "class %s {\n", name)

	inn := make(map[string]*jsonschema.Schema) // "inner" classes
	props := maps.Keys(s.Properties)
	slices.Sort(props)
	for _, n := range props {
		t := s.Properties[n].Types[0]
		if t == "object" {
			cls := fmt.Sprintf("%s%c%s", name, unicode.ToTitle(rune(n[0])), n[1:])
			fmt.Fprintf(w, "  %s: %s", n, cls)
			inn[cls] = s.Properties[n]
		} else if t == "array" {
			fmt.Fprintf(w, "  %s: Array<%s>", n, s.Properties[n].Items2020.Types[0])
		} else {
			fmt.Fprintf(w, "  %s: %s", n, t)
		}
		for _, n2 := range s.Required {
			if n == n2 {
				fmt.Fprintf(w, " | null")
				break
			}
		}
		fmt.Fprintf(w, ";\n")
	}
	fmt.Fprintf(w, "}\n")

	for n, s2 := range inn {
		fmt.Fprintln(w)
		SaveAsClass(w, n, s2)
	}
}
