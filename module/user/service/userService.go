package service

import (
	"time"

	"github.com/edgar0bsj/nerv-desk/module/user/model"
	"github.com/edgar0bsj/nerv-desk/module/user/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func (s *UserService) CreateUser(userCreateDto *UserRegisterDto) (*model.User, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(userCreateDto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:            uuid.New().String(),
		Name:          userCreateDto.Name,
		Email:         userCreateDto.Email,
		Password_hash: string(passHash),
		Role:          userCreateDto.Role,
		Created_at:    time.Now(),
		Updated_at:    time.Now(),
	}

	err = s.repo.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) FindAllUser() ([]*model.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) FindById(userId string) (*model.User, error) {
	return s.repo.FindByID(userId)
}

func (s *UserService) UserUpdate(user *model.User, newUser *UserUpdateDto) (*model.User, error) {

	user.Name = newUser.Name
	user.Email = newUser.Email
	user.Role = newUser.Role
	user.Updated_at = time.Now()

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil

}

func (s *UserService) DeleteUser(userId string) error {
	return s.repo.Delete(userId)
}

func (s *UserService) UserUpdatePasswod(user *model.User, passhash []byte) error {
	user.Password_hash = string(passhash)
	if err := s.repo.Update(user); err != nil {
		return err
	}

	return nil
}
