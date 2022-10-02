package api

import (
	"dev11/internal/model"
	"dev11/internal/storage"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Store is an interface that represent store
type Store interface {
	CreateEvent(event *model.Event) error
	UpdateEvent(userID, eventID int, newEvent *model.Event) error
	DeleteEvent(userID, eventID int)
	GetEventsForWeek(date time.Time, userID int) ([]model.Event, error)
	GetEventsForDay(date time.Time, userID int) ([]model.Event, error)
	GetEventsForMonth(date time.Time, userID int) ([]model.Event, error)
}

// ResultResponse is a struct that contained json field
type ResultResponse struct {
	Result []model.Event `json:"result"`
}

// ErrorResponse is a struct that contained json field
type ErrorResponse struct {
	Err string `json:"error"`
}

// Handler is a struct that contained Store
type Handler struct {
	EventService Store
}

// NewHandler is a handler constructor
func NewHandler() *Handler {
	return &Handler{
		EventService: storage.NewEventStorage(),
	}
}

// Register is a function that registers all routes to mux
func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/create_event", h.CreateEvent)
	mux.HandleFunc("/update_event", h.UpdateEvent)
	mux.HandleFunc("/delete_event", h.DeleteEvent)
	mux.HandleFunc("/events_for_day", h.GetEventsForDay)
	mux.HandleFunc("/events_for_week", h.GetEventsForWeek)
	mux.HandleFunc("/events_for_month", h.GetEventsForMonth)
}

// CreateEvent is a function that gets request data and passes to the service for creating
func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	event, err := h.DecodeJSON(r)
	if err != nil {
		h.ErrorResponse(w, fmt.Errorf("error occured while decoding input value: %v", err), http.StatusBadRequest)
		return
	}

	err = h.EventService.CreateEvent(event)
	if err != nil {
		h.ErrorResponse(w, fmt.Errorf("error occured while : %v", err), http.StatusServiceUnavailable)
		return
	}

	h.ResultResponse(w, []model.Event{*event})
}

// DeleteEvent is a function that gets request data and passes to the service for deleting
func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	event, err := h.DecodeJSON(r)
	if err != nil {
		h.ErrorResponse(w, fmt.Errorf("error while decoding input value: %v", err), http.StatusBadRequest)
		return
	}

	h.EventService.DeleteEvent(event.UserID, event.EventID)
	h.ResultResponse(w, []model.Event{*event})
}

// UpdateEvent is a function that gets request data and passes to the service for updating
func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	event, err := h.DecodeJSON(r)
	if err != nil {
		h.ErrorResponse(w, fmt.Errorf("error occured while decoding input value: %v", err), http.StatusBadRequest)
		return
	}

	err = h.EventService.UpdateEvent(event.UserID, event.EventID, event)
	if err != nil {
		h.ErrorResponse(w, err, http.StatusServiceUnavailable)
		return
	}

	h.ResultResponse(w, []model.Event{*event})
}

// GetEventsForDay is a function that gets request for event for day and returns slice of events
func (h *Handler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")

	uID, err := strconv.Atoi(userID)
	if err != nil || uID < 1 {
		if uID < 1 {
			err = errors.New("userID should be positive")
		}

		h.ErrorResponse(w, err, http.StatusBadRequest)

		return
	}

	eventDate, err := h.ParseDate(date)
	if err != nil {
		h.ErrorResponse(w, err, http.StatusBadRequest)

		return
	}

	events, err := h.EventService.GetEventsForDay(eventDate, uID)
	if err != nil {
		h.ErrorResponse(w, err, http.StatusServiceUnavailable)

		return
	}

	h.ResultResponse(w, events)
}

// GetEventsForWeek is a function that gets request for event for week and returns slice of events
func (h *Handler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")

	uID, err := strconv.Atoi(userID)
	if err != nil || uID < 1 {
		if uID < 1 {
			err = errors.New("userID should be positive")
		}

		h.ErrorResponse(w, err, http.StatusBadRequest)

		return
	}

	eventDate, err := h.ParseDate(date)
	if err != nil {
		h.ErrorResponse(w, err, http.StatusBadRequest)

		return
	}

	events, err := h.EventService.GetEventsForWeek(eventDate, uID)
	if err != nil {
		h.ErrorResponse(w, err, http.StatusServiceUnavailable)

		return
	}

	h.ResultResponse(w, events)
}

// GetEventsForMonth is a function that gets request for event for month and returns slice of events
func (h *Handler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")

	uID, err := strconv.Atoi(userID)
	if err != nil || uID < 1 {
		if uID < 1 {
			err = errors.New("userID should be positive")
		}

		h.ErrorResponse(w, err, http.StatusBadRequest)

		return
	}

	eventDate, err := h.ParseDate(date)
	if err != nil {
		h.ErrorResponse(w, err, http.StatusBadRequest)

		return
	}

	events, err := h.EventService.GetEventsForMonth(eventDate, uID)
	if err != nil {
		h.ErrorResponse(w, err, http.StatusServiceUnavailable)

		return
	}

	h.ResultResponse(w, events)
}

// DecodeJSON is a function that decode json format from request and returns event
func (h *Handler) DecodeJSON(r *http.Request) (*model.Event, error) {
	var event model.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		return nil, err
	}

	if event.UserID < 1 || event.EventID < 1 {
		return nil, errors.New("eventID or userID should pe positive")
	}

	return &event, nil
}

// ResultResponse is a function that returns positive response
func (h *Handler) ResultResponse(w http.ResponseWriter, events []model.Event) {
	w.Header().Set("Content-Type", "application/json")

	result, _ := json.MarshalIndent(&ResultResponse{Result: events}, " ", "")
	_, err := w.Write(result)
	if err != nil {
		h.ErrorResponse(w, err, http.StatusInternalServerError)

		return
	}
}

// ErrorResponse is a function that response with error status
func (h *Handler) ErrorResponse(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")

	jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")

	http.Error(w, string(jsonErr), status)
}

// ParseDate is a function that parsing date from string
func (h *Handler) ParseDate(date string) (time.Time, error) {
	var (
		eventDate time.Time
		err       error
	)

	eventDate, err = time.Parse("2006-01-02T15:04", date)
	if err != nil {
		eventDate, err = time.Parse("2006-01-02", date)
		if err != nil {
			eventDate, err = time.Parse("2006-01-02T15:04:00Z", date)
			if err != nil {
				return time.Time{}, fmt.Errorf("date format: e.g. 2022-05-10T14:10 error: %v", err)
			}
		}
	}

	return eventDate, nil
}
