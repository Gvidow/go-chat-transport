package errors

import (
	"fmt"
	"net/http"
)

type httpError struct {
	message string
	code    int
}

var _ error = (*httpError)(nil)

func NewHTTPError(mes string, code int) *httpError {
	return &httpError{
		message: mes,
		code:    code,
	}
}

func MakeHTTPError(res *http.Response, mes string) error {
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return NewHTTPError(mes, res.StatusCode)
	}
	return nil
}

func (h httpError) Error() string {
	return fmt.Sprintf("%s: http code response: %d", h.message, h.code)
}
