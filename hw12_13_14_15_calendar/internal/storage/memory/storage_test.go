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

		id, err := s.CreateEvent(&event)
		require.NoError(t, err)

		var created *domain.Event = s.events[id]
		event.ID = id

		require.NoError(t, err)
		require.EqualValues(t, &event, created)
	})

	t.Run("DeleteEvent", func(t *testing.T) {
		s := New()
		event := getEvent(t)

		id, err := s.CreateEvent(&event)
		require.NoError(t, err)

		err = s.DeleteEvent(id)
		require.NoError(t, err)

		_, err = s.GetEventByID(event.ID)
		require.Equal(t, storage.ErrEventNotFound, err)
	})

	t.Run("GetEventsByDate", func(t *testing.T) {
		s := New()
		targetDay := time.Date(2007, 1, 2, 1, 30, 0, 0, time.Local)
		event1Jan2007 := getEvent(t)
		event1Jan2007.StartTime = time.Date(2007, 1, 1, 0, 0, 0, 0, time.Local)
		event2Jan2007 := getEvent(t)
		event2Jan2007.StartTime = targetDay
		event2Jan2007Midday := getEvent(t)
		event2Jan2007Midday.StartTime = time.Date(2007, 1, 2, 12, 0, 0, 0, time.Local)
		event3Jan2007 := getEvent(t)
		event3Jan2007.StartTime = time.Date(2007, 1, 3, 0, 0, 0, 0, time.Local)

		_, _ = s.CreateEvent(&event1Jan2007)
		_, _ = s.CreateEvent(&event2Jan2007)
		_, _ = s.CreateEvent(&event2Jan2007Midday)
		_, _ = s.CreateEvent(&event3Jan2007)

		events, err := s.GetEventsByDate(targetDay)
		eventsDates := make([]time.Time, 0)
		for _, evt := range events {
			eventsDates = append(eventsDates, evt.StartTime)
		}
		require.NoError(t, err)
		require.EqualValues(t, []time.Time{event2Jan2007.StartTime, event2Jan2007Midday.StartTime}, eventsDates)
	})

	t.Run("GetEventsByWeek", func(t *testing.T) {
		s := New()

		event1Jan2007 := getEvent(t)
		event1Jan2007.StartTime = time.Date(2007, 1, 1, 3, 0, 0, 0, time.Local)
		event12Jan2007 := getEvent(t)
		event12Jan2007.StartTime = time.Date(2007, 1, 12, 7, 0, 0, 0, time.Local)
		event13Jan2007 := getEvent(t)
		event13Jan2007.StartTime = time.Date(2007, 1, 13, 12, 0, 0, 0, time.Local)
		event20Jan2007 := getEvent(t)
		event20Jan2007.StartTime = time.Date(2007, 1, 20, 0, 0, 0, 0, time.Local)

		_, _ = s.CreateEvent(&event1Jan2007)
		_, _ = s.CreateEvent(&event12Jan2007)
		_, _ = s.CreateEvent(&event13Jan2007)
		_, _ = s.CreateEvent(&event20Jan2007)

		events, err := s.GetEventsByWeek(time.Date(2007, 1, 10, 0, 0, 0, 0, time.Local))
		eventsDates := make([]time.Time, 0)
		for _, evt := range events {
			eventsDates = append(eventsDates, evt.StartTime)
		}
		require.NoError(t, err)
		require.EqualValues(t, []time.Time{event12Jan2007.StartTime, event13Jan2007.StartTime}, eventsDates)
	})
}
