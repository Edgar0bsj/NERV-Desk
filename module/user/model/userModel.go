package model

import (
	"errors"
	"time"
)

// ERROS
var (
	ErrUserNotFound   = errors.New("Usuario Não Encontrado")
	ErrInvalidUser    = errors.New("Dados do Usuario invalido")
	ErrSaveFailedUser = errors.New("Error ao Persistir dados do Usuario")
)

type UserRole string

const (
	RoleUser      UserRole = "USER"
	RoleAttendant UserRole = "ATTENDANT"
	RoleAdmin     UserRole = "ADMIN"
)

// STRUCK USER
type User struct {
	ID            string `gorm:"primaryKey"`
	Name          string
	Email         string `gorm:"unique"`
	Password_hash string
	Role          UserRole
	Created_at    time.Time
	Updated_at    time.Time
}
