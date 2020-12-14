package validator

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

//go test -v -run='TestIsIdCard'
func TestIsIdCard(t *testing.T) {
	//assert.Equal(t, IdValidateReg("420902198710243215"), nil)
	assert.Equal(t, IdValidateArea("420902198710243215"), nil)
	assert.Equal(t, IdValidateBirth("420902198710243215"), nil)
	assert.Equal(t, IdValidateSum("420902198710243215"), nil)
	assert.Equal(t, IsIdCard("420902198710243215"), true)
}

//go test -v -run='TestIsIdPhone'
func TestIsIdPhone(t *testing.T) {
	assert.Equal(t, IsPhone("13430510029"), true)
	assert.Equal(t, IsPhone("134305100292"), false)
	assert.Equal(t, IsPhone("13430510"), false)
	assert.Equal(t, IsPhone("13430510a029"), false)
	assert.Equal(t, IsPhone("+8613430510029"), true)
	assert.Equal(t, IsPhone("+8513430510029"), true)
	assert.Equal(t, IsPhone("+8513430510014349"), true)
	assert.Equal(t, IsPhone("+8613430510029434"), false)
	assert.Equal(t, IsPhone("+86134305100"), false)
}
