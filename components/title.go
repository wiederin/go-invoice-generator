package components

import(
	"github.com/wiederin/go-invoicer/constants"
)

func (doc *Document) Title() {
	title := doc.Config.TextInvoiceType

	// Set x y
	doc.Pdf.SetXY(120, constants.BaseMarginTop)

	// Draw rect
	doc.Pdf.SetFillColor(doc.Config.DarkBgColor[0], doc.Config.DarkBgColor[1], doc.Config.DarkBgColor[2])
	doc.Pdf.Rect(120, constants.BaseMarginTop, 80, 10, "F")

	// Draw text
	doc.Pdf.SetFont(doc.Config.Font, "", 14)
	doc.Pdf.CellFormat(80, 10, doc.encodeString(title), "0", 0, "C", false, 0, "")
}
