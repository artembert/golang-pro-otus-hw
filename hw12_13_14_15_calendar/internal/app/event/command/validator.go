package command

import (
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
)

type ErrValidate struct {
	message string
}

func (e ErrValidate) Error() string { return e.message }

func NewErrValidate(message string) ErrValidate {
	return ErrValidate{message: message}
}

func validateTitle(title string) error {
	if title == "" {
		return NewErrValidate("title is empty")
	}
	return nil
}

func validateStartTime(startTime time.Time) error {
	if startTime.Before(time.Now()) {
		return NewErrValidate("startTime date is in the past")
	}
	return nil
}

func validateDuration(duration time.Duration) error {
	if duration < 0 {
		return NewErrValidate("duration is negative")
	}
	return nil
}

func validateUserID(id domain.UserID) error {
	if string(id) == "" {
		return NewErrValidate("user ID is empty")
	}
	return nil
}

func validateID(id string) error {
	if id == "" {
		return NewErrValidate("ID is empty")
	}
	return nil
}
