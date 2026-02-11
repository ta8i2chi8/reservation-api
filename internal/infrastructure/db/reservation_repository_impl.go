package db

import (
	"reservation-system/internal/domain"
	"reservation-system/internal/repository"

	"gorm.io/gorm"
)

type reservationRepositoryImpl struct {
	db *gorm.DB
}

// NewReservationRepository 予約リポジトリを実装
func NewReservationRepository() repository.ReservationRepository {
	return &reservationRepositoryImpl{
		db: GetDB(),
	}
}

func (r *reservationRepositoryImpl) Create(reservation *domain.Reservation) error {
	return r.db.Create(reservation).Error
}

func (r *reservationRepositoryImpl) FindByID(id uint) (*domain.Reservation, error) {
	var reservation domain.Reservation
	err := r.db.First(&reservation, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrReservationNotFound
		}
		return nil, err
	}
	return &reservation, nil
}

func (r *reservationRepositoryImpl) FindByUserID(userID uint) ([]*domain.Reservation, error) {
	var reservations []*domain.Reservation
	err := r.db.Where("user_id = ?", userID).Find(&reservations).Error
	return reservations, err
}

func (r *reservationRepositoryImpl) Update(reservation *domain.Reservation) error {
	return r.db.Save(reservation).Error
}

func (r *reservationRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.Reservation{}, id).Error
}

func (r *reservationRepositoryImpl) CountByDateAndTime(date string, startTime, endTime string) (int, error) {
	var count int64
	err := r.db.Model(&domain.Reservation{}).
		Where("DATE(created_at) = ? AND time_slot_start_time >= ? AND time_slot_end_time <= ?", date, startTime, endTime).
		Count(&count).Error
	return int(count), err
}
