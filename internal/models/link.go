package models

import (
	"time"
)

type Link struct {
	Id           int64      `db:"id"`
	Url          string     `db:"url"`
	Title        string     `db:"title"`
	Description  *string    `db:"description"`
	Significance int        `db:"significance"`
	PublishedAt  *time.Time `db:"published_at"`
	Tags         []*string  `db:"tags"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
}

type LinkRequest struct {
	Id           *int64     `json:"id"`
	Url          string     `json:"url"`
	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	Significance int        `json:"significance"`
	PublishedAt  *time.Time `json:"published_at"`
	Tags         []*string  `json:"tags"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
