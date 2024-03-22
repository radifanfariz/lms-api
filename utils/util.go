package utils

import "reflect"

// removeDuplicates removes duplicate items from a slice based on a specified key
func RemoveDuplicates(slice interface{}, key string) interface{} {
	encountered := map[interface{}]bool{}
	resultSlice := reflect.MakeSlice(reflect.TypeOf(slice), 0, 0)

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return nil // Return nil if the input is not a slice
	}

	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i)
		keyValue := reflect.Indirect(item).FieldByName(key).Interface()

		if encountered[keyValue] {
			continue // Duplicate found, skip adding to the result
		}
		resultSlice = reflect.Append(resultSlice, item)
		encountered[keyValue] = true
	}

	return resultSlice.Interface()
}
