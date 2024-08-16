package models

type Link struct {
	Id          any    `json:"id"`
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishedAt string `json:"publishedAt"`
	CreatedAt   string `json:"createdAt"`
}
