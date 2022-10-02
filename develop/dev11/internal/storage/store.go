package storage

import (
	"dev11/internal/model"
	"errors"
	"fmt"
	"sync"
	"time"
)

// EventStorage is a struct that contains mutex and database
type EventStorage struct {
	sync.RWMutex
	database map[string]model.Event
}

// NewEventStorage is a EventStorage constructor
func NewEventStorage() *EventStorage {
	return &EventStorage{
		database: make(map[string]model.Event),
	}
}

// CreateEvent is a function that creating new event in data store
func (e *EventStorage) CreateEvent(event *model.Event) error {
	id := fmt.Sprintf("%d%d", event.UserID, event.EventID)

	e.Lock()

	if _, ok := e.database[id]; ok {
		e.Unlock()
		return errors.New("event with such id already exist")
	}

	e.database[id] = *event

	e.Unlock()

	return nil
}

// UpdateEvent is a function that updating event in data store
func (e *EventStorage) UpdateEvent(userID, eventID int, newEvent *model.Event) error {
	combinedID := fmt.Sprintf("%d%d", userID, eventID)

	e.Lock()

	if _, ok := e.database[combinedID]; !ok {
		e.Unlock()
		return fmt.Errorf("there is no event with id: %s", combinedID)
	}

	e.database[combinedID] = *newEvent

	e.Unlock()

	return nil
}

// DeleteEvent is a function that deleting event from data store
func (e *EventStorage) DeleteEvent(userID, eventID int) {
	id := fmt.Sprintf("%d%d", userID, eventID)

	e.Lock()

	delete(e.database, id)

	e.Unlock()
}

// GetEventsForWeek is a function that returns all events for current week
func (e *EventStorage) GetEventsForWeek(date time.Time, userID int) ([]model.Event, error) {
	var eventsForWeek []model.Event

	currYear, currWeek := date.ISOWeek()

	e.RLock()

	for _, event := range e.database {
		eventYear, eventWeek := event.Date.ISOWeek()
		time.Now().ISOWeek()

		if eventYear == currYear && eventWeek == currWeek && userID == event.UserID {
			eventsForWeek = append(eventsForWeek, event)
		}
	}

	e.RUnlock()

	return eventsForWeek, nil
}

// GetEventsForDay is a function that returns all events for current day
func (e *EventStorage) GetEventsForDay(date time.Time, userID int) ([]model.Event, error) {
	var eventsForDay []model.Event

	y, m, d := date.Date()

	e.RLock()

	for _, event := range e.database {
		eventY, eventM, eventD := event.Date.Date()

		if y == eventY && int(eventM) == int(m) && d == eventD && userID == event.UserID {
			eventsForDay = append(eventsForDay, event)
		}
	}

	e.RUnlock()

	return eventsForDay, nil
}

// GetEventsForMonth is a function that returns all events for current month
func (e *EventStorage) GetEventsForMonth(date time.Time, userID int) ([]model.Event, error) {
	var eventsForMonth []model.Event

	y, m, _ := date.Date()

	e.RLock()

	for _, event := range e.database {
		eventY, eventM, _ := event.Date.Date()

		if y == eventY && int(eventM) == int(m) && userID == event.UserID {
			eventsForMonth = append(eventsForMonth, event)
		}
	}

	e.RUnlock()

	return eventsForMonth, nil
}
