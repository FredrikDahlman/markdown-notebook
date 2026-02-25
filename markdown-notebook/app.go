package main

import (
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetNotesDir() string {
	return GetNotesDir()
}

func (a *App) SetNotesDir(dir string) {
	SetNotesDir(dir)
}

func (a *App) ListNotes() ([]Note, error) {
	return ListNotes()
}

func (a *App) GetNote(path string) (Note, error) {
	return ReadNote(path)
}

func (a *App) SaveNote(note Note) error {
	return SaveNote(note)
}

func (a *App) CreateNote() (Note, error) {
	return CreateNote()
}

func (a *App) DeleteNote(path string) error {
	return DeleteNote(path)
}

func (a *App) SearchNotes(query string) ([]Note, error) {
	return SearchNotes(query)
}

func (a *App) GetAllTags() (map[string]int, error) {
	return GetAllTags()
}

func (a *App) GetNotesByTag(tag string) ([]Note, error) {
	return GetNotesByTag(tag)
}
