package types

import "time"

type Event struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type EventHandlerData struct {
	UserId int   `json:"user_id"`
	Event  Event `json:"event"`
}

type UpdateEventHandlerData struct {
	UserId   int   `json:"user_id"`
	Event    Event `json:"event"`
	NewEvent Event `json:"new_event"`
}

type DateIntervalHandlerData struct {
	UserId    int       `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type EventResponse struct {
	Result Event `json:"result"`
}

type EventsListResponse struct {
	Result []Event `json:"result"`
}

type ErrorResponse struct {
	ErrorMsg string `json:"error"`
}
