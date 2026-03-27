package database

import "github.com/devfullcycle/goexpert/9-APIs/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emaild string) (*entity.User, error)
}
