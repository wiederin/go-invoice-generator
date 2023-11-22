package components

func (d *Document) SetHeader(header *HeaderFooter) *Document {
    d.Header = header
    return d
}

func (d *Document) SetFooter(footer *HeaderFooter) *Document {
    d.Footer = footer
    return d
}

func (d *Document) SetCompany(company *Contact) *Document {
	d.Company = company
	return d
}

func (d *Document) SetCustomer(customer *Contact) *Document {
	d.Customer = customer
	return d
}

func (d *Document) SetDescription(desc string) *Document {
	d.Description = desc
	return d
}

func (d *Document) SetVersion(version string) *Document {
	d.Version = version
	return d
}
