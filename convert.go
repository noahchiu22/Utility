package util

import (
	"fmt"
	"reflect"
)

// return num in what digits you want
func IntToDigits[T string | int | int64](num T, digits int) (numStr string) {
	// 將數量改為 001,002...033 等等形式
	numStr = fmt.Sprint(num)
	for len(numStr) < digits {
		numStr = "0" + numStr
	}
	return
}

// 將結構體內的指定的field name轉成指定的型態陣列(支援string, int, float, bool)
// 指定結構體跟回傳陣列型態
func FieldsToSlice[structT comparable, T comparable](
	structSlice []structT, fieldName string) (outSlice []T) {
	// 跑過結構體的陣列
	for _, item := range structSlice {
		var outValue T
		// 結構體的reflect value, 要用來判斷擷取資料的
		v := reflect.ValueOf(item)
		// outValue的reflect value, 要用來寫入資料的
		ov := reflect.ValueOf(&outValue).Elem()

		// 如果outValue的reflect value不能寫入就跳過
		if !ov.CanSet() {
			continue
		}
		// 取得該結構體的某個field資料 by field name
		fieldValue := v.FieldByName(fieldName)
		// 查無欄位名稱
		if !fieldValue.IsValid() {
			fmt.Printf("filed name %s doesn't exist\n", fieldName)
			return
		}
		// 該欄位的型態跟指定的陣列型態不同
		if fieldValue.Kind() != ov.Kind() {
			fmt.Printf("Different kind by %s(%s) and T(%s)\n",
				fieldName, fieldValue.Kind(), ov.Kind())
			return
		}
		// 判斷型態，依照不同型態寫入outValue
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

		// 將outValue append到要回傳的陣列裡
		outSlice = append(outSlice, outValue)
	}
	return
}
