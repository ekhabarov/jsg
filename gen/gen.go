package gen

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io"
	"net/url"
	"sort"
	"strings"

	"github.com/ekhabarov/jsg/ast"
	"github.com/iancoleman/strcase"
)

var (
	ErrWriteStruct  = errors.New("failed to write structure to output")
	ErrWriteImports = errors.New("failed to write imports to output")
	ErrNoProps      = errors.New("there is no properties")
)

type option func(string) string

type piece func() string

// One schema output per file.

func Generate(w io.Writer, s *ast.Schema) error {
	buf := bytes.NewBuffer([]byte{})

	p(buf, header())

	if s.ID == "" {
		fmt.Fprintf(w, "%s", buf.String())

		return nil
	}

	err := structure(buf, s)
	if err != nil {
		return fmt.Errorf("failed to build struct header: %w", err)
	}

	// p(buf, sheader)

	b, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("gofmt failed: %w", err)
	}

	fmt.Fprintf(w, "%s", b)

	return nil
}

func p(w io.Writer, s string, args ...interface{}) {
	fmt.Fprintf(w, s, args...)
}

func header() string { return "package schema\n" }

func structure(out io.Writer, s *ast.Schema) error {
	if len(s.Properties) < 1 {
		return ErrNoProps
	}

	name, err := SchemaName(s.ID)
	if err != nil {
		return fmt.Errorf("failed to find schema name: %w", err)
	}

	// temporary buffer
	w := bytes.NewBuffer([]byte{})
	imports := bytes.NewBuffer([]byte{})
	// map of unique imports
	uimports := map[string]struct{}{}

	fmt.Fprintf(w, "type %s struct {\n", name)

	keys := []string{}

	for n := range s.Properties {
		keys = append(keys, n)
	}

	sort.Strings(keys)

	for _, n := range keys {
		p := s.Properties[n]

		t, imp, err := ast.GoType(p.Type, p.Format)
		if err != nil {
			return fmt.Errorf("failed to find Go type for the schema type %q with format %q", p.Type, p.Format)
		}

		if imp != "" {
			uimports[imp] = struct{}{}
		}

		fmt.Fprintf(w, "%s %s\n", n, t)
	}

	fmt.Fprintln(w, "}")

	if len(uimports) > 0 {
		ik := []string{}
		for i := range uimports {
			ik = append(ik, i)
		}

		sort.Strings(ik)

		for _, k := range ik {
			fmt.Fprintf(imports, "import %q\n", k)
		}

		_, err = io.Copy(out, imports)
		if err != nil {
			return ErrWriteImports
		}
	}

	_, err = io.Copy(out, w)
	if err != nil {
		return ErrWriteStruct
	}

	return nil
}

// schemaName returns schema name exetracted from $id property.
func SchemaName(id string) (string, error) {
	u, err := url.Parse(id)
	if err != nil {
		return "", fmt.Errorf("$id property is invalid: %w", err)
	}

	if u.Path == "" {
		return "", errors.New("invalid schema name")
	}

	f := u.Path[strings.LastIndex(u.Path, "/")+1:]
	f = f[:strings.Index(f, ".")]

	return strcase.ToCamel(f), nil
}
