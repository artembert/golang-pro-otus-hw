package memorystorage

import (
	"testing"
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
	"github.com/stretchr/testify/require"
)

func getEvent(t *testing.T) domain.Event {
	t.Helper()

	return domain.Event{
		Title:        "Title",
		StartTime:    time.Now(),
		Duration:     time.Duration(1) * time.Hour,
		Description:  "Description",
		UserID:       "any-va10e",
		NotifyBefore: time.Duration(1) * time.Hour,
		Notified:     false,
	}
}

func TestStorage(t *testing.T) {
	t.Run("CreateEvent", func(t *testing.T) {
		s := New()
		event := getEvent(t)

		err := s.CreateEvent(event)
		require.NoError(t, err)

		var created domain.Event
		for k := range s.events {
			created = s.events[k]
		}
		event.ID = created.ID

		require.NoError(t, err)
		require.EqualValues(t, event, created)
	})

	t.Run("DeleteEvent", func(t *testing.T) {
		s := New()
		event := getEvent(t)

		err := s.CreateEvent(event)
		require.NoError(t, err)

		var created domain.Event
		for k := range s.events {
			created = s.events[k]
		}

		event.ID = created.ID
		err = s.DeleteEvent(event.ID)
		require.NoError(t, err)

		_, err = s.GetEventByID(event.ID)
		require.Equal(t, storage.ErrEventNotFound, err)
	})
}
