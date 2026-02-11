package db

import (
	"reservation-system/internal/domain"
	"reservation-system/internal/repository"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository ユーザーリポジトリを実装
func NewUserRepository() repository.UserRepository {
	return &userRepositoryImpl{
		db: GetDB(),
	}
}

func (r *userRepositoryImpl) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepositoryImpl) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *userRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}

func (r *userRepositoryImpl) Exists(email string) (bool, error) {
	var count int64
	err := r.db.Model(&domain.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
