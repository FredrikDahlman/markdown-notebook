package main

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Note struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
	Path      string   `json:"path"`
	Created   string   `json:"created"`
	Modified  string   `json:"modified"`
}

type Frontmatter struct {
	Title    string   `yaml:"title"`
	Tags     []string `yaml:"tags"`
	Created  string   `yaml:"created"`
	Modified string   `yaml:"modified"`
}

var notesDir string

func init() {
	homeDir, _ := os.UserHomeDir()
	defaultDir := filepath.Join(homeDir, "Documents", "MarkdownNotebook")
	notesDir = defaultDir
	
	if _, err := os.Stat(defaultDir); os.IsNotExist(err) {
		os.MkdirAll(defaultDir, 0755)
	}
}

func GetNotesDir() string {
	return notesDir
}

func SetNotesDir(dir string) {
	notesDir = dir
}

func ListNotes() ([]Note, error) {
	var notes []Note

	entries, err := os.ReadDir(notesDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		notePath := filepath.Join(notesDir, entry.Name())
		note, err := ReadNote(notePath)
		if err != nil {
			continue
		}
		notes = append(notes, note)
	}

	// Sort by modified date, newest first
	for i := 0; i < len(notes)-1; i++ {
		for j := i + 1; j < len(notes); j++ {
			if notes[j].Modified > notes[i].Modified {
				notes[i], notes[j] = notes[j], notes[i]
			}
		}
	}

	return notes, nil
}

func ReadNote(path string) (Note, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return Note{}, err
	}

	note := Note{
		Path: path,
		ID:   filepath.Base(path),
	}

	// Check for frontmatter
	contentStr := string(content)
	if strings.HasPrefix(contentStr, "---") {
		parts := strings.SplitN(contentStr, "---", 3)
		if len(parts) >= 3 {
			var fm Frontmatter
			if err := yaml.Unmarshal([]byte(parts[1]), &fm); err == nil {
				note.Title = fm.Title
				note.Tags = fm.Tags
				note.Created = fm.Created
				note.Modified = fm.Modified
				note.Content = strings.TrimSpace(parts[2])
			}
		}
	} else {
		note.Content = contentStr
	}

	// Extract title from first H1 if not in frontmatter
	if note.Title == "" {
		lines := strings.Split(note.Content, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "# ") {
				note.Title = strings.TrimPrefix(line, "# ")
				break
			}
		}
	}

	// Fallback to filename
	if note.Title == "" {
		note.Title = strings.TrimSuffix(filepath.Base(path), ".md")
	}

	// Get file modification time
	if info, err := os.Stat(path); err == nil {
		note.Modified = info.ModTime().Format(time.RFC3339)
		if note.Created == "" {
			note.Created = info.ModTime().Format(time.RFC3339)
		}
	}

	return note, nil
}

func SaveNote(note Note) error {
	note.Modified = time.Now().Format(time.RFC3339)

	var content string
	if len(note.Tags) > 0 || note.Created != "" {
		fm := Frontmatter{
			Title:    note.Title,
			Tags:     note.Tags,
			Created:  note.Created,
			Modified: note.Modified,
		}
		fmBytes, _ := yaml.Marshal(fm)
		content = "---\n" + string(fmBytes) + "---\n" + note.Content
	} else {
		content = note.Content
	}

	return os.WriteFile(note.Path, []byte(content), 0644)
}

func CreateNote() (Note, error) {
	timestamp := time.Now().Format("20060102150405")
	filename := "note-" + timestamp + ".md"
	notePath := filepath.Join(notesDir, filename)

	created := time.Now().Format(time.RFC3339)
	content := `---
title: New Note
tags: []
created: ` + created + `
modified: ` + created + `
---

# New Note

`

	if err := os.WriteFile(notePath, []byte(content), 0644); err != nil {
		return Note{}, err
	}

	return ReadNote(notePath)
}

func DeleteNote(path string) error {
	return os.Remove(path)
}

func SearchNotes(query string) ([]Note, error) {
	allNotes, err := ListNotes()
	if err != nil {
		return nil, err
	}

	query = strings.ToLower(query)
	var results []Note

	for _, note := range allNotes {
		// Search in filename
		if strings.Contains(strings.ToLower(note.ID), query) {
			results = append(results, note)
			continue
		}
		// Search in title
		if strings.Contains(strings.ToLower(note.Title), query) {
			results = append(results, note)
			continue
		}
		// Search in tags
		for _, tag := range note.Tags {
			if strings.Contains(strings.ToLower(tag), query) {
				results = append(results, note)
				break
			}
		}
	}

	return results, nil
}

func GetAllTags() (map[string]int, error) {
	notes, err := ListNotes()
	if err != nil {
		return nil, err
	}

	tagCounts := make(map[string]int)
	for _, note := range notes {
		for _, tag := range note.Tags {
			tag = strings.ToLower(tag)
			tagCounts[tag]++
		}
	}

	return tagCounts, nil
}

func GetNotesByTag(tag string) ([]Note, error) {
	allNotes, err := ListNotes()
	if err != nil {
		return nil, err
	}

	tag = strings.ToLower(tag)
	var results []Note

	for _, note := range allNotes {
		for _, noteTag := range note.Tags {
			if strings.ToLower(noteTag) == tag {
				results = append(results, note)
				break
			}
		}
	}

	return results, nil
}
