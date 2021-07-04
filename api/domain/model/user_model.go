package model

type User struct {
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	Password  []byte `gorm:"column:password"`
}

type AcquisitionUser struct {
	ID        int    `gorm:"column:id"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	Password  []byte `gorm:"column:password"`
}
