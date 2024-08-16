package models

type Task struct {
	Id          any    `json:"id"`
	TaskId      any    `json:"taskId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	CreatedAt   string `json:"createdAt"`
}
