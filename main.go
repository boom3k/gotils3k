package gotils3k

import (
	"github.com/tealeg/xlsx"
	"log"
)

//Returns a type xlsx.File -- Save file once completed. Thanks -> "github.com/tealeg/xlsx"
func csv2Xlsx(indexSheetName string, dataSet [][]interface{}) *xlsx.File {
	xlsxFile := xlsx.NewFile()
	sheet, err := xlsxFile.AddSheet(indexSheetName)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	for _, datum := range dataSet {
		thisRow := sheet.AddRow()
		for _, data := range datum {
			cell := thisRow.AddCell()
			cell.Value = data.(string)
		}
	}
	return xlsxFile
}
