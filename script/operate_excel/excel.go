package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {

	fileName := "十四五规划商机任务清单-合并版本.xlsx"
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Printf("%+v\n", f.Sheet)
	return

}
