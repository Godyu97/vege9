package vegeTools

// map[string] interface ---> obj
func MapToObj(m map[string]any, obj any) error {
	j, err := JsonMarshal(m)
	if err != nil {
		return err
	}
	return JsonUnmarshal(j, obj)
}

// obj --->  map[string] interface
func ObjToMap(obj any, m *map[string]any) error {
	j, err := JsonMarshal(obj)
	if err != nil {
		return err
	}
	return JsonUnmarshal(j, m)
}

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package maps defines various functions useful with maps of any type.

// MapKeys returns the keys of the map m.
// The keys will be in an indeterminate order.
func MapKeys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// MapValues returns the values of the map m.
// The values will be in an indeterminate order.
func MapValues[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// MapEqual reports whether two maps contain the same key/value pairs.
// MapValues are compared using ==.
func MapEqual[M1, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

// MapEqualFunc is like MapEqual, but compares values using eq.
// MapKeys are still compared with ==.
func MapEqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](m1 M1, m2 M2, eq func(V1, V2) bool) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || !eq(v1, v2) {
			return false
		}
	}
	return true
}

// MapClear removes all entries from m, leaving it empty.
func MapClear[M ~map[K]V, K comparable, V any](m M) {
	for k := range m {
		delete(m, k)
	}
}

// MapClone returns a copy of m.  This is a shallow clone:
// the new keys and values are set using ordinary assignment.
func MapClone[M ~map[K]V, K comparable, V any](m M) M {
	// Preserve nil in case it matters.
	if m == nil {
		return nil
	}
	r := make(M, len(m))
	for k, v := range m {
		r[k] = v
	}
	return r
}

// MapCopy copies all key/value pairs in src adding them to dst.
// When a key in src is already present in dst,
// the value in dst will be overwritten by the value associated
// with the key in src.
func MapCopy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) {
	for k, v := range src {
		dst[k] = v
	}
}

// MapDeleteFunc deletes any key/value pairs from m for which del returns true.
func MapDeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool) {
	for k, v := range m {
		if del(k, v) {
			delete(m, k)
		}
	}
}
