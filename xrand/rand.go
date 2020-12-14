package xrand

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

func GenRandNumString(length int) string {
	str := ""
	for i := 0; ; i++ {
		n := r.Int63n(1000000000000000)
		str += fmt.Sprintf("%d", n)
		if len(str) >= length {
			break
		}
	}
	return str[:length]
}
