package xgin

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func QueryInt(c *gin.Context, key string, defaultV int) int {
	str, ok := c.GetQuery(key)
	if !ok {
		return defaultV
	}
	v, _ := strconv.Atoi(str)
	return v
}

func QueryBool(c *gin.Context, key string, defaultV bool) bool {
	str, ok := c.GetQuery(key)
	if !ok {
		return defaultV
	}
	return strings.ToLower(str) == "true"
}

func ParamInt(c *gin.Context, key string, defaultV int) int {
	str := c.Param(key)
	if str == "" {
		return defaultV
	}
	v, _ := strconv.Atoi(str)
	return v
}
