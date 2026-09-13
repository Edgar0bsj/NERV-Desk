package repository

import (
	"github.com/edgar0bsj/nerv-desk/module/user/model"
	"gorm.io/gorm"
)

type UserMysqlRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserMysqlRepository {
	return &UserMysqlRepository{
		db: db,
	}
}

func (s *UserMysqlRepository) Save(user *model.User) error {

	result := s.db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *UserMysqlRepository) FindAll() ([]*model.User, error) {
	var users []*model.User

	result := s.db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (s *UserMysqlRepository) FindByID(id string) (*model.User, error) {
	var user model.User

	err := s.db.Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserMysqlRepository) Update(user *model.User) error {
	var userEntity *model.User

	err := s.db.Where("id = ?", user.ID).First(&userEntity).Error

	if err != nil {
		return err
	}

	s.db.Model(&userEntity).Updates(user)

	return nil

}

func (s *UserMysqlRepository) Delete(id string) error {
	var user *model.User

	err := s.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}

	s.db.Delete(&user)

	return nil
}

func (s *UserMysqlRepository) UpdatePassword(id, newPasswordHash string) error {
	var user model.User
	err := s.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}

	err = s.db.Model(&user).Update("Password_hash", newPasswordHash).Error
	if err != nil {
		return err
	}

	return nil
}
