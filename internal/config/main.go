package config

import (
	"github.com/zettadam/adamz-api-go/internal/stores"
)

type Application struct {
	CodeSnippetStore *stores.CodeSnippetStore
	EventStore       *stores.EventStore
	LinkStore        *stores.LinkStore
	NoteStore        *stores.NoteStore
	PostStore        *stores.PostStore
	TaskStore        *stores.TaskStore
}
