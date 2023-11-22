package components

import (
    "github.com/wiederin/go-invoicer/constants"

    "bytes"
    b64 "encoding/base64"
    "image"

    "github.com/go-pdf/fpdf"
)

type Contact struct {
    Name    string   `json:"name,omitempty" validate:"required,min=1,max=256"`
    Logo    []byte   `json:"logo,omitempty"`
    Address *Address `json:"address,omitempty"`

    // AddtionnalInfo to append after contact informations. You can use basic html here (bold, italic tags).
    AddtionnalInfo []string `json:"additional_info,omitempty"`
}

func (c *Contact) appendContactTODoc(
    x float64,
    y float64,
    fill bool,
    logoAlign string,
    doc *Document,
) float64 {
    doc.Pdf.SetXY(x, y)

    // Logo
    if c.Logo != nil {
        // Create filename
        fileName := b64.StdEncoding.EncodeToString([]byte(c.Name))

        // Create reader from logo bytes
        ioReader := bytes.NewReader(c.Logo)

        // Get image format
        _, format, _ := image.DecodeConfig(bytes.NewReader(c.Logo))

        // Register image in pdf
        imageInfo := doc.Pdf.RegisterImageOptionsReader(fileName, fpdf.ImageOptions{
            ImageType: format,
        }, ioReader)

        if imageInfo != nil {
            var imageOpt fpdf.ImageOptions
            imageOpt.ImageType = format
            doc.Pdf.ImageOptions(fileName, doc.Pdf.GetX(), y, 0, 30, false, imageOpt, 0, "")
            doc.Pdf.SetY(y + 30)
        }
    }

    // Name
    if fill {
        doc.Pdf.SetFillColor(
            doc.Config.GreyBgColor[0],
            doc.Config.GreyBgColor[1],
            doc.Config.GreyBgColor[2],
        )
    } else {
        doc.Pdf.SetFillColor(255, 255, 255)
    }

    // Reset x
    doc.Pdf.SetX(x)

    // Name rect
    doc.Pdf.Rect(x, doc.Pdf.GetY(), 70, 8, "F")

    // Set name
    doc.Pdf.SetFont(doc.Config.BoldFont, "B", 10)
    doc.Pdf.Cell(40, 8, doc.encodeString(c.Name))
    doc.Pdf.SetFont(doc.Config.Font, "", 10)

    if c.Address != nil {
        // Address rect
        var addrRectHeight float64 = 17

        if len(c.Address.Address2) > 0 {
            addrRectHeight = addrRectHeight + 5
        }

        if len(c.Address.Country) == 0 {
            addrRectHeight = addrRectHeight - 5
        }

        doc.Pdf.Rect(x, doc.Pdf.GetY()+9, 70, addrRectHeight, "F")

        // Set address
        doc.Pdf.SetFont(doc.Config.Font, "", 10)
        doc.Pdf.SetXY(x, doc.Pdf.GetY()+10)
        doc.Pdf.MultiCell(70, 5, doc.encodeString(c.Address.ToString()), "0", "L", false)
    }

    // Addtionnal info
    if c.AddtionnalInfo != nil {
        doc.Pdf.SetXY(x, doc.Pdf.GetY())
        doc.Pdf.SetFontSize(constants.SmallTextFontSize)
        doc.Pdf.SetXY(x, doc.Pdf.GetY()+2)

        for _, line := range c.AddtionnalInfo {
            doc.Pdf.SetXY(x, doc.Pdf.GetY())
            doc.Pdf.MultiCell(70, 3, doc.encodeString(line), "0", "L", false)
        }

        doc.Pdf.SetXY(x, doc.Pdf.GetY())
        doc.Pdf.SetFontSize(constants.BaseTextFontSize)
    }

    return doc.Pdf.GetY()
}

// appendCompanyContactToDoc append the company contact to the document
func (c *Contact) AppendCompanyContactToDoc(doc *Document) float64 {
    x, y, _, _ := doc.Pdf.GetMargins()
    return c.appendContactTODoc(x, y, true, "L", doc)
}

// appendCompanyContactToDocBottom append the company contact to the document
func (c *Contact) AppendCompanyContactToDocBankDetails(doc *Document) float64 {
    return c.appendContactTODoc(10, doc.Pdf.GetY(), true, "L", doc)
}

// appendCustomerContactToDoc append the customer contact to the document
func (c *Contact) AppendCustomerContactToDoc(doc *Document) float64 {
    return c.appendContactTODoc(130, constants.BaseMarginTop+25, true, "R", doc)
}
