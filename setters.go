package invoicer

func (d *Document) SetHeader(header *HeaderFooter) *Document {
	d.Header = header
	return d
}
