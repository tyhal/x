package fmap

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_flatten(t *testing.T) {
	m := map[string]int{"b": 1, "a": 2, "c": 3}
	flattened := New(m)

	slices.SortFunc(flattened, flattened.CmpV(false /*asc*/))
	expected := FlattenedKVs[string, int]{
		{"b", 1},
		{"a", 2},
		{"c", 3},
	}
	assert.Equal(t, expected, flattened)

	slices.SortFunc(flattened, flattened.CmpV(true /*desc*/))
	expected = FlattenedKVs[string, int]{
		{"c", 3},
		{"a", 2},
		{"b", 1},
	}
	assert.Equal(t, expected, flattened)

	slices.SortFunc(flattened, flattened.CmpK(false /*asc*/))
	expected = FlattenedKVs[string, int]{
		{"a", 2},
		{"b", 1},
		{"c", 3},
	}
	assert.Equal(t, expected, flattened)
}
