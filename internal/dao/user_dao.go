package dao

import (
	"llmcloud/internal/model"

	"gorm.io/gorm"
)

type UserDao interface {
	CheckFieldExists(fied string, value interface{}) (bool, error)
	CreateUser(user *model.User) error
	GetUserByName(name string) (*model.User, error)
}

type userDao struct {
	db *gorm.DB
}

func (r *userDao) CheckFieldExists(field string, value interface{}) (bool, error) {
	var count int64
	err := r.db.Model(&model.User{}).Where(field+" = ?", value).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *userDao) CreateUser(user *model.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userDao) GetUserByName(name string) (*model.User, error) {
	var user model.User
	err := r.db.Model(&model.User{}).Where("username = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserDao(db *gorm.DB) UserDao {
	return &userDao{db: db}
}
