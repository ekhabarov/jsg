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

var ErrNoProps = errors.New("there is no properties")

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

	sheader, err := structHeader(s)
	if err != nil {
		return fmt.Errorf("failed to build struct header: %w", err)
	}

	p(buf, sheader)

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

func structHeader(s *ast.Schema) (string, error) {
	if len(s.Properties) < 1 {
		return "", ErrNoProps
	}

	name, err := SchemaName(s.ID)
	if err != nil {
		return "", fmt.Errorf("failed to find schema name: %w", err)
	}

	// fmt.Sprintf("%v")

	w := bytes.NewBuffer([]byte(fmt.Sprintf("type %s struct {", name)))

	keys := []string{}

	for n, _ := range s.Properties {
		keys = append(keys, n)
	}

	sort.Sort(sort.StringSlice(keys))

	for _, n := range keys {
		p := s.Properties[n]
		fmt.Fprintf(w, "%s %s\n", n, p.Type)
	}

	w.WriteString("}\n")

	return w.String(), nil
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
