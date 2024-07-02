package service

import (
	"errors"
	"regexp"
	"tpm-tujuh/model"
	"tpm-tujuh/repository"
)

type UserService struct {
	UserPgRepo *repository.UserPgRepo
}

func (s *UserService) Create(user *model.User) (int, error) {

	if !validateEmail(user.Email) {
		return 0, errors.New("email Invalid")
	}
	return s.UserPgRepo.Create(user)
}

func (s *UserService) Get(user *model.User) (*model.User, error) {
	if !validateEmail(user.Email) {
		return nil, errors.New("email Invalid")
	}
	return s.UserPgRepo.Get(user)
}

func validateEmail(email string) bool {
	re := regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`)
	return re.MatchString(email)
}
