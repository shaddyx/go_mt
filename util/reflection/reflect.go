package reflection

import (
	"encoding/json"
	"fmt"
	"reflect"
)

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

func UnmarshalTo(jsn string, output any) error {
	bJsn := []byte(jsn)
	var err error
	if !IsPointer(output) {
		return fmt.Errorf("type is not a pointer: %s", reflect.TypeOf(output))
	} else if IsStringPointer(output) {
		output = jsn
	} else if IsStructurePointer(output) {
		err = json.Unmarshal(bJsn, output)
	} else if IsMapPointer(output) {
		err = json.Unmarshal(bJsn, output)
	} else {
		err = fmt.Errorf("unknown output type: %s", reflect.TypeOf(output).Name())
	}
	return err
}
