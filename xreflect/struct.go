package xreflect

import (
	"reflect"
)

// StructCopy copies the filed of a to b
// a and b must be a pointer to a struct
func StructCopy(a interface{}, b interface{}) {
	va := reflect.ValueOf(a).Elem()
	vb := reflect.ValueOf(b).Elem()
	for i := 0; i < va.NumField(); i++ {
		item1 := va.Field(i)
		type1 := va.Type().Field(i)
		item2 := vb.FieldByName(type1.Name)

		if !item2.IsValid() || item1.Kind() != item2.Kind() {
			continue
		}
		if !item2.CanSet() {
			continue
		}
		item2.Set(item1)
	}
}
