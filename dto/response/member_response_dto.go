package dto

import (
	"encoding/json"
)

type MemberSummary struct {
	Id       int64           `json:"id"`
	Email    string          `json:"email"`
	Password string          `json:"password"`
	Mobile   string          `json:"mobile"`
	Name     string          `json:"name"`
	Nickname string          `json:"nickname"`
	Role     string          `json:"role"`
	Created  json.RawMessage `json:"created"`
	Updated  json.RawMessage `json:"updated"`
}

type MemberJwtToken struct {
	SignUpped       bool
	Token           string
	ActiveUser      bool
	HadMobileNumber bool
	OrgMember       bool
}
