package vege

// Ptr2Value return ptr's value; if ptr==nil,return zero Ts
func Ptr2Value[T any](ptr *T) T {
	var res T
	if ptr != nil {
		res = *ptr
	}
	return res
}

// Value2Ptr return &v
func Value2Ptr[T any](v T) *T {
	return &v
}

func NewPtr[T any](v T) *T {
	r := new(T)
	*r = v
	return r
}
