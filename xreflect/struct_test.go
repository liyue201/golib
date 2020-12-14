package xreflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructCopy(t *testing.T) {
	type A struct {
		Host     string
		Port     int
		User     string
		Password string
		Db       int
		Exa      string
	}

	type B struct {
		Host     string
		User     string
		Port     int
		Password string
		Db       string
		Exb      string
	}

	a := A{
		Host:     "aaa",
		Port:     6616,
		User:     "test",
		Password: "123456",
		Db:       1,
		Exa:      "eeee",
	}
	b := B{}

	StructCopy(&a, &b)

	t.Logf("a: %+v", a)
	t.Logf("b: %+v", b)

	assert.Equal(t, a.Host, b.Host)
}
