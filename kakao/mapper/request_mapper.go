package mapper

import (
	"study-service/common"
	"study-service/kakao/entity"
	"time"
)

func NewMemberForKakao(userInfo map[string]interface{}) entity.KaKaoMember {
	now := time.Now()

	profileImage := ""
	nickname := ""
	email := ""
	var gender *string
	var ageRange *string
	phoneNumber := ""
	if properties, ok := userInfo["properties"]; ok {
		if properties.(map[string]interface{})["profile_image"] != nil {
			profileImage = properties.(map[string]interface{})["profile_image"].(string)
		}
		if properties.(map[string]interface{})["nickname"] != nil {
			nickname = properties.(map[string]interface{})["nickname"].(string)
		}
	}
	if kakaoAccount, ok := userInfo["kakao_account"]; ok {
		profile := kakaoAccount.(map[string]interface{})["profile"]
		if profile != nil {
			if profile.(map[string]interface{})["profile_image_url"] != nil {
				profileImage = profile.(map[string]interface{})["profile_image_url"].(string)
			}
			if profile.(map[string]interface{})["nickname"] != nil {
				nickname = profile.(map[string]interface{})["nickname"].(string)
			}
		}
		if kakaoAccount.(map[string]interface{})["gender"] != nil {
			genderValue := kakaoAccount.(map[string]interface{})["gender"].(string)
			gender = &genderValue
		}
		if kakaoAccount.(map[string]interface{})["age_range"] != nil {
			ageRangeValue := kakaoAccount.(map[string]interface{})["age_range"].(string)
			ageRange = &ageRangeValue
		}
		if kakaoAccount.(map[string]interface{})["email"] != nil {
			email = kakaoAccount.(map[string]interface{})["email"].(string)
		}
		if kakaoAccount.(map[string]interface{})["phone_number"] != nil {
			phoneNumber = kakaoAccount.(map[string]interface{})["phone_number"].(string)
		}
	}

	return entity.KaKaoMember{
		KakaoId:               int64(userInfo["id"].(float64)),
		Nickname:              nickname,
		ProfileImage:          profileImage,
		Email:                 email,
		Gender:                gender,
		AgeRange:              ageRange,
		Mobile:                common.SetEncrypt(phoneNumber),
		CreatedAt:             now,
		UpdatedAt:             now,
	}
}

func UpdateMemberForKakao(userInfo map[string]interface{}, m *entity.KaKaoMember) {
	profileImage := ""
	nickname := ""
	var gender *string
	var ageRange *string
	if kakaoAccount, ok := userInfo["kakao_account"]; ok {
		profile := kakaoAccount.(map[string]interface{})["profile"]
		if profile != nil {
			if profile.(map[string]interface{})["profile_image_url"] != nil {
				profileImage = profile.(map[string]interface{})["profile_image_url"].(string)
			}
			if profile.(map[string]interface{})["nickname"] != nil {
				nickname = profile.(map[string]interface{})["nickname"].(string)
			}
		}
		if kakaoAccount.(map[string]interface{})["gender"] != nil {
			genderValue := kakaoAccount.(map[string]interface{})["gender"].(string)
			gender = &genderValue
		}
		if kakaoAccount.(map[string]interface{})["age_range"] != nil {
			ageRangeValue := kakaoAccount.(map[string]interface{})["age_range"].(string)
			ageRange = &ageRangeValue
		}
	}
	m.Nickname = nickname
	m.ProfileImage = profileImage
	m.Gender = gender
	m.AgeRange = ageRange
	m.UpdatedAt = time.Now()
}
