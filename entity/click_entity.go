package entity

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"study-service/common"
	"time"
)

type Click struct {
	Id        int64     `xorm:"id pk autoincr"`
	CreatedAt time.Time `xorm:"created" `
}

func (s *Click) Create(ctx context.Context) error {

	if rowsAffected, err := common.GetDB(ctx).Insert(s); err != nil {
		log.Errorln(err)
		return err
	} else if rowsAffected == 0 {
		msg := "가입이 되지 않았습니다. 시스템관리자에게 문의하여 주세요"
		log.Errorln(msg)
		return errors.New(msg)
	}
	return nil
}

func (Click) TableName() string {
	return "click"
}
