package entity

import (
	"context"
	"study-service/common"
	"time"

	log "github.com/sirupsen/logrus"
)


type KaKaoMember struct {
	Id                    int64      `xorm:"id pk autoincr" `
	KakaoId               int64      `xorm:"kakao_id" `
	Nickname              string     `xorm:"nickname" `
	ProfileImage          string     `xorm:"profile_image" `
	Mobile                string     `xorm:"mobile"`
	Email                 string     `xorm:"email" `
	Gender                *string    `xorm:"gender" `
	AgeRange              *string    `xorm:"age_range" `
	CreatedAt             time.Time  `xorm:"created" `
	UpdatedAt             time.Time  `xorm:"updated" `
	WithdrawAt            *time.Time `xorm:"withdraw_at" `

}


func (KaKaoMember) TableName() string {
	return "members"
}

func (m *KaKaoMember) Get(ctx context.Context) (bool, error) {
	log.Traceln("")

	has, err := common.GetDB(ctx).Get(m)
	if err != nil {
		log.Errorln(err)
	}

	return has, err
}

func (m *KaKaoMember) Withdraw() {
	now := time.Now()
	m.WithdrawAt = &now
}

//func (m *KaKaoMember) UpdateMobile(updateParam requestDto.UpdateMemberParam) {
//	m.Mobile = common.SetEncrypt(updateParam.Mobile)
//}

func (m KaKaoMember) IsSignUp() bool {
	return m.Id != int64(0)
}

func (m KaKaoMember) IsActiveUser() bool {
	return m.WithdrawAt == nil
}

func (m KaKaoMember) HasMobileNumber() bool {
	return m.Mobile != ""
}


