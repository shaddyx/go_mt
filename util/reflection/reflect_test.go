package reflection

import (
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
