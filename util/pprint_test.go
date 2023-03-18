package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDump(t *testing.T) {
	s := struct {
		A string
		B int
	}{
		A: "123",
		B: 12,
	}
	res, _ := Dump(s)
	assert.Equal(t, "{\n  \"A\": \"123\",\n  \"B\": 12\n}", res)

}
