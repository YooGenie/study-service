package dto

import (
	"time"
)

type KaKaoMemberInformation struct {
	Id                    int64      `json:"id"`
	WithdrawAt            *time.Time `json:"withdrawAt"`
}



func (m KaKaoMemberInformation) IsWithdraw() bool {
	if m.WithdrawAt != nil {
		return true
	}
	return false
}


type MaskedMemberSummary struct {
	Id           int64  `xorm:"id" json:"id"`
	Nickname     string `xorm:"nickname" json:"nickname"`
	MemberMobile string `xorm:"-" json:"mobile"`
	Mobile       string `xorm:"mobile" json:"-"`
}
