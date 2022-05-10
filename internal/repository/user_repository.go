package repository

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nickname string
}

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	GetUser(userID string) (*User, error)
	ListUsers() ([]*User, error)
}

func NewUserRepository(db *gorm.DB) (UserRepository, error) {
	return &userRepository{db: db}, nil
}

func (ur userRepository) CreateUser(user *User) (*User, error) {
	if err := ur.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur userRepository) GetUser(userID string) (*User, error) {
	var user User
	if err := ur.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur userRepository) ListUsers() ([]*User, error) {
	var users []*User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
