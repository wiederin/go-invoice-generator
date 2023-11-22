package invoicer

import (
    "errors"

    "github.com/creasty/defaults"
    "github.com/go-pdf/fpdf"
    "github.com/leekchan/accounting"
)

var ErrInvalidDocumentType = errors.New("invalid document type")

// Init returns new document with provided types and defaults
func Init(config *Config) (*Document, error) {
    _ = defaults.Set(config)

    if config.TextTypeInvoice != Invoice {
        return nil, ErrInvalidDocumentType
    }

    doc := &Document{
        Config: config,
    }

    // Prepare pdf
    doc.pdf = fpdf.New("P", "mm", "A4", "")
    doc.Config.UnicodeTranslateFunc = doc.pdf.UnicodeTranslatorFromDescriptor("")
    // todo: add custom fonts
    //doc.pdf.AddFont("OpenSans", "", "OpenSans-Regular.ttf")

    // Prepare accounting
    doc.ac = accounting.Accounting{
        Symbol:    doc.Config.CurrencySymbol,
        Precision: doc.Config.CurrencyPrecision,
        Thousand:  doc.Config.CurrencyThousand,
        Decimal:   doc.Config.CurrencyDecimal,
    }

    return doc, nil
}
