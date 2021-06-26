package util

import (
	"errors"
	"fmt"
	"net/http"
)

// apiError
// 独自定義エラー
type apiError struct {
	APIErr    error
	ErrorCode ErrorCode
	ErrorInfo string
}

// Error
// エラー内容を返却
func (e apiError) Error() string {
	return fmt.Sprintf("ErrorCode: %d, ErrorInfo: %s, Error: %s", e.ErrorCode, e.ErrorInfo, e.APIErr)
}

// Errorf
// エラーをラップして、独自エラーを形成.
func Errorf(ec ErrorCode, ei string, format string, a ...interface{}) error {
	return &apiError{
		APIErr:    fmt.Errorf(format, a...),
		ErrorCode: ec,
		ErrorInfo: ei,
	}
}

// GetErrorCode
// エラーコードを取得する.
func GetErrorCode(err error) ErrorCode {
	var apiError *apiError
	if errors.As(err, &apiError) {
		return apiError.ErrorCode
	}

	// エラーハンドリングが正しくできていればこのエラーは返却しない
	return UnknownError
}

// errorResponse
// エラーレスポンス
type errorResponse struct {
	APIID        string    `json:"apiId"`
	ErrorCode    ErrorCode `json:"errorCode"`
	ErrorMessage string    `json:"errorMessage"`
}

// GetErrorResponse
// エラーレスポンスを形成して返却.
func GetErrorResponse(apiID string, ec ErrorCode) (errorResponse, int) {
	var status int
	res := errorResponse{
		APIID:        apiID,
		ErrorCode:    ec,
		ErrorMessage: "",
	}

	switch ec {
	// 共通エラー
	case ErrorCode00001:
		status = http.StatusInternalServerError
		res.ErrorMessage = "Unexpected error"
	case ErrorCode00002:
		status = http.StatusBadRequest
		res.ErrorMessage = "Request body can't read"
	case ErrorCode00003:
		status = http.StatusBadRequest
		res.ErrorMessage = "Request body json parse error"
	case ErrorCode00004:
		status = http.StatusBadRequest
		res.ErrorMessage = "Validation error"
	case ErrorCode00005:
		status = http.StatusServiceUnavailable
		res.ErrorMessage = "DB connection error"

	// 独自APIエラー
	case ErrorCode10000:
		status = http.StatusBadRequest
		res.ErrorMessage = "Failed to create a new request"
	default:
		status = http.StatusInternalServerError
		res.ErrorMessage = "Unknown error"
	}

	return res, status
}
