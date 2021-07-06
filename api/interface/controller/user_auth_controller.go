package controller

import (
	"api/usecase/interactor"
	"api/util"
	"net/http"
)

const userAuthenticationAPIID = "LA03"

type UserAuthenticationController struct {
	i interactor.UserAuthenticationInteractor
}

func NewUserAuthenticationController(i interactor.UserAuthenticationInteractor) *UserAuthenticationController {
	return &UserAuthenticationController{
		i: i,
	}
}

func (ctrl *UserAuthenticationController) UserAuthentication(w http.ResponseWriter, r *http.Request) {
	defer panicErrorResponse(w, userAuthenticationAPIID)

	logger := util.NewStdLogger()

	cookie, errNoCookie := r.Cookie("jwt")
	if errNoCookie != nil {
		e := util.Errorf(util.ErrorCode00006, "", "errorMessage : %w", errNoCookie)
		logger.Errorf("%s", e)
		errorResponse, status := util.GetErrorResponse(userAuthenticationAPIID, util.ErrorCode00006)
		responseJSON(w, status, errorResponse)
	}

	user, err := ctrl.i.UserAuthentication(cookie.Value)
	if err != nil {
		logger.Errorf("%s", err)
		errorCode := util.GetErrorCode(err)
		errorResponse, status := util.GetErrorResponse(userAuthenticationAPIID, errorCode)
		responseJSON(w, status, errorResponse)
		return
	}

	responseJSON(w, http.StatusOK, user)
}
