package controller

import (
	"fmt"
	"net/http"
	"study-service/common/errors"
	"study-service/config"
	requestDto "study-service/dto/request"
	service2 "study-service/service"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

const (
	SignUpTypeUnlink = "unlink"
	SignUpTypeNew    = "new"
)

type AuthController struct {
}

func (controller AuthController) Init(g *echo.Group) {
	g.POST("/login", controller.AuthAdminWithEmailAndPassword)
	g.GET("/kakao", controller.RedirectKakaoLoginPage)
	g.GET("", controller.AuthWithKakao)
}

func (AuthController) AuthAdminWithEmailAndPassword(ctx echo.Context) (err error) {
	var adminSignIn requestDto.AdminSignIn //아이디와 비밀번호
	if err = ctx.Bind(&adminSignIn); err != nil {
		return errors.ApiParamValidError(err)
	}

	if err = adminSignIn.Validate(ctx); err != nil {
		return err
	}

	jwtToken, err := service2.AuthService().AuthWithSignIdPassword(ctx.Request().Context(), adminSignIn)
	if err != nil {
		if err == errors.ErrAuthentication {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	refreshToken, err := ctx.Cookie("refreshToken")
	if err != nil || len(refreshToken.Value) == 0 {
		cookie := new(http.Cookie)
		cookie.Name = "refreshToken"
		cookie.Value = jwtToken.RefreshToken
		cookie.HttpOnly = true
		cookie.Path = "/"
		ctx.SetCookie(cookie)
	} else {
		refreshToken.Value = jwtToken.RefreshToken
		refreshToken.HttpOnly = true
		refreshToken.Path = "/"
		ctx.SetCookie(refreshToken)
	}

	result := map[string]string{}
	result["accessToken"] = jwtToken.AccessToken
	return ctx.JSON(http.StatusOK, result)

}

func (AuthController) RedirectKakaoLoginPage(ctx echo.Context) (err error) {
	//클라이언트에서 카카로 로그인 요청을 하면 여기 접속한다.
	state := requestDto.Unlink{}
	if err = ctx.Bind(&state); err != nil {
		return errors.ApiParamValidError(err)
	}
	var kakaoUrl string
	//서비스는 Kakao API에 get 한다.
	kakaoUrl = fmt.Sprintf("https://kauth.kakao.com/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code",
		config.Config.Kakao.RestApiKey, config.Config.Kakao.RedirectURL)
	//카카오에서는 클라이언트한테 카카오 계정 로그인 요청한다.
	//클라이언트는 카카오 계정으로 로그인한다.

	if state.State == SignUpTypeUnlink {
		return ctx.Redirect(http.StatusPermanentRedirect, kakaoUrl+"&state=unlink")
	} else if state.State == SignUpTypeNew { //탈퇴한 후 회원가입 눌렀을 때
		return ctx.Redirect(http.StatusPermanentRedirect, kakaoUrl+"&state=new")
	}

	//사용자가 처음 가입한 사람이면
	// 동의 화면을 출력해서 클라이언트에 보여준다.
	// 클라이언트는 동의하고 시작한다.
	// 카카오는 서비스에 Redirect URL로 인가 코드 전달한다.

	return ctx.Redirect(http.StatusPermanentRedirect, kakaoUrl)
}

func (AuthController) AuthWithKakao(ctx echo.Context) (err error) {
	log.Trace("")
	state := requestDto.Unlink{} //이미 가입을 했으면 빈칸으로 나온다.

	if err = ctx.Bind(&state); err != nil {
		return errors.ApiParamValidError(err)
	}
	authorizeCode := ctx.QueryParam("code") // URL 주소에서 code를 뽑아 온다.
	kakaoErr := ctx.QueryParam("error")     //URL 주소에서 error를 뽑아 온다.

	if state.State == SignUpTypeUnlink {
		if err = service2.AuthService().UnlinkWithKakao(ctx.Request().Context(), authorizeCode); err != nil {
			log.Error(err)
		}
		logoutUri := fmt.Sprintf("https://kauth.kakao.com/oauth/logout?client_id=%s&logout_redirect_uri=%s",
			config.Config.Kakao.RestApiKey, config.Config.Kakao.LogoutRedirectURL)
		return ctx.Redirect(http.StatusMovedPermanently, logoutUri)
	}
	loginUri := "http://localhost:3000/login"
	if len(kakaoErr) > 0 {
		return ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("%s?error=%s", loginUri, kakaoErr))
	}
	var redirectUri string

	memberJwtToken, err := service2.AuthService().NewMemberJwtToken(ctx.Request().Context(), authorizeCode, state.State)

	if err != nil {
		redirectUri = fmt.Sprintf("%s?error=%s", loginUri, err.Error())
	} else if memberJwtToken.SignUpped {
		redirectUri = fmt.Sprintf("%s?token=%s&isSignUp=true&isActiveUser=%t&hasMobileNumber=%t&orgMember=%t", loginUri, memberJwtToken.Token, memberJwtToken.ActiveUser, memberJwtToken.HadMobileNumber, memberJwtToken.OrgMember)
	} else {
		redirectUri = fmt.Sprintf("%s?token=%s&isSignUp=false&isActiveUser=%t&hasMobileNumber=%t&orgMember=%t", loginUri, memberJwtToken.Token, memberJwtToken.ActiveUser, memberJwtToken.HadMobileNumber, memberJwtToken.OrgMember)
	}
	return ctx.Redirect(http.StatusMovedPermanently, redirectUri)
}
