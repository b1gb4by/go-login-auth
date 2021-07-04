package model

type RegisterUserRequestParam struct {
	FirstName       string `validate:"required" json:"first_name"`
	LastName        string `validate:"required" json:"last_name"`
	Email           string `validate:"required" json:"email"`
	Password        string `validate:"required" json:"password"`
	ConfirmPassword string `validate:"required" json:"confirm_password"`
}

type LoginAuthenticationRequestParam struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
