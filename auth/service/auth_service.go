package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"study-service/common/errors"
	auth "study-service/config/authentication"
	requestDto "study-service/dto/request"
	responseDto "study-service/dto/response"
	"study-service/kakao/adapter"
	"study-service/kakao/entity"
	"study-service/kakao/mapper"
	"study-service/kakao/repository"
	"study-service/kakao/service"
	member "study-service/member/entity"
	memberService "study-service/member/service"
	"study-service/security"
	"sync"
)

var (
	authServiceOnce     sync.Once
	authServiceInstance *authService
)

func AuthService() *authService {
	authServiceOnce.Do(func() {
		authServiceInstance = &authService{}
	})
	return authServiceInstance
}

type authService struct {
}

const (
	SignUpTypeUnlink = "unlink"
	SignUpTypeNew    = "new"
)


func (this authService) NewMemberJwtToken(ctx context.Context, authorizeCode string, signUpType string) (responseDto.MemberJwtToken, error) {
	log.Traceln("")

	isSignUp, member, err := this.AuthWithKakao(ctx, authorizeCode, signUpType)
	if err != nil {
		return responseDto.MemberJwtToken{
			SignUpped:       false,
			Token:           "",
			ActiveUser:      false,
			HadMobileNumber: false,
		}, err
	}

	token, err := auth.CreateToken(member.Id, member.Nickname, "member")
	orgMember := false


	return responseDto.MemberJwtToken{
		SignUpped:       isSignUp,
		Token:           token,
		ActiveUser:      member.IsActiveUser(),
		HadMobileNumber: member.HasMobileNumber(),
		OrgMember:       orgMember,
	}, err
}


func (authService) AuthWithKakao(ctx context.Context, authorizeCode string, signUpType string) (isSignUp bool, member entity.KaKaoMember, err error) {
	userInfo, _, _, err := adapter.KakaoAdapter().GetKakaoUserInfo(ctx, authorizeCode)
	if err != nil {
		return
	}
	userKakaoId := int64(userInfo["id"].(float64))
	var findMember entity.KaKaoMember
	if signUpType == SignUpTypeNew {
		findMember, err = repository.KaKaoMemberRepository().FindByKakaoId(ctx, userKakaoId) //아이디있는지 체크
		if err != nil {
			return
		}
	} else {
		findMember, err = repository.KaKaoMemberRepository().FindByKakaoIdWithoutWithdraw(ctx, userKakaoId) //탈퇴했는지 체크
		if err != nil {
			return
		}
	}
	if findMember.IsSignUp() { // 이미 가입했으면 업데이트
		mapper.UpdateMemberForKakao(userInfo, &findMember)
		if err = service.MemberService().Update(ctx, &findMember); err != nil {
			return
		}

		isSignUp = false
		member = findMember
		return
	}

	// 신규가입
	newMember := mapper.NewMemberForKakao(userInfo)
	if _, err = service.MemberService().Create(ctx, &newMember); err != nil {
		return
	}
	isSignUp = true
	member = newMember
	return
}

//토큰 받기
//서비스는 post 보낸다
//카카오쪽에서 토큰을 발급한다.
//
//사용자 로그인 처리
//ID 토큰 유효성 검증을 서비스에서 한다.
//발급받은 토큰으로 사용자 정보 조회 서비스 회원 정보 확인 또는 가입 처리
//서비스에서 로그인 완료를 클라이언트에게 보낸다.

func (authService) UnlinkWithKakao(ctx context.Context, authorizeCode string) error {
	//카카오 정보를 가져온다
	kakaoUserInfo, _, token, err := adapter.KakaoAdapter().GetKakaoUserInfo(ctx, authorizeCode)

	member, err := service.MemberService().GetMemberByKakaoId(ctx, int64(kakaoUserInfo["id"].(float64)))
	if err = service.MemberService().Withdraw(ctx, member.Id); err != nil {
		return err
	}
	if err = adapter.KakaoAdapter().Unlink(ctx, token); err != nil {
		return err
	}
	return nil
}

func (authService) AuthWithSignIdPassword(ctx context.Context, signIn requestDto.AdminSignIn) (token security.JwtToken, err error) {
	memberEntity, err := memberService.MemberService().GetMemberById(ctx, signIn.Email)
	if err != nil {
		return
	}
	//비밀번호 유효성
	err = member.Member{}.ValidatePassword(signIn.Password)
	if err != nil {
		err = errors.ErrAuthentication
		return
	}

	token, err = security.JwtAuthentication{}.GenerateJwtToken(security.UserClaim{
		Id:    memberEntity.Email,
		Name:  "유지니",
		Roles: memberEntity.Role,
	})

	return
}
