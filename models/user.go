package models

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username" gorm:"type:varchar(255);unique;not null"`
	Email    string `json:"email" gorm:"type:varchar(255);unique;not null"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Address  string `json:"address"`
}
