package service

import (
	"context"
	"errors"
	"study-service/common"
	"study-service/config"
	"sync"
)

var (
	emailServiceOnce     sync.Once
	emailServiceInstance *emailService
)

func EmailService() *emailService {
	emailServiceOnce.Do(func() {
		emailServiceInstance = &emailService{}
	})

	return emailServiceInstance
}

type emailService struct {
}

type MailContent struct {
	Password       string //email open시 사용할 password
	Type           string
	DonationId     int64
	DocuNo         string
	RegistrationNo string
	DonorType      string
	Content        map[string]interface{}
}

func (emailService) SendMessage(ctx context.Context) (err error) {
	// 전송할 메일 메세지생성
	message := common.NewMessage(config.Config.Mail.Sender, config.Config.Mail.Receipt.Subject, true)

	mailTo := []string{"genie201207@gmail.com"} //받는사람
	mailCC := []string{}                        //참조
	mailBCC := []string{}                       //숨은참조

	// 메일 수신자 및 참조자 지정
	if len(mailTo) == 0 {
		return errors.New("Mail 수신자가 지정되어 있지 않습니다.")
	}
	message.To = mailTo

	if len(mailCC) > 0 {
		message.CC = mailCC
	}

	if len(mailBCC) > 0 {
		message.BCC = mailBCC
	}

	// 메일바디셋팅
	if err := message.SetMailBody(config.Config.Mail.Receipt.Path.MailBody, struct{}{}); err != nil {
		return err
	}

	// 메일전송
	return message.Send()

}
