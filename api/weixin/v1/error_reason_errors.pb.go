// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsWeixinCodeUnspecified(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_WEIXIN_CODE_UNSPECIFIED.String() && e.Code == 400
}

func ErrorWeixinCodeUnspecified(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_WEIXIN_CODE_UNSPECIFIED.String(), fmt.Sprintf(format, args...))
}

func IsWeixinCodeError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_WEIXIN_CODE_ERROR.String() && e.Code == 404
}

func ErrorWeixinCodeError(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_WEIXIN_CODE_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsWxCode2sessionError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_WX_CODE2SESSION_ERROR.String() && e.Code == 500
}

func ErrorWxCode2sessionError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_WX_CODE2SESSION_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsWxDecryptDataError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_WX_DECRYPT_DATA_ERROR.String() && e.Code == 500
}

func ErrorWxDecryptDataError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_WX_DECRYPT_DATA_ERROR.String(), fmt.Sprintf(format, args...))
}
