package controller

import (
	"api/util"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
)

type AppController struct {
	HealthCheck *HealthCheckController
}

func NewControllers(hc *HealthCheckController) *AppController {
	return &AppController{
		HealthCheck: hc,
	}
}

func responseJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	res, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if _, err := w.Write(res); err != nil {
			logger := util.NewStdLogger()
			logger.Errorf("%s", err)
		}
	}
}

func panicErrorResponse(w http.ResponseWriter, apiID string) {
	if panicErr := recover(); panicErr != nil {
		fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		errorResponse, status := util.GetErrorResponse(apiID, util.ErrorCode00001)
		responseJSON(w, status, errorResponse)
	}
}