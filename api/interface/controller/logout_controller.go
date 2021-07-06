package controller

import (
	"net/http"
	"time"
)

const logoutAPIID = "LA04"

type LogoutController struct{}

func NewLogoutController() *LogoutController {
	return &LogoutController{}
}

func (_ *LogoutController) Logout(w http.ResponseWriter, _ *http.Request) {
	defer panicErrorResponse(w, logoutAPIID)

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	responseJSON(w, http.StatusNoContent, "")
}
