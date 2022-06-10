package sorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinInt(t *testing.T) {
	assert.Equal(t, 1, Min(1, 2, 3))
	assert.Equal(t, 1, Min(3, 2, 1))
	assert.Equal(t, 1, Min(3, 1, 2))
	assert.Equal(t, 3, Min(3))
}

func TestMaxInt(t *testing.T) {
	assert.Equal(t, 3, Max(1, 2, 3))
	assert.Equal(t, 3, Max(3, 2, 1))
	assert.Equal(t, 3, Max(3, 1, 2))
	assert.Equal(t, 1, Max(1))
}

func TestMinString(t *testing.T) {
	assert.Equal(t, "a", Min("a", "ab", "abc"))
	assert.Equal(t, "a", Min("abc", "ab", "a"))
	assert.Equal(t, "a", Min("abc", "a", "ab"))
	assert.Equal(t, "abc", Min("abc"))
}

func TestMaxString(t *testing.T) {
	assert.Equal(t, "abc", Max("a", "ab", "abc"))
	assert.Equal(t, "abc", Max("abc", "ab", "a"))
	assert.Equal(t, "abc", Max("abc", "a", "ab"))
	assert.Equal(t, "abc", Max("abc"))
}
