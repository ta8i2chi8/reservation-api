package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"reservation-system/internal/domain"
	"reservation-system/internal/usecase"
	"reservation-system/pkg/response"
	"reservation-system/pkg/validator"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userUseCase: usecase.NewUserUseCase(),
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req usecase.CreateUserRequest
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

	user, err := h.userUseCase.CreateUser(&req)
	if err != nil {
		if err == domain.ErrDuplicateEmail {
			response.BadRequest(w, "Email already exists")
			return
		}
		response.InternalServerError(w, "Failed to create user")
		return
	}

	response.Created(w, user)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("id")
	if userIDStr == "" {
		response.BadRequest(w, "User ID is required")
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		response.BadRequest(w, "Invalid user ID")
		return
	}

	user, err := h.userUseCase.GetUser(uint(userID))
	if err != nil {
		if err == domain.ErrUserNotFound {
			response.NotFound(w, "User not found")
			return
		}
		response.InternalServerError(w, "Failed to get user")
		return
	}

	response.Success(w, user)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req usecase.LoginRequest
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

	resp, err := h.userUseCase.Login(&req)
	if err != nil {
		if err == domain.ErrInvalidCredentials {
			response.Unauthorized(w, "Invalid credentials")
			return
		}
		response.InternalServerError(w, "Failed to login")
		return
	}

	response.Success(w, resp)
}
