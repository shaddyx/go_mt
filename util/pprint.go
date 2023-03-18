package util

import (
	"encoding/json"
)

func Dump(data interface{}) (string, error) {
	b, err := json.MarshalIndent(data, "", "  ")
	return string(b), err
}
