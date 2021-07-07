package util

import (
	"errors"
	"fmt"
	"net/http"
)

type apiError struct {
	APIErr    error
	ErrorCode ErrorCode
	ErrorInfo string
}

type errorResponse struct {
	APIID        string    `json:"apiId"`
	ErrorCode    ErrorCode `json:"errorCode"`
	ErrorMessage string    `json:"errorMessage"`
}

func (e apiError) Error() string {
	return fmt.Sprintf("ErrorCode: %d, ErrorInfo: %s, Error: %s", e.ErrorCode, e.ErrorInfo, e.APIErr)
}

func Errorf(ec ErrorCode, ei string, format string, a ...interface{}) error {
	return &apiError{
		APIErr:    fmt.Errorf(format, a...),
		ErrorCode: ec,
		ErrorInfo: ei,
	}
}

func GetErrorCode(err error) ErrorCode {
	var apiError *apiError
	if errors.As(err, &apiError) {
		return apiError.ErrorCode
	}
	return UnknownError
}

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
	case ErrorCode00006:
		status = http.StatusUnauthorized
		res.ErrorMessage = "The specified cookie does not exist"

	// 独自APIエラー
	case ErrorCode10000:
		status = http.StatusBadRequest
		res.ErrorMessage = "Password does not match"
	case ErrorCode10001:
		status = http.StatusInternalServerError
		res.ErrorMessage = "Failed to generate hash value"
	case ErrorCode10002:
		status = http.StatusInternalServerError
		res.ErrorMessage = "Failed to insert data"
	case ErrorCode10003:
		status = http.StatusInternalServerError
		res.ErrorMessage = "Failed to retrieve data"
	case ErrorCode10004:
		status = http.StatusNotFound
		res.ErrorMessage = "User not found"
	case ErrorCode10005:
		status = http.StatusBadRequest
		res.ErrorMessage = "Incorrect password"
	case ErrorCode10006:
		status = http.StatusUnprocessableEntity
		res.ErrorMessage = "Failed to get token"
	case ErrorCode10007:
		status = http.StatusUnauthorized
		res.ErrorMessage = "Unauthorized"
	case ErrorCode10008:
		status = http.StatusInternalServerError
		res.ErrorMessage = "Failed to send to SMTP"
	case ErrorCode10009:
		status = http.StatusBadRequest
		res.ErrorMessage = "Invalid token"
	case ErrorCode10010:
		status = http.StatusInternalServerError
		res.ErrorMessage = "Failed to update password"
	default:
		status = http.StatusInternalServerError
		res.ErrorMessage = "Unknown error"
	}

	return res, status
}
