package invoicer

import (
    "github.com/go-pdf/fpdf"
    "github.com/leekchan/accounting"
)

type Document struct {
    pdf *fpdf.Fpdf
    ac  accounting.Accounting

    Config      *Config        `json:"config,omitempty"`
    Header       *HeaderFooter `json:"header,omitempty"`
    Footer       *HeaderFooter `json:"footer,omitempty"`
}

// returns the underlying *fpdf.Fpdf used to build document
func (doc *Document) Pdf() *fpdf.Fpdf {
    return doc.pdf
}

// SetUnicodeTranslator to use
// See https://pkg.go.dev/github.com/go-pdf/fpdf#UnicodeTranslator
func (doc *Document) SetUnicodeTranslator(fn UnicodeTranslateFunc) {
    doc.Config.UnicodeTranslateFunc = fn
}

// encodeString encodes the string using doc.Options.UnicodeTranslateFunc
func (doc *Document) encodeString(str string) string {
    return doc.Config.UnicodeTranslateFunc(str)
}
