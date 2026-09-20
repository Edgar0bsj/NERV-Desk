package service

import "github.com/edgar0bsj/nerv-desk/module/user/model"

type UserLoginDto struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6,max=25"`
}

type UserRegisterDto struct {
	Name     string         `validate:"required,max=50"`
	Email    string         `validate:"required,email"`
	Password string         `validate:"required,min=6,max=25"`
	Role     model.UserRole `validate:"required,oneof=USER ATTENDANT ADMIN"`
}

type UserUpdateDto struct {
	Name  string         `validate:"required,max=50"`
	Email string         `validate:"required,email"`
	Role  model.UserRole `validate:"required,oneof=USER ATTENDANT ADMIN"`
}

type UserPasswordChange struct {
	Password string `validate:"required,min=6,max=25"`
}
