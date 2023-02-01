package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
)

func convertToPdf(images []string, outfile string) error {
	// Create a new PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")

	for _, imageFile := range images {
		// Open the input file
		in, err := os.Open(imageFile)
		if err != nil {
			return err
		}
		defer in.Close()

		// Decode the input image
		// img, _, err := image.Decode(in)
		// if err != nil {
		// 	return err
		// }

		// Add the image to the PDF
		pdf.AddPage()
		pdf.Image(imageFile, 50, 0, 100, 297, false, "", 0, "")
	}

	// Create the output file
	out, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the PDF to the output file
	return pdf.Output(out)
}

func main() {
	outfile := flag.String("o", "", "output file")
	flag.Parse()

	if *outfile == "" || len(flag.Args()) == 0 {
		fmt.Println("usage: image-to-pdf -o <output file> <input files>")
		os.Exit(1)
	}

	err := convertToPdf(flag.Args(), *outfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
