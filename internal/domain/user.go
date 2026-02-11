package domain

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User ユーザーエンティティ
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"not null"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser 新規ユーザーを作成
func NewUser(email, password, name string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
	}, nil
}

// CheckPassword パスワードを検証
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
