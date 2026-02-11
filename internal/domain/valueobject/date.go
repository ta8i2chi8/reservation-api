package valueobject

import (
	"time"
)

type Date struct {
	value time.Time
}

func NewDate(t time.Time) (*Date, error) {
	if t.IsZero() {
		return nil, ErrInvalidDate
	}

	return &Date{value: t}, nil
}

func (d *Date) Value() time.Time {
	return d.value
}

func (d *Date) String() string {
	return d.value.Format("2006-01-02")
}

func (d *Date) Equals(other *Date) bool {
	if other == nil {
		return false
	}
	return d.value.Equal(other.value)
}

func (d *Date) IsBefore(other *Date) bool {
	if other == nil {
		return false
	}
	return d.value.Before(other.value)
}

func (d *Date) IsAfter(other *Date) bool {
	if other == nil {
		return false
	}
	return d.value.After(other.value)
}
