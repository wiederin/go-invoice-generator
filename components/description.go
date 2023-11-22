package components

func (doc *Document) AddDescription() {
	if len(doc.Description) > 0 {
		doc.Pdf.SetY(doc.Pdf.GetY() + 10)
		doc.Pdf.SetFont(doc.Config.Font, "", 10)
		doc.Pdf.MultiCell(190, 5, doc.encodeString(doc.Description), "B", "L", false)
	}
}
