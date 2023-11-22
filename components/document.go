package components

import (
    "github.com/go-pdf/fpdf"
    "github.com/leekchan/accounting"
)

type Document struct {
    Pdf *fpdf.Fpdf
    Accounting  accounting.Accounting

    Config      *Config        `json:"config,omitempty"`
    Header       *HeaderFooter `json:"header,omitempty"`
    Footer       *HeaderFooter `json:"footer,omitempty"`
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
