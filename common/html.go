package common

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"html/template"
	"os"
)

// HTML 템블릿
func ParseHtmlTemplate(filePath string, data interface{}) (string, error) {
	log.Traceln("")

	t, err := template.ParseFiles(filePath)
	if err != nil {
		fmt.Println(err)

		return "", err
	}

	//HTML 파일 생성
	fileName := "test.html"

	fileWriter, err := os.Create(fileName)

	if err != nil {
		return "", err
	}
	if err := t.Execute(fileWriter, data); err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	buff := new(bytes.Buffer)
	if err = t.Execute(buff, data); err != nil {
		fmt.Println("여기 에러? ", err)
		return "", err
	}

	defer os.Remove(fileName)

	return buff.String(), nil
}
