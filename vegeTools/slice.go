package vegeTools

func ItemIsInSlice[T comparable](item T, slice []T) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func SliceIsInSlice[T comparable](item []T, slice []T) bool {
	for _, i := range item {
		found := false
		for _, s := range slice {
			if s == i {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func SliceSub[T comparable](s1 []T, s2 []T) []T {
	diff := make([]T, 0)
	tmp := make(map[T]struct{}, 0)
	for _, v2 := range s2 {
		if _, ok := tmp[v2]; ok == false {
			tmp[v2] = struct{}{}
		}
	}
	for _, v1 := range s1 {
		if _, ok := tmp[v1]; ok == false {
			diff = append(diff, v1)
		}
	}
	return diff
}

func SliceUnique[T comparable](arr []T) []T {
	tmp := make(map[T]struct{})
	l := len(arr)
	if l == 0 {
		return arr
	}
	rel := make([]T, 0, l)
	for _, item := range arr {
		_, ok := tmp[item]
		if ok {
			continue
		}
		tmp[item] = struct{}{}
		rel = append(rel, item)
	}
	return rel[:len(tmp)]
}

func SliceEqual[T comparable](s1, s2 []T) bool {
	// If one is nil, the other must also be nil.
	if (s1 == nil) != (s2 == nil) {
		return false
	}
	if len(s1) != len(s2) {
		return false
	}
	for i, _ := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
