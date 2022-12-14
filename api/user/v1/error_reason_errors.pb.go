// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUserUnspecified(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_UNSPECIFIED.String() && e.Code == 400
}

func ErrorUserUnspecified(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_USER_UNSPECIFIED.String(), fmt.Sprintf(format, args...))
}

func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_NOT_FOUND.String() && e.Code == 404
}

func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsUserGetError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_GET_ERROR.String() && e.Code == 500
}

func ErrorUserGetError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_USER_GET_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsUserCreateError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_CREATE_ERROR.String() && e.Code == 500
}

func ErrorUserCreateError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_USER_CREATE_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsUserUpdateError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_UPDATE_ERROR.String() && e.Code == 500
}

func ErrorUserUpdateError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_USER_UPDATE_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsUserCacheNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_CACHE_NOT_FOUND.String() && e.Code == 401
}

func ErrorUserCacheNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_USER_CACHE_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsUserCacheUnmarshalError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_CACHE_UNMARSHAL_ERROR.String() && e.Code == 401
}

func ErrorUserCacheUnmarshalError(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_USER_CACHE_UNMARSHAL_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsUserCacheMarshalError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_CACHE_MARSHAL_ERROR.String() && e.Code == 500
}

func ErrorUserCacheMarshalError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_USER_CACHE_MARSHAL_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsUserCacheSetError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_CACHE_SET_ERROR.String() && e.Code == 500
}

func ErrorUserCacheSetError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_USER_CACHE_SET_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsAddressNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ADDRESS_NOT_FOUND.String() && e.Code == 404
}

func ErrorAddressNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_ADDRESS_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsAddressGetError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ADDRESS_GET_ERROR.String() && e.Code == 500
}

func ErrorAddressGetError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_ADDRESS_GET_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsAddressCreateError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ADDRESS_CREATE_ERROR.String() && e.Code == 500
}

func ErrorAddressCreateError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_ADDRESS_CREATE_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsAddressConflict(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ADDRESS_CONFLICT.String() && e.Code == 409
}

func ErrorAddressConflict(format string, args ...interface{}) *errors.Error {
	return errors.New(409, ErrorReason_ADDRESS_CONFLICT.String(), fmt.Sprintf(format, args...))
}

func IsAddressEditError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ADDRESS_EDIT_ERROR.String() && e.Code == 500
}

func ErrorAddressEditError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_ADDRESS_EDIT_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsAddressEditNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ADDRESS_EDIT_NOT_FOUND.String() && e.Code == 401
}

func ErrorAddressEditNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_ADDRESS_EDIT_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsAddressDeleteError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ADDRESS_DELETE_ERROR.String() && e.Code == 500
}

func ErrorAddressDeleteError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_ADDRESS_DELETE_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsAddressDeleteNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ADDRESS_DELETE_NOT_FOUND.String() && e.Code == 401
}

func ErrorAddressDeleteNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_ADDRESS_DELETE_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}
