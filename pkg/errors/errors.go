package errors

import (
	"errors"
	"net/http"
)

type MyError struct {
	message string `json:"message"`	 // 对外提供
	inner string `json:"interMessage"` // 对内定位使用
	code int `json:"code"`  // HTTP Code
}

func (m *MyError)Error() string {
	return m.message
}

func (m *MyError)Message() string {
	return m.message
}

func (m *MyError)Inner() string {
	return m.inner
}

func (m *MyError)Code() int{
	return m.code
}

// http code 500
func New(message string) error {
	return &MyError{
		message: message,
		inner:  message,
		code: http.StatusInternalServerError,
	}
}

// code is 500,
func NewInner(message, inner string) error {
	return &MyError{
		message: message,
		inner:  inner,
		code: http.StatusInternalServerError,
	}
}

func NewCode(message, inner string, code int) error {
	return &MyError{
		message: message,
		inner:  inner,
		code: code,
	}
}

func Is(err, target error) bool{
	return errors.Is(err, target)
}

func As(err error, target interface{}) bool{
	return errors.As(err, target)
}

func Wrap() {
}

func UnWrap() {
}

var (
	NotFound = "数据不存在"
	ErrNotFound = NewCode(NotFound, NotFound, http.StatusOK)
)

// 数据未找到，返回200 并返回提示信息
func DataNotFound() error {
	return ErrNotFound
}
