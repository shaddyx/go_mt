package reflection

import "reflect"

func IsPointer(s any) bool {
	return reflect.TypeOf(s).Kind() == reflect.Pointer
}

func IsMapPointer(s any) bool {
	return IsPointer(s) &&
		reflect.TypeOf(s).Elem().Kind() == reflect.Map
}

func IsStructurePointer(s any) bool {
	return IsPointer(s) &&
		reflect.TypeOf(s).Elem().Kind() == reflect.Struct
}

func IsStringPointer(s any) bool {
	return IsPointer(s) &&
		reflect.TypeOf(s).Elem().Kind() == reflect.String
}
