package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Donkey struct {
	name string
}

func TestAddBool(t *testing.T) {
	orig := []bool{true, false, true}
	actual := Add(orig, true)

	assert.Equal(t, 3, len(orig))
	assert.Equal(t, 4, len(actual))

	assert.Equal(t, []bool{true, false, true, true}, actual)
}

func TestAddString(t *testing.T) {
	orig := []string{"a", "b", "c"}
	actual := Add(orig, "d")

	assert.Equal(t, 3, len(orig))
	assert.Equal(t, 4, len(actual))

	assert.Equal(t, []string{"a", "b", "c", "d"}, actual)
}

func TestAddDonkey(t *testing.T) {
	orig := []Donkey{{name: "sam"}, {name: "mark"}}
	actual := Add(orig, Donkey{name: "mary"})

	assert.Equal(t, 2, len(orig))
	assert.Equal(t, 3, len(actual))
}

func TestAddEmpty(t *testing.T) {
	var orig []string
	actual := Add(orig, "sam")

	assert.Equal(t, 0, len(orig))
	assert.Equal(t, 1, len(actual))
}
