package auth

import (
	"fmt"
	"net/url"
	"github.com/liyue201/golib/xhash"
	"sort"
)

func Sign(method, path string, queries url.Values, secret string) string {
	queryStr := sortQueries(queries)
	return sign(method, path, queryStr, secret)
}

func VerifySign(method, path string, queries url.Values, secret string, signStr string) bool {
	s := Sign(method, path, queries, secret)
	return signStr == s
}

func sortQueries(queryMap url.Values) string {
	var keysList []string
	for k := range queryMap {
		if k != "sign" {
			keysList = append(keysList, k)
		}
	}
	sort.Strings(keysList)
	var queryStr string
	for i, key := range keysList {
		if i > 0 {
			queryStr += "&"
		}
		v := queryMap[key]
		if len(v) == 1 {
			queryStr += fmt.Sprintf("%s=%v", key, queryMap[key][0])
		} else {
			queryStr += fmt.Sprintf("%s=%v", key, queryMap[key])
		}
	}
	return queryStr
}

func sign(method, path, queryStr, secret string) string {
	str := method + path + queryStr + secret
	return xhash.Md5([]byte(str))
}
