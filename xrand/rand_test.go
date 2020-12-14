package xrand

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGenRandNumString(t *testing.T) {
	for i := 1; i <= 1000; i++ {
		rn := GenRandNumString(i)
		assert.Equal(t, len(rn), i)
	}
}
