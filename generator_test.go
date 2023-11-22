package invoicer

import (
    "github.com/wiederin/go-invoicer/components"
    "github.com/wiederin/go-invoicer/constants"

    "errors"
    "testing"
)

func TestInitWithInvalidType(t *testing.T) {
    _, err := Init(&components.Config{
        TextInvoiceType: "INVALID",
    })

    if errors.Is(err, ErrInvalidDocumentType) {
        return
    }

    t.Fatalf("expected ErrInvalidDocumentType, got %v", err)
}

func TestInit(t *testing.T) {
    doc, err := Init(&components.Config{
        TextInvoiceType:   constants.Invoice,
        TextRefTitle:      "Invoice No.",
        CurrencyPrecision: 2,
    })

    if err != nil {
        t.Fatalf("error generating pdf. got error %v", err)
    }

    doc.SetHeader(&components.HeaderFooter{
        Text:       "<center>Test header content.</center>",
        Pagination: true,
    })


    doc.SetFooter(&components.HeaderFooter{
        Text:       "<center>Test footer content</center>",
        Pagination: true,
    })


    pdf, err := Build(doc)
    if err != nil {
        t.Errorf(err.Error())
    }

    err = pdf.OutputFileAndClose("test_generator_init.pdf")

    if err != nil {
        t.Errorf(err.Error())
    }
}
