package constants

const (
    // Invoice define the "invoice" document type
    Invoice string = "INVOICE"

    // BaseMargin define base margin used in documents
    BaseMargin float64 = 10

    // BaseMarginTop define base margin top used in documents
    BaseMarginTop float64 = 20

    // HeaderMarginTop define base header margin top used in documents
    HeaderMarginTop float64 = 5

    // MaxPageHeight define the maximum height for a single page
    MaxPageHeight float64 = 260
)

var (
	// BaseTextFontSize define the base font size for text in document
	BaseTextFontSize float64 = 8

	// SmallTextFontSize define the small font size for text in document
	SmallTextFontSize float64 = 7

	// ExtraSmallTextFontSize define the extra small font size for text in document
	ExtraSmallTextFontSize float64 = 6

	// LargeTextFontSize define the large font size for text in document
	LargeTextFontSize float64 = 10
)

const (
	// ItemColNameOffset ...
	ItemColNameOffset float64 = 10

	// ItemColUnitPriceOffset ...
	ItemColUnitPriceOffset float64 = 80

	// ItemColQuantityOffset ...
	ItemColQuantityOffset float64 = 103

	// ItemColCommissionOffset ...
	ItemColCommissionOffset float64 = 125

	// ItemColTotalHTOffset ...
	ItemColTotalHTOffset float64 = 113

	// ItemColDiscountOffset ...
	ItemColDiscountOffset float64 = 140

	// ItemColTaxOffset ...
	ItemColTaxOffset float64 = 157

	// ItemColTotalTTCOffset ...
	ItemColTotalTTCOffset float64 = 175
)
