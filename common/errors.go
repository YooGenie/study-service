package common

import (
	"errors"
	"menu-service/dtos"
)

var (
	ErrNoResult                          = errors.New("no result")
	ErrNoDeletable                       = errors.New("no deletable")
	ErrAuthorization                     = errors.New("no authorization")
	ErrDuplicationRegister               = errors.New("duplication register")
	ErrNoFileUpload                      = errors.New("no file upload")
	ErrNoMatch                           = errors.New("no match")
	ErrUserYetApproval                   = errors.New("yet approval")
	ErrUserReject                        = errors.New("reject")
	ErrDuplicationBusinessNumberRegister = errors.New("duplication business registration number register")
	ErrShippingDestination               = errors.New("no shipping destination")
)

var (
	APIErrorNoResult                          = dtos.ApiError{Code: 10001, Message: "no result"}
	APIErrorNoDeletable                       = dtos.ApiError{Code: 10002, Message: "no deletable"}
	APIErrorAuthorization                     = dtos.ApiError{Code: 10003, Message: "no authorization"}
	APIErrorDuplicationRegister               = dtos.ApiError{Code: 10004, Message: "duplication register"}
	APIErrorFileUpload                        = dtos.ApiError{Code: 10005, Message: "no file upload"}
	APIErrorNoMatch                           = dtos.ApiError{Code: 10006, Message: "no match"}
	APIErrorUserYetApproval                   = dtos.ApiError{Code: 10007, Message: "yet approval"}
	APIErrorUserReject                        = dtos.ApiError{Code: 10008, Message: "reject"}
	APIErrorDuplicationBusinessNumberRegister = dtos.ApiError{Code: 10009, Message: "duplication business registration number register"}
	APIErrorShippingDestination               = dtos.ApiError{Code: 10010, Message: "no shipping destination"}
)
