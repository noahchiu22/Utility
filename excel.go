package util

import (
	"reflect"

	"github.com/xuri/excelize/v2"
)

func CreateExcel[S comparable](data []S, path, filename string, headers ...string) (err error) {
	f := excelize.NewFile()

	defer func() {
		if err := f.Close(); err != nil {
			return
		}
	}()

	for col, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(col+1, 1)
		f.SetCellValue("Sheet1", cell, header)
	}

	for row, item := range data {
		rft := reflect.ValueOf(item)
		for col := 0; col < len(headers); col++ {
			cell, _ := excelize.CoordinatesToCellName(col+1, row+2)
			field := rft.Field(col)

			var value interface{}
			switch field.Type().String() {
			case "int64":
				value = field.Int()
			case "float64":
				value = field.Float()
			default:
				value = field.String()
			}
			f.SetCellValue("Sheet1", cell, value)
		}
	}

	// Save spreadsheet by the given path.
	if err = f.SaveAs(path + filename); err != nil {
		return
	}

	return
}
