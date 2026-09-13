package service

import "github.com/edgar0bsj/nerv-desk/module/user/model"

type UserServiceInterface interface {
	SaveUser(userDTO *UserRegisterDto) error
	FindAllUsers() ([]*model.User, error)
	FindByIdUser(id string) (*model.User, error)
	UpdateUser(userDTO *UserUpdateDto) error
	DeleteUser(id string) error
	UpdatePassword(id string, newPassword UserUpdatePassword) error
}
