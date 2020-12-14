package auth

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestSortQueries(t *testing.T) {
	query := "b=555&j=6af&a=grgfg"
	expect := "a=grgfg&b=555&j=6af"
	v, _ := url.ParseQuery(query)
	assert.Equal(t, expect, sortQueries(v))
}

func TestSignUrl(t *testing.T) {
	secret := "niafidsf0gergr"
	query1 := "b=555&j=6af&a=grgfg"
	query2 := "a=grgfg&b=555&j=6af"

	v, _ := url.ParseQuery(query1)
	signStr := Sign("POST", "api/v1/file/add", v, secret)
	expect := sign("POST", "api/v1/file/add", query2, secret)

	t.Logf("sign=%s", signStr)
	assert.Equal(t, expect, signStr)
}

func TestCheckSign(t *testing.T) {

}
