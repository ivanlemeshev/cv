package pdf

import (
	"fmt"

	"github.com/go-pdf/fpdf"

	"github.com/ivanlemeshev/cv/internal/cv"
)

const (
	font         = "Arial"
	fontBold     = "B"
	fontSize     = 16
	oritentation = "P"  // Portrait
	pageSize     = "A4" // A4 paper size
	unit         = "mm" // Millimeters
)

func Generate(cvData cv.CV, outputFile string) error {
	pdfDocument := fpdf.New(oritentation, unit, pageSize, "")

	pdfDocument.AddPage()
	pdfDocument.SetFont(font, fontBold, fontSize)
	pdfDocument.Cell(40, 10, cvData.FirstName+" "+cvData.LastName)

	if err := pdfDocument.OutputFileAndClose(outputFile); err != nil {
		return fmt.Errorf("faile to write PDF file: %w", err)
	}

	return nil
}
