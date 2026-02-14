package vege

// --------------------------
// 基础指针/值转换 (Basic)
// --------------------------

// Ptr2Value returns ptr's value; if ptr==nil, return zero value of T.
func Ptr2Value[T any](ptr *T) T {
	if ptr != nil {
		return *ptr
	}
	var zero T
	return zero
}

// Ptr2ValueWithDefault returns ptr's value; if ptr==nil, return the specific defaultVal.
// 场景：如果指针为空，我想返回 -1 或者 "unknown"，而不是 0 或 ""。
func Ptr2ValueWithDefault[T any](ptr *T, defaultVal T) T {
	if ptr != nil {
		return *ptr
	}
	return defaultVal
}

// Value2Ptr returns &v.
func Value2Ptr[T any](v T) *T {
	return &v
}

// --------------------------
// 切片批量转换 (Slice Batch)
// --------------------------

// SliceValue2Ptr converts a slice of values to a slice of pointers.
// 场景：批量构造数据库模型或 API 请求体。
func SliceValue2Ptr[T any](s []T) []*T {
	if s == nil {
		return nil
	}
	result := make([]*T, len(s))
	for i := range s {
		result[i] = &s[i] // 注意：这里取的是切片底层数组元素的地址
	}
	return result
}

// SlicePtr2Value converts a slice of pointers to a slice of values.
// Nil pointers in the slice will be converted to zero values.
// 场景：处理 API 响应，将 []*User 转为 []User。
func SlicePtr2Value[T any](s []*T) []T {
	if s == nil {
		return nil
	}
	result := make([]T, len(s))
	for i, v := range s {
		result[i] = Ptr2Value(v)
	}
	return result
}

// --------------------------
// Map 批量转换 (Map Batch)
// --------------------------

// MapValue2Ptr converts map[K]V to map[K]*V.
func MapValue2Ptr[K comparable, V any](m map[K]V) map[K]*V {
	if m == nil {
		return nil
	}
	result := make(map[K]*V, len(m))
	for k, v := range m {
		val := v // 创建副本以获取地址
		result[k] = &val
	}
	return result
}

// MapPtr2Value converts map[K]*V to map[K]V.
func MapPtr2Value[K comparable, V any](m map[K]*V) map[K]V {
	if m == nil {
		return nil
	}
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = Ptr2Value(v)
	}
	return result
}

// --------------------------
// 零值与空值逻辑 (Zero & Nil Logic)
// --------------------------

// PtrOrNil returns nil if v is the zero value of T, otherwise returns &v.
// 场景：在更新数据库时，如果字段是零值（如空字符串），不想传空指针，而是传 nil 以避免覆盖。
// 注意：T 必须是可比较类型 (comparable)。
func PtrOrNil[T comparable](v T) *T {
	var zero T
	if v == zero {
		return nil
	}
	return &v
}

// IsNilOrZero checks if a pointer is nil OR if it points to a zero value.
func IsNilOrZero[T comparable](ptr *T) bool {
	if ptr == nil {
		return true
	}
	var zero T
	return *ptr == zero
}

// --------------------------
// 条件辅助 (Conditional)
// --------------------------

// If returns trueVal if condition is true, otherwise falseVal.
// 场景：由行三元运算符： vege.If(age > 18, "Adult", "Minor")
func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// IfPtr returns &trueVal if condition is true, otherwise &falseVal.
func IfPtr[T any](condition bool, trueVal, falseVal T) *T {
	if condition {
		return &trueVal
	}
	return &falseVal
}
