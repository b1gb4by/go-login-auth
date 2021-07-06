package model

type User struct {
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	Password  []byte `gorm:"column:password"`
}

type AcquisitionUser struct {
	ID        int    `json:"id" gorm:"column:id"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	Email     string `json:"email" gorm:"column:email"`
	Password  []byte `json:"-" gorm:"column:password"`
}
