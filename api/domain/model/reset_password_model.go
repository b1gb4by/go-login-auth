package model

type ResetPassword struct {
	Email string `gorm:"column:email"`
	Token string `gorm:"column:token"`
}
