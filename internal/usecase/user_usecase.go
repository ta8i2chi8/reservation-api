package usecase

import (
	"reservation-system/internal/domain"
	"reservation-system/internal/infrastructure/db"
	"reservation-system/internal/infrastructure/jwt"
	"reservation-system/internal/repository"
)

// UserUseCase ユーザーユースケース
type UserUseCase struct {
	userRepo repository.UserRepository
}

// NewUserUseCase ユーザーユースケースを作成
func NewUserUseCase() *UserUseCase {
	return &UserUseCase{
		userRepo: db.NewUserRepository(),
	}
}

// CreateUserRequest ユーザー作成リクエスト
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// CreateUser ユーザーを作成
func (uc *UserUseCase) CreateUser(req *CreateUserRequest) (*domain.User, error) {
	exists, err := uc.userRepo.Exists(req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, domain.ErrDuplicateEmail
	}

	user, err := domain.NewUser(req.Email, req.Password, req.Name)
	if err != nil {
		return nil, err
	}

	err = uc.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser ユーザーを取得
func (uc *UserUseCase) GetUser(id uint) (*domain.User, error) {
	return uc.userRepo.FindByID(id)
}

// LoginRequest ログインリクエスト
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse ログインレスポンス
type LoginResponse struct {
	Token string       `json:"token"`
	User  *domain.User `json:"user"`
}

// Login ログイン
func (uc *UserUseCase) Login(req *LoginRequest) (*LoginResponse, error) {
	user, err := uc.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	if err := user.CheckPassword(req.Password); err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	token, err := jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: token,
		User: &domain.User{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil
}
