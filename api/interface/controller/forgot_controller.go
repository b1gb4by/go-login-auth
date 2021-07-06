package controller

import (
	"api/domain/model"
	"api/infrastructure/validation"
	"api/usecase/interactor"
	"api/util"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const forgotAPIID = "LA05"

type ForgotController struct {
	i interactor.ForgotInteractor
}

func NewForgotController(i interactor.ForgotInteractor) *ForgotController {
	return &ForgotController{
		i: i,
	}
}

func (ctrl *ForgotController) Forgot(w http.ResponseWriter, r *http.Request) {
	defer panicErrorResponse(w, userRegisterAPIID)

	logger := util.NewStdLogger()

	b, readErr := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if readErr != nil {
		e := util.Errorf(util.ErrorCode00002, "", "errorMessage : %w", readErr)
		logger.Errorf("%s", e)
		errorResponse, status := util.GetErrorResponse(forgotAPIID, util.ErrorCode00002)
		responseJSON(w, status, errorResponse)
		return
	}

	var req model.ForgotRequestParam
	JSONErr := json.Unmarshal(b, &req)
	if JSONErr != nil {
		e := util.Errorf(util.ErrorCode00003, "", "errorMessage : %w", JSONErr)
		logger.Errorf("%s", e)
		errorResponse, status := util.GetErrorResponse(forgotAPIID, util.ErrorCode00003)
		responseJSON(w, status, errorResponse)
		return
	}

	reqLog := logger.AddJSONKey("RequestParam", req)
	logger.Printf("%s", reqLog)

	if validErr := validation.RequestBodyValidate(&req); validErr != nil {
		j, _ := json.Marshal(req)
		e := util.Errorf(util.ErrorCode00004, string(j), "errorMessage : %w", validErr)
		logger.Errorf("%s", e)
		errorResponse, status := util.GetErrorResponse(forgotAPIID, util.GetErrorCode(e))
		errorResponse.ErrorMessage += " - " + validErr.Error()
		responseJSON(w, status, errorResponse)
		return
	}

	if err := ctrl.i.Forgot(req); err != nil {
		logger.Errorf("%s", err)
		errorCode := util.GetErrorCode(err)
		errorResponse, status := util.GetErrorResponse(forgotAPIID, errorCode)
		responseJSON(w, status, errorResponse)
		return
	}

	responseJSON(w, http.StatusNoContent, "")
}
