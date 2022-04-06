package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

func main() {
	excelFile, err := excelize.OpenFile("MedianIncome.xlsx")
	if err != nil {
		log.Fatalln(err)
	}
	all_rows, err := excelFile.GetRows("h08")
	if err != nil {
		log.Fatalln(err)
	}
	for number, row := range all_rows {
		if number < 10 {
			continue
		}
		if len(row) <= 1 {
			continue
		}
		fmt.Println(row[0], " \t: ", row[1])
	}
}
