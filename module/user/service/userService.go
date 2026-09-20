package service

import (
	"github.com/edgar0bsj/nerv-desk/module/user/model"
	"github.com/edgar0bsj/nerv-desk/module/user/repository"
)

type UserService struct {
	repo repository.UserStorage
}

func New(repo repository.UserStorage) *UserService {
	return &UserService{repo}
}

func (s *UserService) FindByEmail(email string) (*model.User, error) {
	return s.repo.FindByEmail(email)
}
