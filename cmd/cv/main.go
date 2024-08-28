package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/go-pdf/fpdf"
	"gopkg.in/yaml.v3"

	"github.com/ivanlemeshev/cv/internal/cv"
)

var (
	inputFile  = flag.String("file", "cv.yml", "Path to the YAML file containing the CV data")
	outputFile = flag.String("output", "cv.pdf", "Path to the generated PDF file")
)

func main() {
	flag.Parse()

	fmt.Println("Generating CV...")

	fmt.Println("Input file:", *inputFile)
	fmt.Println("Output file:", *outputFile)

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var cvData cv.CV

	if err := yaml.Unmarshal(data, &cvData); err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("First name: %s\n", cvData.FirstName)
	fmt.Printf("Last name: %s\n", cvData.LastName)

	const oritentation = "P" // Portrait
	const unit = "mm"        // Millimeters
	const size = "A4"        // A4 paper size

	pdf := fpdf.New(oritentation, unit, size, "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, cvData.FirstName+" "+cvData.LastName)

	if err := pdf.OutputFileAndClose(*outputFile); err != nil {
		log.Fatalf("error: %v", err)
	}
}
