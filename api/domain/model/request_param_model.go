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

type ForgotRequestParam struct {
	Email string `validate:"required" json:"email"`
}

type ResetRequestParam struct {
	Password        string `validate:"required" json:"password"`
	ConfirmPassword string `validate:"required" json:"confirm_password"`
	Token           string `validate:"required" json:"token" gorm:"column:token"`
}
