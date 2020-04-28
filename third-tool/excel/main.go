package main

import "github.com/360EntSecGroup-Skylar/excelize"

func main() {
	xlsx := excelize.NewFile()
	sheet := xlsx.NewSheet("sheet1")

	data := map[string]string{
		"B1": "语文",
		"C1": "数学",
		"D1": "英语",
		"E1": "理综",

		"A2": "啊俊",
		"A3": "小杰",
		"A4": "老王",

		"B2": "112",
		"C2": "115",
		"D2": "128",
		"E2": "255",

		"B3": "100",
		"C3": "90",
		"D3": "110",
		"E3": "200",

		"B4": "70",
		"C4": "140",
		"D4": "60",
		"E4": "265",
	}

	for k, v := range data {
		xlsx.SetCellValue("sheet1", k, v)
	}
	xlsx.SetActiveSheet(sheet)
	err := xlsx.SaveAs("./2020-04-28.xlsx")
	if err != nil {
		panic(err)
	}
}
