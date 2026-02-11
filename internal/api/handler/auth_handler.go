package handler

import (
	"encoding/json"
	"net/http"

	"reservation-system/internal/domain"
	"reservation-system/internal/usecase"
	"reservation-system/pkg/response"
	"reservation-system/pkg/validator"
)

type AuthHandler struct {
	authUseCase *usecase.AuthUseCase
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authUseCase: usecase.NewAuthUseCase(),
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req usecase.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	v := validator.NewValidator()
	v.Required("email", req.Email).
		Email("email", req.Email).
		Required("password", req.Password).
		MinLength("password", req.Password, 6).
		Required("name", req.Name).
		MinLength("name", req.Name, 2)

	if v.HasErrors() {
		response.BadRequest(w, v.GetFirstError())
		return
	}

	resp, err := h.authUseCase.Register(&req)
	if err != nil {
		if err == domain.ErrDuplicateEmail {
			response.BadRequest(w, "Email already exists")
			return
		}
		response.InternalServerError(w, "Failed to register user")
		return
	}

	response.Created(w, resp)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req usecase.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	v := validator.NewValidator()
	v.Required("email", req.Email).
		Email("email", req.Email).
		Required("password", req.Password)

	if v.HasErrors() {
		response.BadRequest(w, v.GetFirstError())
		return
	}

	resp, err := h.authUseCase.Authenticate(&req)
	if err != nil {
		if err == domain.ErrInvalidCredentials {
			response.Unauthorized(w, "Invalid credentials")
			return
		}
		response.InternalServerError(w, "Failed to authenticate")
		return
	}

	response.Success(w, resp)
}

func (h *AuthHandler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	v := validator.NewValidator()
	v.Required("token", req.Token)

	if v.HasErrors() {
		response.BadRequest(w, v.GetFirstError())
		return
	}

	claims, err := h.authUseCase.ValidateToken(req.Token)
	if err != nil {
		response.Unauthorized(w, "Invalid token")
		return
	}

	response.Success(w, claims)
}
