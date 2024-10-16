package util

import (
	"reflect"
	"strconv"
	"time"
)

// 依照soap給的field來製作一個struct
// func makeNewStruct(columns []Field, ignoreFields ...string) reflect.Value {
// 	// 要讓第一個字大寫來製作field
// 	title := cases.Title(language.English, cases.NoLower)

// 	// 同義於 fields := []reflect.StructField{} with columns length
// 	fields := make([]reflect.StructField, 0, len(columns))

// 	// 跑過每個應該建立的欄位
// 	for _, field := range columns {
// 		// 判斷忽略欄位，沒有給的話這裡會是false
// 		if slices.Contains(ignoreFields, field.Name) {
// 			continue
// 		}
// 		fmt.Println(field.Name)
// 		columnName := columnMap[field.Name]
// 		fieldName := title.String(columnName)
// 		fieldType, _ := GetTypeAndValue(field.Type, field.Value)

// 		// 新增一個field with type and tag
// 		fields = append(fields, reflect.StructField{
// 			Name: fieldName,
// 			Type: fieldType,
// 			Tag:  reflect.StructTag(fmt.Sprintf(`gorm:"column:%s"`, columnName)),
// 		})
// 	}

// 	// 將剛剛做的struct實體化
// 	newStructType := reflect.StructOf(fields)
// 	newStruct := reflect.New(newStructType).Elem()

// 	for _, field := range columns {
// 		// 判斷忽略欄位，沒有給的話這裡會是false
// 		if slices.Contains(ignoreFields, field.Name) {
// 			continue
// 		}
// 		_, fieldValue := GetTypeAndValue(field.Type, field.Value)
// 		fieldName := title.String(columnMap[field.Name])

// 		// set value to the field by name
// 		newStruct.FieldByName(fieldName).Set(fieldValue)
// 	}
// 	// fmt.Printf("New struct: %+v\n", newStruct)

// 	return newStruct
// }

// get reflect type and value
func GetTypeAndValue(dataType, valueStr string) (reflectType reflect.Type, reflectValue reflect.Value) {
	switch dataType {
	case "string":
		reflectValue = reflect.ValueOf(valueStr)
		reflectType = reflect.TypeOf(valueStr)
		return
	case "numeric":
		value, _ := strconv.Atoi(valueStr)
		reflectValue = reflect.ValueOf(float64(value))
		reflectType = reflect.TypeOf(float64(value))
		return
	case "date":
		layout := "20060102"
		loc, _ := time.LoadLocation("Local")
		value, err := time.ParseInLocation(layout, valueStr, loc)
		if err != nil {
			value = time.Time{}
		}
		reflectValue = reflect.ValueOf(value)
		reflectType = reflect.TypeOf(value)
		return
	default:
		panic("unknown data type")
	}
}

// 自動把有對應到的欄位寫入到emptyStruct
func InheritStruct[T comparable](emptyStruct T, structValue reflect.Value) {
	woReflect := reflect.ValueOf(&emptyStruct).Elem()

	// 如果newStruct的fieldName有出現在需要被寫入的struct裡就直接set值進去
	for i := 0; i < structValue.NumField(); i++ {
		fieldName := structValue.Type().Field(i).Name

		if woReflect.FieldByName(fieldName).CanSet() {
			switch value := structValue.Field(i).Interface().(type) {
			case string:
				woReflect.FieldByName(fieldName).SetString(value)
			case float64:
				woReflect.FieldByName(fieldName).SetFloat(value)
			case time.Time:
				woReflect.FieldByName(fieldName).Set(reflect.ValueOf(value))
			}
		}
	}
}
