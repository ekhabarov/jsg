package gen

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io"
	"sort"

	"github.com/ekhabarov/jsg/ast"
	"github.com/ekhabarov/jsg/lib"
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

	if err := structure(buf, s); err != nil {
		return fmt.Errorf("failed to build struct header: %w", err)
	}

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

	name, err := lib.URLName(s.ID)
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

		t, imp, err := ast.GoType(p.Type, p.Format, p.Ref)
		if err != nil {
			return fmt.Errorf("go type not found: schema type: %q, format: %v", p.Type, p.Format)
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
