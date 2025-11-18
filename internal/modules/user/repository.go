package user

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]User, error)
	Create(u *User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &userRepository{db}
}

func (r *userRepository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) Create(u *User) error {
	return r.db.Create(u).Error
}
