package lib

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/iancoleman/strcase"
)

// URLName returns schema name exetracted from $id property.
func URLName(id string) (string, error) {
	u, err := url.Parse(id)
	if err != nil {
		return "", fmt.Errorf("$id property is invalid: %w", err)
	}

	if u.Path == "" {
		return "", errors.New("invalid url")
	}

	f := u.Path[strings.LastIndex(u.Path, "/")+1:]
	f = f[:strings.Index(f, ".")]

	return strcase.ToCamel(f), nil
}
