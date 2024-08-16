package models

type Event struct {
	Id          any    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"slug"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	CreatedAt   string `json:"createdAt"`
}
