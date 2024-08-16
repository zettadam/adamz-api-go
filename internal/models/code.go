package models

type CodeSnippet struct {
	Id          any    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Body        string `json:"body"`
	PublishedAt string `json:"publishedAt"`
	CreatedAt   string `json:"createdAt"`
}
