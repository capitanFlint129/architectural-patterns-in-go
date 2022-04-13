package types

import "time"

type Event struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type CreateEventData struct {
	UserId int   `json:"user_id"`
	Event  Event `json:"event"`
}

type CreateEventResponse struct {
	Result Event `json:"result"`
}

type UpdateEventData struct {
	UserId int   `json:"user_id"`
	Event  Event `json:"event"`
}

type UpdateEventResponse struct {
	Result Event `json:"result"`
}

type ErrorResponse struct {
	ErrorMsg string `json:"error"`
}
