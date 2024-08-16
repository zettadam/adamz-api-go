package models

type Note struct {
	Id          any    `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Body        string `json:"body"`
	PublishedAt string `json:"publishedAt"`
	CreatedAt   string `json:"createdAt"`
}
