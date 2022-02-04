package common

import (
	"strings"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"

	log "github.com/sirupsen/logrus"
)

func HtmlToPdfByte(htmlStr string) ([]byte, error) {
	log.Traceln("")

	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(htmlStr)))
	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fileName := "./genie.pdf"
	//Your Pdf Name
	err = pdfg.WriteFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return pdfg.Bytes(), nil
}


func HtmlToPdfString(htmlStr string) (string, error) {
	log.Traceln("")

	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		log.Println(err)
		return "", err
	}

	pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(htmlStr)))
	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Println(err)
		return "", err
	}

	fileName := "./genie.pdf"
	//Your Pdf Name

	err = pdfg.WriteFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return fileName, nil
}
