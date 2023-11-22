package components

import (
    "github.com/go-pdf/fpdf"
    "github.com/leekchan/accounting"
)

type Document struct {
    // config
    Accounting   accounting.Accounting

    // fields
    Pdf          *fpdf.Fpdf
	  Date         string        `json:"date,omitempty"`
    Ref          string        `json:"ref,omitempty" validate:"required,min=1,max=32"`
	  Version      string        `json:"version,omitempty" validate:"max=32"`

    // components
    Config       *Config        `json:"config,omitempty"`
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
