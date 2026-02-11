package repository

import "reservation-system/internal/domain"

// UserRepository ユーザーリポジトリインターフェース
type UserRepository interface {
	Create(user *domain.User) error
	FindByID(id uint) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id uint) error
	Exists(email string) (bool, error)
}
