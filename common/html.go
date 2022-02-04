package common

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"html/template"
	"os"
)

func ParseHtmlTemplate(filePath string, data interface{}) (string, error) {
	log.Traceln("")

	t, err := template.ParseFiles(filePath)
	if err != nil {
		fmt.Println(err)

		return "", err
	}
	fmt.Println("t : ",t)
	//HTML 파일 생성
	fileName := "test.html"

	fileWriter , err := os.Create(fileName)
	fmt.Println("fileWriter : ",fileWriter)
	if err != nil {
		return "", err
	}
	if err := t.Execute(fileWriter, data); err != nil {
		return "", err
	}

	buff := new(bytes.Buffer)
	if err = t.Execute(buff, data); err != nil {
		fmt.Println(err)
		return "", err
	}

	return buff.String(), nil
}
