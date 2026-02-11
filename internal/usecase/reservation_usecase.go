package usecase

import (
	"reservation-system/internal/domain"
	"reservation-system/internal/infrastructure/db"
	"reservation-system/internal/repository"
)

type ReservationUseCase struct {
	reservationRepo repository.ReservationRepository
	userRepo        repository.UserRepository
}

func NewReservationUseCase() *ReservationUseCase {
	return &ReservationUseCase{
		reservationRepo: db.NewReservationRepository(),
		userRepo:        db.NewUserRepository(),
	}
}

type CreateReservationRequest struct {
	UserID   uint             `json:"user_id"`
	TimeSlot *domain.TimeSlot `json:"time_slot"`
}

type CreateReservationResponse struct {
	Reservation *domain.Reservation `json:"reservation"`
}

func (uc *ReservationUseCase) CreateReservation(req *CreateReservationRequest) (*CreateReservationResponse, error) {
	_, err := uc.userRepo.FindByID(req.UserID)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}

	count, err := uc.reservationRepo.CountByDateAndTime(
		req.TimeSlot.Date.Format("2006-01-02"),
		req.TimeSlot.StartTime,
		req.TimeSlot.EndTime,
	)
	if err != nil {
		return nil, err
	}

	if !req.TimeSlot.IsAvailable(count) {
		return nil, domain.ErrCapacityExceeded
	}

	reservation, err := domain.NewReservation(req.UserID, req.TimeSlot)
	if err != nil {
		return nil, err
	}

	err = uc.reservationRepo.Create(reservation)
	if err != nil {
		return nil, err
	}

	return &CreateReservationResponse{
		Reservation: reservation,
	}, nil
}

func (uc *ReservationUseCase) GetReservation(id uint) (*domain.Reservation, error) {
	return uc.reservationRepo.FindByID(id)
}

func (uc *ReservationUseCase) GetUserReservations(userID uint) ([]*domain.Reservation, error) {
	return uc.reservationRepo.FindByUserID(userID)
}

type ConfirmReservationRequest struct {
	ReservationID uint `json:"reservation_id"`
	UserID        uint `json:"user_id"`
}

func (uc *ReservationUseCase) ConfirmReservation(req *ConfirmReservationRequest) error {
	reservation, err := uc.reservationRepo.FindByID(req.ReservationID)
	if err != nil {
		return err
	}

	if reservation.UserID != req.UserID {
		return domain.ErrUnauthorized
	}

	err = reservation.Confirm()
	if err != nil {
		return err
	}

	return uc.reservationRepo.Update(reservation)
}

func (uc *ReservationUseCase) CancelReservation(reservationID, userID uint) error {
	reservation, err := uc.reservationRepo.FindByID(reservationID)
	if err != nil {
		return err
	}

	if reservation.UserID != userID {
		return domain.ErrUnauthorized
	}

	err = reservation.Cancel()
	if err != nil {
		return err
	}

	return uc.reservationRepo.Update(reservation)
}
