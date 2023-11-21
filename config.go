package invoicer

// UnicodeTranslateFunc
type UnicodeTranslateFunc func(string) string

// Options for Document
type Config struct {
    // fonts
    Font     string `default:"Helvetica"`
	BoldFont string `default:"Helvetica"`

	UnicodeTranslateFunc UnicodeTranslateFunc
}
