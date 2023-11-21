package invoicer

import (
	"bytes"
	"fmt"
	"time"

	"github.com/go-pdf/fpdf"
)

// Build pdf document from config 
func (doc *Document) Build() (*fpdf.Fpdf, error) {
    // todo: validate document data

	// Build base doc
	doc.pdf.SetMargins(BaseMargin, BaseMarginTop, BaseMargin)
	doc.pdf.SetXY(10, 10)

    // todo: add invoice data

	return doc.pdf, nil
}
