package vimeo

import (
	"fmt"
	"net/url"
	"strings"
)

type QueryParam interface {
	Get() (key, value string)
}

type Fields []string

func (f Fields) Get() (string, string) {
	return "fields", strings.Join(f, ",")
}

type Page int

func (p Page) Get() (string, string) {
	return "page", fmt.Sprint(p)
}

type PerPage int

func (p PerPage) Get() (string, string) {
	return "per_page", fmt.Sprint(p)
}

func addQueryParam(s string, params ...QueryParam) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	q := u.Query()
	for _, p := range params {
		q.Set(p.Get())
	}

	u.RawQuery = q.Encode()

	return u.String(), nil
}
