package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/ivanlemeshev/cv/internal/cv"
	"github.com/ivanlemeshev/cv/internal/pdf"
)

var (
	inputFile = flag.String(
		"file",
		"cv.yml",
		"Path to the YAML file containing the CV data",
	)
	outputFile = flag.String(
		"output",
		"cv.pdf",
		"Path to the generated PDF file",
	)
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

	err = pdf.Generate(cvData, *outputFile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
