package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {
	var startNum uint64
	f := excelize.NewFile()
	startNum = 898611222220217723

	//for i := 1; i < 50001; i++ {
	for i := 1; i < 11; i++ {
		newNum := startNum + 1
		startNum = newNum
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i), strconv.FormatUint(newNum, 10))
	}

	if err := f.SaveAs("物联网卡号.xlsx"); err != nil {
		fmt.Println(err)
	}
	// // Create a new sheet.
	// index := f.NewSheet("Sheet2")
	// // Set value of a cell.
	// f.SetCellValue("Sheet2", "A2", "Hello world.")

	// Set active sheet of the workbook.
	//f.SetActiveSheet(index)
	// Save spreadsheet by the given path.

}
