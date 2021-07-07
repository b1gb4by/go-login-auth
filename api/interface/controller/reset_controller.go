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

const resetAPIID = "LA06"

type ResetController struct {
	i interactor.ResetInteractor
}

func NewResetController(i interactor.ResetInteractor) *ResetController {
	return &ResetController{
		i: i,
	}
}

func (ctrl *ResetController) Reset(w http.ResponseWriter, r *http.Request) {
	defer panicErrorResponse(w, resetAPIID)

	logger := util.NewStdLogger()

	b, readErr := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if readErr != nil {
		e := util.Errorf(util.ErrorCode00002, "", "errorMessage : %w", readErr)
		logger.Errorf("%s", e)
		errorResponse, status := util.GetErrorResponse(resetAPIID, util.ErrorCode00002)
		responseJSON(w, status, errorResponse)
		return
	}

	var req model.ResetRequestParam
	JSONErr := json.Unmarshal(b, &req)
	if JSONErr != nil {
		e := util.Errorf(util.ErrorCode00003, "", "errorMessage : %w", JSONErr)
		logger.Errorf("%s", e)
		errorResponse, status := util.GetErrorResponse(resetAPIID, util.ErrorCode00003)
		responseJSON(w, status, errorResponse)
		return
	}

	reqLog := logger.AddJSONKey("RequestParam", req)
	logger.Printf("%s", reqLog)

	if validErr := validation.RequestBodyValidate(&req); validErr != nil {
		j, _ := json.Marshal(req)
		e := util.Errorf(util.ErrorCode00004, string(j), "errorMessage : %w", validErr)
		logger.Errorf("%s", e)
		errorResponse, status := util.GetErrorResponse(resetAPIID, util.GetErrorCode(e))
		errorResponse.ErrorMessage += " - " + validErr.Error()
		responseJSON(w, status, errorResponse)
		return
	}

	if err := ctrl.i.Reset(req); err != nil {
		logger.Errorf("%s", err)
		errorCode := util.GetErrorCode(err)
		errorResponse, status := util.GetErrorResponse(userRegisterAPIID, errorCode)
		responseJSON(w, status, errorResponse)
		return
	}

	responseJSON(w, http.StatusNoContent, "")

}
