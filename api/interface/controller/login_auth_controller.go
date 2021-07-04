package controller

import (
	"api/domain/model"
	"api/infrastructure/validation"
	"api/usecase/interactor"
	"api/util"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const loginAuthenticationAPIID = "LA02"

type LoginAuthenticationController struct {
	i interactor.LoginAuthenticationInteractor
}

func NewLoginAuthenticationController(i interactor.LoginAuthenticationInteractor) *LoginAuthenticationController {
	return &LoginAuthenticationController{
		i: i,
	}
}

func (ctrl *LoginAuthenticationController) LoginAuthentication(w http.ResponseWriter, r *http.Request) {
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

	var req model.LoginAuthenticationRequestParam
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

	user, token, err := ctrl.i.LoginAuthentication(req)
	if err != nil {
		logger.Errorf("%s", err)
		errorCode := util.GetErrorCode(err)
		errorResponse, status := util.GetErrorResponse(loginAuthenticationAPIID, errorCode)
		responseJSON(w, status, errorResponse)
		return
	}
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	responseJSON(w, http.StatusOK, user)
}
