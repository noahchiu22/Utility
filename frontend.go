package util

import (
	"fmt"
	"reflect"
)

// 計算前端合併需要的rowSpan，需要先使用util.FieldsToSlice產生mergeCol
func CalculateRowSpan[T comparable](structWithRowSpan []T, mergeCol []string) (err error) {
	// 結構體陣列的長度跟合併欄位的長度不同就回傳錯誤
	if len(structWithRowSpan) != len(mergeCol) {
		Log("CountRowSpan 陣列長度錯誤")
		fmt.Println("CountRowSpan 陣列長度錯誤")
		return fmt.Errorf("操作錯誤")
	}

	// 紀錄當前合併欄位值
	currentValue := ""
	for i := 0; i < len(structWithRowSpan); i++ {
		// 取出第i個結構體的RowSpan用來寫入數量
		v := reflect.ValueOf(&structWithRowSpan[i]).Elem().FieldByName("RowSpan")

		// 無此欄位
		if !v.IsValid() {
			Log("CountRowSpan 欄位名稱錯誤")
			fmt.Println("CountRowSpan 欄位名稱錯誤")
			return fmt.Errorf("操作錯誤")
		}
		// 欄位型態不是int系列
		if !v.CanInt() {
			Log("CountRowSpan 欄位型態錯誤")
			fmt.Println("CountRowSpan 欄位型態錯誤")
			return fmt.Errorf("操作錯誤")
		}

		// 如果換了一個合併欄位值
		if currentValue != mergeCol[i] {
			// 先寫入當下欄位值，要用來往後跑迴圈
			currentValue = mergeCol[i]
			// 紀錄有幾筆重複的資料
			rowSpan := 0
			// 從第i個資料往後跑
			for i+rowSpan < len(mergeCol) {
				// 如果換欄位值了就跳開不計算rowSpan了
				if mergeCol[i+rowSpan] != currentValue {
					break
				}
				tv := reflect.ValueOf(&structWithRowSpan[i+rowSpan]).Elem().FieldByName("RowSpan")
				// 將所有算過的資料的rowSpan寫成0，離開這個迴圈之後第一筆會被寫回rowSpan
				tv.SetInt(0)
				rowSpan++
			}

			// 寫入計算後的rowSpan數量，此時i是合併欄位裡的第一筆
			v.SetInt(int64(rowSpan))
			// index往後加rowSpan-1的數量(因為已經寫成0了)
			i += (rowSpan - 1)
		}
	}

	return
}
