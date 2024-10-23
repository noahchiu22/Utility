# Go utility
## Convert struct slice into a slice
```go
// Convert a specified field name in a struct to a slice of a specified type (supports string, int, float, bool)
// Specify the struct type and the return slice type
func FieldsToSlice[structT comparable, T comparable](
	structSlice []structT, fieldName string) (outSlice []T) {
	// Iterate through the slice of structs
	for _, item := range structSlice {
		var outValue T
		// Reflect value of the struct, used to access field data
		v := reflect.ValueOf(item)
		// Reflect value of outValue, used to set the data
		ov := reflect.ValueOf(&outValue).Elem()

		// Skip if the reflect value of outValue cannot be set
		if !ov.CanSet() {
			continue
		}
		// Get the value of the specified field by field name
		fieldValue := v.FieldByName(fieldName)
		// Field name does not exist
		if !fieldValue.IsValid() {
			fmt.Printf("field name %s doesn't exist\n", fieldName)
			return
		}
		// Field type is different from the specified slice type
		if fieldValue.Kind() != ov.Kind() {
			fmt.Printf("Different kind for %s(%s) and T(%s)\n",
				fieldName, fieldValue.Kind(), ov.Kind())
			return
		}
		// Determine the type and assign the value to outValue accordingly
		switch fieldValue.Kind() {
		case reflect.String:
			ov.SetString(fieldValue.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			ov.SetInt(fieldValue.Int())
		case reflect.Float32, reflect.Float64:
			ov.SetFloat(fieldValue.Float())
		case reflect.Bool:
			ov.SetBool(fieldValue.Bool())
		}

		// Append outValue to the return slice
		outSlice = append(outSlice, outValue)
	}
	return
}
```
