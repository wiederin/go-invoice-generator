package invoicer

import (
	"github.com/go-pdf/fpdf"
)

// Build pdf document from config
func (doc *Document) Build() (*fpdf.Fpdf, error) {
    // todo: validate document data

	// Build base doc
	doc.pdf.SetMargins(BaseMargin, BaseMarginTop, BaseMargin)
	doc.pdf.SetXY(10, 10)

    // todo: add invoice data

	if doc.Header != nil {
		if err := doc.Header.applyHeader(doc); err != nil {
			return nil, err
		}
	}


	if doc.Footer != nil {
		if err := doc.Footer.applyFooter(doc); err != nil {
			return nil, err
		}
	}

	return doc.pdf, nil
}
