package components
import (
		"github.com/wiederin/go-invoicer/constants"

		"fmt"
		"time"
)
func (doc *Document) Meta() {
    refString := fmt.Sprintf("%s %s", doc.Config.TextRefTitle, doc.Ref)

    doc.Pdf.SetXY(120, constants.BaseMarginTop+11)
    doc.Pdf.SetFont(doc.Config.Font, "", 8)
    doc.Pdf.CellFormat(80, 4, doc.encodeString(refString), "0", 0, "R", false, 0, "")


    date := time.Now().Format("02/01/2006")
    if len(doc.Date) > 0 {
        date = doc.Date
    }
    dateString := fmt.Sprintf("%s: %s", doc.Config.TextDateTitle, date)
    doc.Pdf.SetXY(120, constants.BaseMarginTop+19)
    doc.Pdf.SetFont(doc.Config.Font, "", 8)
    doc.Pdf.CellFormat(80, 4, doc.encodeString(dateString), "0", 0, "R", false, 0, "")

    if len(doc.Version) > 0 {
        versionString := fmt.Sprintf("%s: %s", doc.Config.TextVersionTitle, doc.Version)
        doc.Pdf.SetXY(120, constants.BaseMarginTop+15)
        doc.Pdf.SetFont(doc.Config.Font, "", 8)
        doc.Pdf.CellFormat(80, 4, doc.encodeString(versionString), "0", 0, "R", false, 0, "")
    }
}
