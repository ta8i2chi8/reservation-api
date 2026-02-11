package repository

import "reservation-system/internal/domain"

// ReservationRepository 予約リポジトリインターフェース
type ReservationRepository interface {
	Create(reservation *domain.Reservation) error
	FindByID(id uint) (*domain.Reservation, error)
	FindByUserID(userID uint) ([]*domain.Reservation, error)
	Update(reservation *domain.Reservation) error
	Delete(id uint) error
	CountByDateAndTime(date string, startTime, endTime string) (int, error)
}
