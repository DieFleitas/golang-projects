package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

const (
	bannerHt = 94.0
	xIndent  = 40.0
)

func main() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v, height=%v\n", w, h)
	pdf.AddPage()

	//top maroon banner
	pdf.SetFillColor(10, 60, 79)
	pdf.Polygon([]gofpdf.PointType{
		{0, 0},
		{w, 0},
		{w, bannerHt},
		{0, bannerHt * 0.9},
	}, "F")
	pdf.Polygon([]gofpdf.PointType{
		{0, h},
		{0, h - (bannerHt * 0.2)},
		{w, h - (bannerHt * 0.1)},
		{w, h},
	}, "F")

	// banner - invoice
	pdf.SetFont("arial", "B", 40)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt := pdf.GetFontSize()
	pdf.Text(xIndent, bannerHt-(bannerHt/2.0)+lineHt/3.1, "Text")

	// banner - phone, email, github
	pdf.SetFont("arial", "", 12)
	_, lineHt = pdf.GetFontSize()
	pdf.MoveTo(w - xIndent - 2.0*124.0, (bannerHt - (lineHt*1.5*3.0))/2.0)
	pdf.MultiCell(124.0, lineHt*1.5, "(123) 4567-8910\ndiegofleitas98@outlook.es\ngithub.com/DieFleitas.com", gofpdf.BorderNone, gofpdf.AlignRight, false)

	// banner - address
	pdf.SetFont("arial", "", 12)
	_, lineHt = pdf.GetFontSize()
	pdf.MoveTo(w - xIndent -124.0, (bannerHt - (lineHt*1.5*3.0))/2.0)
	pdf.MultiCell(124.0, lineHt*1.5, "Calle falsa 123\nBuenos Aires, AR\n1768", gofpdf.BorderNone, gofpdf.AlignRight, false)

	// Grid
	drawGrid((pdf))

	err := pdf.OutputFileAndClose("p2.pdf")
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
		pdf.SetTextColor(200, 200, 200)

		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}

	for y := 0.0; y < h; y = y + (h / 20.0) {
		if y < bannerHt*0.9 {
			pdf.SetTextColor(200, 200, 200)
		} else {
			pdf.SetTextColor(80, 80, 80)
		}
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
}
