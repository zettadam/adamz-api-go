package models

import (
	"time"
)

type (
	Post struct {
		Id           int64      `db:"id"`
		Title        string     `db:"title"`
		Slug         string     `db:"slug"`
		Abstract     *string    `db:"abstract"`
		Body         *string    `db:"body"`
		Significance int        `db:"significance"`
		Tags         []*string  `db:"tags"`
		PublishedAt  *time.Time `db:"published_at"`
		CreatedAt    time.Time  `db:"created_at"`
		UpdatedAt    *time.Time `db:"updated_at"`
	}

	PostRequest struct {
		Id           *int64     `json:"id"`
		Title        string     `json:"title"`
		Slug         string     `json:"slug"`
		Abstract     *string    `json:"abstract"`
		Body         *string    `json:"body"`
		Significance int        `json:"significance"`
		Tags         []*string  `json:"tags"`
		PublishedAt  *time.Time `json:"published_at"`
	}
)
