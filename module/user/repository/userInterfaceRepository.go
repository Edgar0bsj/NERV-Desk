package repository

import "github.com/edgar0bsj/nerv-desk/module/user/model"

type UserStorage interface {
	Save(user *model.User) error
	FindAll() ([]*model.User, error)
	FindByID(id string) (*model.User, error)
	Update(User *model.User) error
	Delete(id string) error
	FindByEmail(email string) (*model.User, error)
}
