package reflection

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPointer(t *testing.T) {
	var m map[string]any
	var i int = 1

	assert.True(t, IsPointer(&m))
	assert.False(t, IsPointer(m))
	assert.False(t, IsPointer(i))
	assert.True(t, IsPointer(&i))

}

func TestIsMapPointer(t *testing.T) {
	var m map[string]any
	var i int = 1

	assert.True(t, IsMapPointer(&m))
	assert.False(t, IsMapPointer(m))
	assert.False(t, IsMapPointer(i))
	assert.False(t, IsMapPointer(&i))
}

func TestIsStructurePointer(t *testing.T) {
	var m struct {
	}
	var i int = 1

	assert.True(t, IsStructurePointer(&m))
	assert.False(t, IsStructurePointer(m))
	assert.False(t, IsStructurePointer(i))
	assert.False(t, IsStructurePointer(&i))
}

func TestIsStringPointer(t *testing.T) {
	var m string
	var i int = 1

	assert.True(t, IsStringPointer(&m))
	assert.False(t, IsStringPointer(m))
	assert.False(t, IsStringPointer(i))
	assert.False(t, IsStringPointer(&i))
}

func TestUnmarshalTo(t *testing.T) {
	s := "123"
	f := 1.25
	i := 1
	m := make(map[string]any)
	err := UnmarshalTo("1", &s)
	assert.NoError(t, err)
	err = UnmarshalTo("2.5", &f)
	assert.Error(t, err)
	err = UnmarshalTo("2", &i)
	assert.Error(t, err)
	err = UnmarshalTo(`{"a":1}`, &m)
	assert.NoError(t, err)
	fmt.Println("s", s)
	fmt.Println("m", m)
	//assert.True(t, IsStringPointer(&m))
}
