package components

// UnicodeTranslateFunc
type UnicodeTranslateFunc func(string) string

// Config for Document
type Config struct {
    // fonts
    Font     string `default:"Helvetica"`
    BoldFont string `default:"Helvetica"`

    // colors
    DarkBgColor   []int `default:"[212,212,212]" json:"dark_bg_color,omitempty"`

    // currency
    CurrencySymbol    string `default:"€ " json:"currency_symbol,omitempty"`
    CurrencyPrecision int    `default:"2" json:"currency_precision,omitempty"`
    CurrencyDecimal   string `default:"." json:"currency_decimal,omitempty"`
    CurrencyThousand  string `default:" " json:"currency_thousand,omitempty"`

    // text
    TextInvoiceType      string `default:"INVOICE" json:"text_type_invoice,omitempty"`
    TextRefTitle         string `default:"Ref." json:"text_ref_title,omitempty"`
	  TextDateTitle        string `default:"Date" json:"text_date_title,omitempty"`
    TextVersionTitle     string `default:"Version" json:"text_version_title,omitempty"`

    UnicodeTranslateFunc UnicodeTranslateFunc
}
