package repositories

import (
	"backEnd/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(email string) (models.User, error)
	CheckAuth(ID int) (models.User, error)
	GetAllUser() ([]models.User, error)
	GetUserById(ID int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUserById(ID int) (models.User, error) {
	var user models.User
	query := "SELECT * From users WHERE id = ?"
	err := r.db.Raw(query, ID).Scan(&user).Error
	return user, err
}

func (r *repository) GetAllUser() ([]models.User, error) {
	var users []models.User
	query := "SELECT * From users"
	err := r.db.Raw(query).Scan(&users).Error
	return users, err
}

func (r *repository) Login(username string) (models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE user_name = ?"
	err := r.db.Raw(query, username).First(&user).Error

	return user, err
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) CheckAuth(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}
