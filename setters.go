package invoicer

func (d *Document) SetHeader(header *HeaderFooter) *Document {
    d.Header = header
    return d
}

func (d *Document) SetFooter(footer *HeaderFooter) *Document {
    d.Footer = footer
    return d
}
