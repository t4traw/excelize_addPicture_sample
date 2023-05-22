package main

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/xuri/excelize/v2"
)

func main() {
	xlsx, _ := excelize.OpenFile("template_1.xlsx")
	pos, _ := xlsx.SearchSheet("Sheet1", "[photo]")
	graphicOptions := excelize.GraphicOptions{
		AutoFit:         true,
		LockAspectRatio: true,
	}
	xlsx.AddPicture("Sheet1", pos[0], "400x300.jpg", &graphicOptions)
	xlsx.SaveAs("output.xlsx")
}
