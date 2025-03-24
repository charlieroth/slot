package errs

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
)

type ErrCode struct {
	value int
}

func (ec ErrCode) Value() int {
	return ec.value
}

func (ec ErrCode) String() string {
	return codeNames[ec]
}

func (ec *ErrCode) UnmarshalText(data []byte) error {
	errName := string(data)
	v, exists := codeNumbers[errName]
	if !exists {
		return fmt.Errorf("err code %q does not exist", errName)
	}

	*ec = v
	return nil
}

func (ec ErrCode) MarshalText() ([]byte, error) {
	return []byte(ec.String()), nil
}

func (ec ErrCode) Equal(ec2 ErrCode) bool {
	return ec.value == ec2.value
}

// ================================

type Error struct {
	Code     ErrCode `json:"code"`
	Message  string  `json:"message"`
	FuncName string  `json:"-"`
	FileName string  `json:"-"`
}

func New(code ErrCode, err error) *Error {
	pc, filename, line, _ := runtime.Caller(1)
	return &Error{
		Code:     code,
		Message:  err.Error(),
		FuncName: runtime.FuncForPC(pc).Name(),
		FileName: fmt.Sprintf("%s:%d", filename, line),
	}
}

func Newf(code ErrCode, format string, v ...any) *Error {
	pc, filename, line, _ := runtime.Caller(1)
	return &Error{
		Code:     code,
		Message:  fmt.Sprintf(format, v...),
		FuncName: runtime.FuncForPC(pc).Name(),
		FileName: fmt.Sprintf("%s:%d", filename, line),
	}
}

func NewError(err error) *Error {
	var errsErr *Error
	if errors.As(err, &errsErr) {
		return errsErr
	}

	return New(Internal, err)
}

func (e *Error) Error() string {
	return e.Message
}

// ================================

type FieldError struct {
	Field string `json:"field"`
	Err   string `json:"error"`
}

type FieldErrors []FieldError

func (fe *FieldErrors) Add(field string, err error) {
	*fe = append(*fe, FieldError{
		Field: field,
		Err:   err.Error(),
	})
}

func (fe FieldErrors) ToError() *Error {
	return New(InvalidArgument, fe)
}

func (fe FieldErrors) Error() string {
	d, err := json.Marshal(fe)
	if err != nil {
		return err.Error()
	}
	return string(d)
}
