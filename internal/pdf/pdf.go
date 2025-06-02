package pdf

import (
	"fmt"
	"os"
	"strings"

	"codeberg.org/go-pdf/fpdf"
	"codeberg.org/go-pdf/fpdf/contrib/barcode"
	"github.com/anvidev/rma-portal/internal/store"
)

func TicketLabelPDF(ticket *store.Ticket) (*os.File, func(), error) {
	f, err := os.CreateTemp("", "file.pdf")
	closeFunc := func() {
		os.Remove(f.Name())
		f.Close()
	}
	if err != nil {
		return nil, closeFunc, err
	}

	model := "Ikke angivet"
	if ticket.Model != nil {
		model = *ticket.Model
	}
	serialNumber := "Ikke angivet"
	if ticket.SerialNumber != nil {
		serialNumber = *ticket.SerialNumber
	}
	var categories []string
	for _, c := range ticket.Categories {
		categories = append(categories, store.Categories[c])
	}

	pdf := fpdf.New("P", "mm", "A6", "")

	left, _, right, _ := pdf.GetMargins()
	pageWidth, _ := pdf.GetPageSize()

	t := pdf.UnicodeTranslatorFromDescriptor("")

	pdf.SetTitle(fmt.Sprintf("Skancode A/S RMA - #%s", ticket.ID), false)
	pdf.SetAuthor("Skancode A/S", false)
	pdf.SetAutoPageBreak(false, 0)
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 15, "Skancode A/S - RMA Label", "1", 1, "CM", false, 0, "")
	pdf.SetFont("Arial", "B", 7)
	pdf.WriteAligned(0, 15, "RMA ID:", "")
	pdf.Ln(4)
	pdf.SetFont("Arial", "", 9)
	pdf.WriteAligned(0, 15, ticket.ID, "")
	pdf.Ln(7.75)
	pdf.SetFont("Arial", "B", 7)
	pdf.WriteAligned(0, 10, "Kategori", "")
	pdf.Ln(4)
	pdf.SetFont("Arial", "", 9)
	pdf.WriteAligned(0, 10, strings.Join(categories, ", "), "")
	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 7)
	pdf.WriteAligned(0, 10, "Model", "")
	pdf.Ln(4)
	pdf.SetFont("Arial", "", 9)
	pdf.WriteAligned(0, 10, model, "")
	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 7)
	pdf.WriteAligned(0, 10, "Serienr.", "")
	pdf.Ln(7.5)
	pdf.SetFont("Arial", "", 9)
	pdf.WriteAligned(0, 4, serialNumber, "")
	pdf.Line(10, 75.5, pageWidth-right, 75.5)

	key := barcode.RegisterPdf417(pdf, ticket.ID, 6, 5)
	barcode.Barcode(pdf, key, left, 82, 85, 20, false)

	pdf.Line(10, 108.5, pageWidth-right, 108.5)

	pdf.SetFont("Arial", "B", 10)
	pdf.Text(right, 118, t("Hvis fundet, returnér venligst til Skancode A/S"))
	pdf.SetFont("Arial", "", 10)
	pdf.Text(right, 124, "Skancode A/S")
	pdf.Text(right, 129, "Hejrevang 13, 1. t.v")
	pdf.Text(right, 134, t("3450, Allerød"))
	pdf.Text(right, 139, "Danmark")

	if err = pdf.Output(f); err != nil {
		return nil, closeFunc, err
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		return nil, closeFunc, err
	}

	return f, closeFunc, nil
}
