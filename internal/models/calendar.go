package models

import (
	"time"
)

type Event struct {
	Id          int64      `db:"id"`
	Title       string     `db:"title"`
	Description *string    `db:"slug"`
	StartTime   *time.Time `db:"start_time"`
	EndTime     *time.Time `db:"end_time"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

type EventRequest struct {
	Id          *int64     `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"slug"`
	StartTime   *time.Time `json:"start_time"`
	EndTime     *time.Time `json:"end_time"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
