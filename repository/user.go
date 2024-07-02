package repository

import (
	"errors"
	"tpm-tujuh/model"

	"gorm.io/gorm"
)

type UserPgRepo struct {
	DB *gorm.DB
}

func (r *UserPgRepo) Create(user *model.User) (int, error) {
	err := r.DB.Debug().Create(&user).Error

	if err != nil {
		err = errors.New("Create email failed")
	}
	return user.UserId, err
}

func (r *UserPgRepo) Get(user *model.User) (*model.User, error) {
	var responseUser model.User
	response := r.DB.Debug().Where(
		"email = ?",
		user.Email).First(&responseUser)

	err := response.Error
	if err != nil {
		return nil, err
	}

	if response.RowsAffected == 0 {
		err = errors.New("Get user failed")
	}

	return &responseUser, err
}
