// Package fmap is a utility for flattening maps into slices of key-value pairs
package fmap

import "cmp"

type flatKV[K, V cmp.Ordered] struct {
	K K
	V V
}

// FlattenedKVs can be used to convert a map into a sorted slice of key-value pairs
type FlattenedKVs[K, V cmp.Ordered] []flatKV[K, V]

// CmpV returns the sort function for the values
func (FlattenedKVs[K, V]) CmpV(desc bool) func(a, b flatKV[K, V]) int {
	return func(a, b flatKV[K, V]) int {
		if desc {
			a, b = b, a
		}
		return cmp.Compare(a.V, b.V)
	}
}

// CmpK returns the sort function for the keys
func (FlattenedKVs[K, V]) CmpK(desc bool) func(a, b flatKV[K, V]) int {
	return func(a, b flatKV[K, V]) int {
		if desc {
			a, b = b, a
		}
		return cmp.Compare(a.K, b.K)
	}
}

// New returns a sorted slice of key-value pairs from a map
func New[K, V cmp.Ordered](m map[K]V) FlattenedKVs[K, V] {
	out := make([]flatKV[K, V], 0, len(m))
	for k, v := range m {
		out = append(out, flatKV[K, V]{k, v})
	}
	return out
}
