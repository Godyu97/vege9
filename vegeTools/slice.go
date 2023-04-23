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
