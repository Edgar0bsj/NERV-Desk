package service

import (
	"time"

	"github.com/edgar0bsj/nerv-desk/module/user/model"
	"github.com/edgar0bsj/nerv-desk/module/user/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserStorage
}

func New(repo repository.UserStorage) *UserService {
	return &UserService{repo}
}

// -- SaveUser
func (s *UserService) SaveUser(userDTO *UserRegisterDto) error {
	validation := validator.New()
	if err := validation.Struct(userDTO); err != nil {
		return err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userEntity := model.User{
		ID:            uuid.New().String(),
		Name:          userDTO.Name,
		Email:         userDTO.Email,
		Password_hash: string(passwordHash),
		Role:          userDTO.Role,
		Created_at:    time.Now(),
		Updated_at:    time.Now(),
	}

	if err := s.repo.Save(&userEntity); err != nil {
		return err
	}

	return nil
}

// -- FindAllUsers
func (s *UserService) FindAllUsers() ([]*model.User, error) {
	return s.repo.FindAll()
}

// -- FindByIdUser
func (s *UserService) FindByIdUser(id string) (*model.User, error) {
	return s.repo.FindByID(id)
}

// -- UpdateUser
func (s *UserService) UpdateUser(userDTO *UserUpdateDto) error {
	validation := validator.New()
	if err := validation.Struct(userDTO); err != nil {
		return err
	}

	user, err := s.repo.FindByID(userDTO.ID)
	if err != nil {
		return err
	}

	user.Name = userDTO.Name
	user.Email = userDTO.Email
	user.Role = userDTO.Role
	user.Updated_at = time.Now()

	if err := s.repo.Update(user); err != nil {
		return err
	}

	return nil
}

// -- DeleteUser
func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}

// -- UpdatePassword
func (s *UserService) UpdatePassword(id string, newPassword UserUpdatePassword) error {
	validation := validator.New()

	if err := validation.Struct(newPassword); err != nil {
		return err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.repo.UpdatePassword(id, string(passwordHash))
}
