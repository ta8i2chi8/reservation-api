package domain

import (
	"time"
)

// ReservationStatus 予約ステータス
type ReservationStatus string

const (
	StatusPending   ReservationStatus = "pending"
	StatusConfirmed ReservationStatus = "confirmed"
	StatusCancelled ReservationStatus = "cancelled"
)

// Reservation 予約エンティティ（集約ルート）
type Reservation struct {
	ID        uint              `json:"id" gorm:"primaryKey"`
	UserID    uint              `json:"user_id" gorm:"not null"`
	TimeSlot  *TimeSlot         `json:"time_slot" gorm:"embedded"`
	Status    ReservationStatus `json:"status" gorm:"not null;default:'pending'"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

// NewReservation 新規予約を作成
func NewReservation(userID uint, timeSlot *TimeSlot) (*Reservation, error) {
	if userID == 0 {
		return nil, ErrInvalidUser
	}
	if timeSlot == nil {
		return nil, ErrInvalidTimeSlot
	}

	return &Reservation{
		UserID:   userID,
		TimeSlot: timeSlot,
		Status:   StatusPending,
	}, nil
}

// Confirm 予約を確定
func (r *Reservation) Confirm() error {
	if r.Status != StatusPending {
		return ErrReservationNotPending
	}
	r.Status = StatusConfirmed
	return nil
}

// Cancel 予約をキャンセル
func (r *Reservation) Cancel() error {
	if r.Status == StatusCancelled {
		return ErrReservationAlreadyCancelled
	}
	r.Status = StatusCancelled
	return nil
}
