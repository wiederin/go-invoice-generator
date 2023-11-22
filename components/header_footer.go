package components

import (
    "fmt"

    "github.com/wiederin/go-invoicer/constants"

    "github.com/creasty/defaults"
    "github.com/go-pdf/fpdf"
)

type HeaderFooter struct {
    UseCustomFunc bool    `json:"-"`
    Text          string  `json:"text,omitempty"`
    FontSize      float64 `json:"font_size,omitempty" default:"7"`
    Pagination    bool    `json:"pagination,omitempty"`
}

type fnc func()

// ApplyFunc allow user to apply custom func
func (hf *HeaderFooter) ApplyFunc(pdf *fpdf.Fpdf, fn fnc) {
    pdf.SetHeaderFunc(fn)
}

func (hf *HeaderFooter) ApplyHeader(doc *Document) error {
    if err := defaults.Set(hf); err != nil {
        return err
    }

    if !hf.UseCustomFunc {
        doc.Pdf.SetHeaderFunc(func() {
            currentY := doc.Pdf.GetY()
            currentX := doc.Pdf.GetX()

            doc.Pdf.SetTopMargin(constants.HeaderMarginTop)
            doc.Pdf.SetY(constants.HeaderMarginTop)

            doc.Pdf.SetLeftMargin(constants.BaseMargin)
            doc.Pdf.SetRightMargin(constants.BaseMargin)

            // Parse Text as html (simple)
            doc.Pdf.SetFont(doc.Config.Font, "", hf.FontSize)
            _, lineHt := doc.Pdf.GetFontSize()
            html := doc.Pdf.HTMLBasicNew()
            html.Write(lineHt, doc.encodeString(hf.Text))

            // Apply pagination
            if !hf.Pagination {
                doc.Pdf.AliasNbPages("") // Will replace {nb} with total page count
                doc.Pdf.SetY(constants.HeaderMarginTop + 8)
                doc.Pdf.SetX(195)
                doc.Pdf.CellFormat(
                    10,
                    5,
                    doc.encodeString(fmt.Sprintf("Page %d/{nb}", doc.Pdf.PageNo())),
                    "0",
                    0,
                    "R",
                    false,
                    0,
                    "",
                )
            }

            doc.Pdf.SetY(currentY)
            doc.Pdf.SetX(currentX)
            doc.Pdf.SetMargins(constants.BaseMargin, constants.BaseMarginTop, constants.BaseMargin)
        })
    }

    return nil
}

func (hf *HeaderFooter) ApplyFooter(doc *Document) error {
    if err := defaults.Set(hf); err != nil {
        return err
    }

    if !hf.UseCustomFunc {
        doc.Pdf.SetFooterFunc(func() {
            currentY := doc.Pdf.GetY()
            currentX := doc.Pdf.GetX()

            doc.Pdf.SetTopMargin(constants.HeaderMarginTop)
            doc.Pdf.SetY(287 - constants.HeaderMarginTop)

            // Parse Text as html (simple)
            doc.Pdf.SetFont(doc.Config.Font, "", hf.FontSize)
            _, lineHt := doc.Pdf.GetFontSize()
            html := doc.Pdf.HTMLBasicNew()
            html.Write(lineHt, doc.encodeString(hf.Text))

            // Apply pagination
            if hf.Pagination {
                doc.Pdf.AliasNbPages("") // Will replace {nb} with total page count
                doc.Pdf.SetY(287 - constants.HeaderMarginTop - 8)
                doc.Pdf.SetX(195)
                doc.Pdf.CellFormat(
                    10,
                    5,
                    doc.encodeString(fmt.Sprintf("Page %d/{nb}", doc.Pdf.PageNo())),
                    "0",
                    0,
                    "R",
                    false,
                    0,
                    "",
                )
            }

            doc.Pdf.SetY(currentY)
            doc.Pdf.SetX(currentX)
            doc.Pdf.SetMargins(constants.BaseMargin, constants.BaseMarginTop, constants.BaseMargin)
        })
    }

    return nil
}
