package invoicer

import (
    "errors"

    "github.com/wiederin/go-invoicer/components"
    "github.com/wiederin/go-invoicer/constants"

    "github.com/creasty/defaults"
    "github.com/go-pdf/fpdf"
    "github.com/leekchan/accounting"
)

var ErrInvalidDocumentType = errors.New("invalid document type")

// Init returns new document with provided types and defaults
func Init(config *components.Config) (*components.Document, error) {
    _ = defaults.Set(config)

    if config.TextInvoiceType != constants.Invoice {
        return nil, ErrInvalidDocumentType
    }

    doc := &components.Document{
        Config: config,
    }

    // Prepare pdf
    doc.Pdf = fpdf.New("P", "mm", "A4", "")
    doc.Config.UnicodeTranslateFunc = doc.Pdf.UnicodeTranslatorFromDescriptor("")
    // todo: add custom fonts
    //doc.pdf.AddFont("OpenSans", "", "OpenSans-Regular.ttf")

    // Prepare accounting
    doc.Accounting = accounting.Accounting{
        Symbol:    doc.Config.CurrencySymbol,
        Precision: doc.Config.CurrencyPrecision,
        Thousand:  doc.Config.CurrencyThousand,
        Decimal:   doc.Config.CurrencyDecimal,
    }

    return doc, nil
}
