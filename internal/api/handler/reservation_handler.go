package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"reservation-system/internal/domain"
	"reservation-system/internal/usecase"
	"reservation-system/pkg/response"
	"reservation-system/pkg/validator"
)

type ReservationHandler struct {
	reservationUseCase *usecase.ReservationUseCase
}

func NewReservationHandler() *ReservationHandler {
	return &ReservationHandler{
		reservationUseCase: usecase.NewReservationUseCase(),
	}
}

func (h *ReservationHandler) CreateReservation(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID    uint   `json:"user_id"`
		Date      string `json:"date"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
		Capacity  int    `json:"capacity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	v := validator.NewValidator()
	v.Required("user_id", strconv.Itoa(int(req.UserID))).
		Required("date", req.Date).
		Required("start_time", req.StartTime).
		Required("end_time", req.EndTime).
		Required("capacity", strconv.Itoa(req.Capacity))

	if v.HasErrors() {
		response.BadRequest(w, v.GetFirstError())
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		response.BadRequest(w, "Invalid date format. Use YYYY-MM-DD")
		return
	}

	timeSlot, err := domain.NewTimeSlot(date, req.StartTime, req.EndTime, req.Capacity)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	createReq := &usecase.CreateReservationRequest{
		UserID:   req.UserID,
		TimeSlot: timeSlot,
	}

	resp, err := h.reservationUseCase.CreateReservation(createReq)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			response.NotFound(w, "User not found")
		case domain.ErrCapacityExceeded:
			response.BadRequest(w, "Capacity exceeded")
		default:
			response.InternalServerError(w, "Failed to create reservation")
		}
		return
	}

	response.Created(w, resp)
}

func (h *ReservationHandler) GetReservation(w http.ResponseWriter, r *http.Request) {
	reservationIDStr := r.URL.Query().Get("id")
	if reservationIDStr == "" {
		response.BadRequest(w, "Reservation ID is required")
		return
	}

	reservationID, err := strconv.ParseUint(reservationIDStr, 10, 32)
	if err != nil {
		response.BadRequest(w, "Invalid reservation ID")
		return
	}

	reservation, err := h.reservationUseCase.GetReservation(uint(reservationID))
	if err != nil {
		if err == domain.ErrReservationNotFound {
			response.NotFound(w, "Reservation not found")
			return
		}
		response.InternalServerError(w, "Failed to get reservation")
		return
	}

	response.Success(w, reservation)
}

func (h *ReservationHandler) GetUserReservations(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		response.BadRequest(w, "User ID is required")
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		response.BadRequest(w, "Invalid user ID")
		return
	}

	reservations, err := h.reservationUseCase.GetUserReservations(uint(userID))
	if err != nil {
		response.InternalServerError(w, "Failed to get reservations")
		return
	}

	response.Success(w, reservations)
}

func (h *ReservationHandler) ConfirmReservation(w http.ResponseWriter, r *http.Request) {
	var req usecase.ConfirmReservationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	v := validator.NewValidator()
	v.Required("reservation_id", strconv.Itoa(int(req.ReservationID))).
		Required("user_id", strconv.Itoa(int(req.UserID)))

	if v.HasErrors() {
		response.BadRequest(w, v.GetFirstError())
		return
	}

	err := h.reservationUseCase.ConfirmReservation(&req)
	if err != nil {
		switch err {
		case domain.ErrReservationNotFound:
			response.NotFound(w, "Reservation not found")
		case domain.ErrUnauthorized:
			response.Forbidden(w, "Not authorized to confirm this reservation")
		case domain.ErrReservationNotPending:
			response.BadRequest(w, "Reservation is not pending")
		default:
			response.InternalServerError(w, "Failed to confirm reservation")
		}
		return
	}

	response.Success(w, map[string]string{"message": "Reservation confirmed"})
}

func (h *ReservationHandler) CancelReservation(w http.ResponseWriter, r *http.Request) {
	reservationIDStr := r.URL.Query().Get("reservation_id")
	userIDStr := r.URL.Query().Get("user_id")

	if reservationIDStr == "" || userIDStr == "" {
		response.BadRequest(w, "Reservation ID and User ID are required")
		return
	}

	reservationID, err := strconv.ParseUint(reservationIDStr, 10, 32)
	if err != nil {
		response.BadRequest(w, "Invalid reservation ID")
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		response.BadRequest(w, "Invalid user ID")
		return
	}

	err = h.reservationUseCase.CancelReservation(uint(reservationID), uint(userID))
	if err != nil {
		switch err {
		case domain.ErrReservationNotFound:
			response.NotFound(w, "Reservation not found")
		case domain.ErrUnauthorized:
			response.Forbidden(w, "Not authorized to cancel this reservation")
		default:
			response.InternalServerError(w, "Failed to cancel reservation")
		}
		return
	}

	response.Success(w, map[string]string{"message": "Reservation cancelled"})
}
