package vegeTools

// map[string] interface ---> obj
func MapToObj(m map[string]interface{}, obj interface{}) error {
	j, err := Marshal(m)
	if err != nil {
		return err
	}
	return Unmarshal(j, obj)
}

// obj --->  map[string] interface
func ObjToMap(obj interface{}, m *map[string]interface{}) error {
	j, err := Marshal(obj)
	if err != nil {
		return err
	}
	return Unmarshal(j, m)
}
