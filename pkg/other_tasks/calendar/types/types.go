package types

import "time"

type Event struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type HandlerEventData struct {
	UserId int   `json:"user_id"`
	Event  Event `json:"event"`
}

type HandlerDateData struct {
	UserId int       `json:"user_id"`
	Date   time.Time `json:"date"`
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
