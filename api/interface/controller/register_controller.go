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

const loginAuthenticationAPIID = "LA01"

type RegisterUserController struct {
	i interactor.RegisterUserInteractor
}

func NewRegisterUserController(i interactor.RegisterUserInteractor) *RegisterUserController {
	return &RegisterUserController{
		i: i,
	}
}

func (ctrl *RegisterUserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	defer panicErrorResponse(w, loginAuthenticationAPIID)

	logger := util.NewStdLogger()

	b, readErr := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if readErr != nil {
		e := util.Errorf(util.ErrorCode00002, "", "errorMessage : %w", readErr)
		logger.Errorf("%s", e)
		errorResponse, status := util.GetErrorResponse(loginAuthenticationAPIID, util.ErrorCode00002)
		responseJSON(w, status, errorResponse)
		return
	}

	var req model.RegisterUserRequestParam
	JSONErr := json.Unmarshal(b, &req)
	if JSONErr != nil {
		e := util.Errorf(util.ErrorCode00003, "", "errorMessage : %w", JSONErr)
		logger.Errorf("%s", e)
		errorResponse, status := util.GetErrorResponse(loginAuthenticationAPIID, util.ErrorCode00003)
		responseJSON(w, status, errorResponse)
		return
	}

	reqLog := logger.AddJSONKey("RequestParam", req)
	logger.Printf("%s", reqLog)

	if validErr := validation.RequestBodyValidate(&req); validErr != nil {
		j, _ := json.Marshal(req)
		e := util.Errorf(util.ErrorCode00004, string(j), "errorMessage : %w", validErr)
		logger.Errorf("%s", e)
		errorResponse, status := util.GetErrorResponse(loginAuthenticationAPIID, util.GetErrorCode(e))
		errorResponse.ErrorMessage += " - " + validErr.Error()
		responseJSON(w, status, errorResponse)
		return
	}

	if err := ctrl.i.RegisterUser(req); err != nil {
		logger.Errorf("%s", err)
		errorCode := util.GetErrorCode(err)
		errorResponse, status := util.GetErrorResponse(loginAuthenticationAPIID, errorCode)
		responseJSON(w, status, errorResponse)
		return
	}

	responseJSON(w, http.StatusNoContent, "")
}
