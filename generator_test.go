package invoicer

import (
    "errors"
    "testing"
)

func TestInitWithInvalidType(t *testing.T) {
    _, err := Init(&Config{
        TextTypeInvoice: "INVALID",
    })

    if errors.Is(err, ErrInvalidDocumentType) {
        return
    }

    t.Fatalf("expected ErrInvalidDocumentType, got %v", err)
}

func TestInit(t *testing.T) {
    doc, err := Init(&Config{
        TextTypeInvoice:   Invoice,
        TextRefTitle:      "Invoice No.",
        CurrencyPrecision: 2,
    })

    if err != nil {
        t.Fatalf("error generating pdf. got error %v", err)
    }

    doc.SetHeader(&HeaderFooter{
        Text:       "<center>Test header content.</center>",
        Pagination: true,
    })


    doc.SetFooter(&HeaderFooter{
        Text:       "<center>Test footer content</center>",
        Pagination: true,
    })


    pdf, err := doc.Build()
    if err != nil {
        t.Errorf(err.Error())
    }

    err = pdf.OutputFileAndClose("test_generator_init.pdf")

    if err != nil {
        t.Errorf(err.Error())
    }
}
