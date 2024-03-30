package exceptions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OkexAPIException struct {
	Code       string
	Message    string
	StatusCode int
	Response   *http.Response
	Request    *http.Request
}

func NewOkexAPIException(response *http.Response) *OkexAPIException {
	var res map[string]interface{}
	json.NewDecoder(response.Body).Decode(&res)

	code := "None"
	message := "System error"
	if val, ok := res["code"]; ok {
		code = val.(string)
	}
	if val, ok := res["msg"]; ok {
		message = val.(string)
	}

	return &OkexAPIException{
		Code:       code,
		Message:    message,
		StatusCode: response.StatusCode,
		Response:   response,
		Request:    response.Request,
	}
}

func (e *OkexAPIException) Error() string {
	return fmt.Sprintf("API Request Error(code=%s): %s", e.Code, e.Message)
}

type OkexRequestException struct {
	Message string
}

func NewOkexRequestException(message string) *OkexRequestException {
	return &OkexRequestException{Message: message}
}

func (e *OkexRequestException) Error() string {
	return fmt.Sprintf("OkexRequestException: %s", e.Message)
}

type OkexParamsException struct {
	Message string
}

func NewOkexParamsException(message string) *OkexParamsException {
	return &OkexParamsException{Message: message}
}

func (e *OkexParamsException) Error() string {
	return fmt.Sprintf("OkexParamsException: %s", e.Message)
}
