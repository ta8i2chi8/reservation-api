package usecase

import (
	"reservation-system/internal/domain"
	"reservation-system/internal/infrastructure/db"
	"reservation-system/internal/infrastructure/jwt"
	"reservation-system/internal/repository"
)

type AuthUseCase struct {
	userRepo repository.UserRepository
}

func NewAuthUseCase() *AuthUseCase {
	return &AuthUseCase{
		userRepo: db.NewUserRepository(),
	}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type RegisterResponse struct {
	Token string       `json:"token"`
	User  *domain.User `json:"user"`
}

func (uc *AuthUseCase) Register(req *RegisterRequest) (*RegisterResponse, error) {
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

	token, err := jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &RegisterResponse{
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

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  *domain.User `json:"user"`
}

func (uc *AuthUseCase) Authenticate(req *AuthRequest) (*AuthResponse, error) {
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

	return &AuthResponse{
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

func (uc *AuthUseCase) ValidateToken(tokenString string) (*jwt.Claims, error) {
	return jwt.ValidateToken(tokenString)
}
