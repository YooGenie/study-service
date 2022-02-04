package service

import (
	"fmt"
	"os"
	"study-service/common"
	responseDto "study-service/dto/response"
	"sync"
)

var (
	pdfServiceOnce     sync.Once
	pdfServiceInstance *pdfService
)

func PdfService() *pdfService {
	pdfServiceOnce.Do(func() {
		pdfServiceInstance = &pdfService{}
	})

	return pdfServiceInstance
}

type pdfService struct {
}

func (pdfService) MakeHtmlByte(id int64) (filePDFName []byte, err error) {

	dataHTML := &responseDto.PdfHTML{
		Id:        id,
		Type:     "안녕",
	}

	htmlGenerated, err := common.ParseHtmlTemplate("template/test.html", dataHTML)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer os.Remove(htmlGenerated)

	filePDFName, err = common.HtmlToPdfByte(htmlGenerated)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return filePDFName, err
}


func (pdfService) MakeHtmlString(id int64) (filePDFName string, err error) {

	dataHTML := &responseDto.PdfHTML{
		Id:        id,
		Type:     "안녕",
	}

	htmlGenerated, err := common.ParseHtmlTemplate("template/test.html", dataHTML)
	if err != nil {
		fmt.Println(err)
		return "", err
	}



	defer os.Remove(htmlGenerated)

	filePDFName, err = common.HtmlToPdfString(htmlGenerated)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return filePDFName, err
}
