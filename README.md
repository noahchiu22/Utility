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
## Create excel with field tag
```go
// Using any excel tag as title in the struct to create excel
// And it will automatically fit column width by the cell content
func CreateExcel[S comparable](data []S, path, filename string) (err error) {
	f := excelize.NewFile()

	defer func() {
		if err := f.Close(); err != nil {
			return
		}
	}()

	headers := []string{}
	for row, item := range data {
		t := reflect.TypeOf(item)
		// 第一筆資料依照struct的tag excel製作headers
		if row == 0 {
			for i := 0; i < t.NumField(); i++ {
				excelTag := t.Field(i).Tag.Get("excel")
				if excelTag != "" {
					headers = append(headers, excelTag)
				}
			}

			for col, header := range headers {
				cell, _ := excelize.CoordinatesToCellName(col+1, 1)
				f.SetCellValue("Sheet1", cell, header)
			}
		}
		v := reflect.ValueOf(item)
		for col := 0; col < len(headers); col++ {
			cell, _ := excelize.CoordinatesToCellName(col+1, row+2)
			field := v.Field(col)
			excelTag := t.Field(col).Tag.Get("excel")
			// 沒有excelTag就跳過
			if excelTag == "" {
				continue
			}

			value := field.Interface()
			switch temp := value.(type) {
			case int64:
				value = temp
			case float64:
				value = temp
			case time.Time:
				value = temp.Format("2006-01-02 15:04:05")
			default:
				value = temp
			}
			f.SetCellValue("Sheet1", cell, value)
		}
	}

	AutofitColumn(f)

	fmt.Println(path + filename)
	// Save spreadsheet by the given path.
	if err = f.SaveAs(path + filename); err != nil {
		return
	}

	return
}
```