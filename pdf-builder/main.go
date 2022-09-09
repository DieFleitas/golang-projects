package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v, height=%v\n", w, h)
	pdf.AddPage()

	// Text stuff
	pdf.MoveTo(0, 0)
	pdf.SetFont("arial", "B", 15)
	_, lineHt := pdf.GetFontSize()
	pdf.SetTextColor(255, 0, 0)
	pdf.Text(0, lineHt, "Hello, World")
	pdf.MoveTo(0, lineHt*2.0)

	pdf.SetFont("times", "", 16)
	pdf.SetTextColor(100, 100, 100)
	_, lineHt = pdf.GetFontSize()
	pdf.MultiCell(0, lineHt*1.5, "", gofpdf.BorderNone, gofpdf.AlignRight, false)

	// Shapes
	pdf.SetFillColor(0, 255, 0)
	pdf.SetDrawColor(0, 0, 255)
	pdf.Rect(10, 100, 100, 100, "FD")
	pdf.SetFillColor(100, 200, 200)
	pdf.Polygon([]gofpdf.PointType{
		{110, 250},
		{160, 300},
		{110, 350},
		{60, 300},
		{70, 230},
	}, "F")

	pdf.ImageOptions("images/jump.png", 275, 275, 92, 0, false, gofpdf.ImageOptions{
		ReadDpi: true,
	}, 0, "")

	// Grid
	drawGrid((pdf))

	err := pdf.OutputFileAndClose("p1.pdf")
	if err != nil {
		panic(err)
	}
}

func drawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()

	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)

	for x := 0.0; x < w; x = x + (w / 20.0) {
		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}

	for y := 0.0; y < h; y = y + (h / 20.0) {
		pdf.Line(0, y, w, h)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
}