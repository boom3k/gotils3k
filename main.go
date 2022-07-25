package gotils3k

import (
	"github.com/tealeg/xlsx"
	"log"
)

// CSV2XLSX Returns a type xlsx.File -- Save file once completed. Thanks -> "github.com/tealeg/xlsx"
func CSV2XLSX(indexSheetName string, dataSet [][]string) *xlsx.File {
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
			cell.Value = data
		}
	}
	return xlsxFile
}

//CSV2XLSX_RAW - Handles interface slices for CSV2XLSX
func CSV2XLSX_RAW(indexSheetName string, dataSet [][]interface{}) *xlsx.File {
	var asStringSlice [][]string
	for _, columnData := range dataSet {
		var data []string
		for _, datum := range columnData {
			data = append(data, datum.(string))
		}
		asStringSlice = append(asStringSlice, data)
	}
	return CSV2XLSX(indexSheetName, asStringSlice)
}
