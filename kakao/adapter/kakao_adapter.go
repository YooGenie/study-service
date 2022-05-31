package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"study-service/common"
	"study-service/config"
	"sync"

	log "github.com/sirupsen/logrus"
)

var (
	kakaoAdapterOnce     sync.Once
	kakaoAdapterInstance *kakaoAdapter
)

func KakaoAdapter() *kakaoAdapter {
	kakaoAdapterOnce.Do(func() {
		kakaoAdapterInstance = &kakaoAdapter{}
	})

	return kakaoAdapterInstance
}

type kakaoAdapter struct {
}


func (adpater kakaoAdapter) GetKakaoUserInfo(ctx context.Context, authorizeCode string) (map[string]interface{}, map[string]interface{}, string, error) {
	log.Traceln("")

	// 1.인증 코드로 사용 토큰을 받기
	// https://developers.kakao.com/docs/latest/ko/kakaologin/rest-api#request-token
	data := url.Values{}
	data.Add("grant_type", "authorization_code")
	data.Add("client_id", config.Config.Kakao.RestApiKey)
	data.Add("redirect_uri", config.Config.Kakao.RedirectURL)
	data.Add("code", authorizeCode)

	response, err := http.Post("https://kauth.kakao.com/oauth/token", "application/x-www-form-urlencoded;charset=utf-8", bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Error("KaKao Token Error : ", err)
		return nil, nil, "", err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, "",err
	}

	var userToken = map[string]interface{}{}
	if err = json.Unmarshal(b, &userToken); err != nil {
		return nil, nil,"", err
	}
	if userToken["access_token"] == nil {
		return nil, nil, "",errors.New("kakao api get error : access_token is null")
	}
	accessToken := userToken["access_token"].(string)
	// 2. 토큰으로 사용자 정보 요청
	// https://developers.kakao.com/docs/latest/ko/user-mgmt/rest-api#req-user-info
	userInfo, err := adpater.requestKakaoAPI(ctx, http.MethodPost, "https://kapi.kakao.com/v2/user/me", accessToken)
	if err != nil {
		log.Error("KaKao User Info Error : ", err)
		return nil, nil, "",err
	}
	userInfoLog, logErr := common.Struct2Json(userInfo)
	if logErr != nil {
		log.Info("Authorization_code: ", authorizeCode, ", Kakao User Info: ", userInfo)
	} else {
		log.Info("Authorization_code: ", authorizeCode, ", Kakao User Info: ", userInfoLog)
	}

	termInfo, err := adpater.requestKakaoAPI(ctx, http.MethodGet, "https://kapi.kakao.com/v1/user/service/terms", accessToken)
	if err != nil {
		log.Error("KaKao Term Info Error : ", err)
		return nil, nil, "",err
	}

	return userInfo, termInfo, accessToken, nil
}

func (adpater kakaoAdapter) Unlink(ctx context.Context, token string) error {
	log.Traceln("")
	unlinkInfo, err := adpater.requestKakaoAPI(ctx, http.MethodPost, "https://kapi.kakao.com/v1/user/unlink", token)
	if err != nil {
		return err
	}
	log.Info("Unlink Info : ", unlinkInfo)
	return nil
}

func (kakaoAdapter) requestKakaoAPI(ctx context.Context, httpMethod string, url string, accessToken string) (map[string]interface{}, error) {
	req, err := http.NewRequest(httpMethod, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}

	var response *http.Response
	response, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	var info = map[string]interface{}{}
	defer response.Body.Close()
	var b []byte
	b, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	} else {
		if err = json.Unmarshal(b, &info); err != nil {
			return nil, err
		}
	}

	return info, nil
}
