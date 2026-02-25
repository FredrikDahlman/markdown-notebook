## 1. Project Setup

- [x] 1.1 Initialize Wails project with Svelte template
- [x] 1.2 Configure project structure (Go backend + frontend directories)
- [x] 1.3 Set up development build and run commands
- [x] 1.4 Verify empty shell builds to .app on macOS
- [x] 1.5 Add required Go dependencies (go-yaml, fsnotify)
- [x] 1.6 Add required frontend dependencies (CodeMirror, markdown-it, shiki)

## 2. Go Backend - Core Services

- [x] 2.1 Implement Note struct with Content, Title, Tags, Path fields
- [x] 2.2 Implement ListNotes() - scan directory, return note list
- [x] 2.3 Implement GetNote(id) - read single note with frontmatter
- [x] 2.4 Implement SaveNote(note) - write note to disk
- [x] 2.5 Implement DeleteNote(id) - move to trash
- [x] 2.6 Implement CreateNote() - create new note with template
- [x] 2.7 Implement frontmatter parsing (YAML → Note struct)
- [x] 2.8 Implement frontmatter serialization (Note struct → YAML)
- [x] 2.9 Set up notes directory configuration (default + custom)

## 3. Go Backend - Search & Indexing

- [x] 3.1 Implement GetAllTags() - extract unique tags from all notes
- [x] 3.2 Implement GetNotesByTag(tag) - filter notes by tag
- [x] 3.3 Implement SearchNotes(query) - substring search in filenames and frontmatter
- [ ] 3.4 Set up fsnotify file watcher for notes directory
- [ ] 3.5 Implement event bus for file changes (notify frontend)

## 4. Frontend - Layout & Components

- [x] 4.1 Create main layout with three-panel design (sidebar, editor, preview)
- [x] 4.2 Implement resizable split pane between editor and preview
- [x] 4.3 Build notes list sidebar component
- [x] 4.4 Build tag sidebar component with counts
- [x] 4.5 Build search input component

## 5. Frontend - Editor (CodeMirror)

- [x] 5.1 Integrate CodeMirror 6 with markdown mode
- [x] 5.2 Add syntax highlighting for markdown
- [x] 5.3 Implement theme (matching macOS light/dark)
- [x] 5.4 Wire up change events to trigger preview updates
- [x] 5.5 Add keyboard shortcuts (Cmd+S save, Cmd+F find)
- [x] 5.6 Implement undo/redo support

## 6. Frontend - Live Preview

- [x] 6.1 Integrate markdown-it for markdown rendering
- [x] 6.2 Integrate Shiki for code block syntax highlighting
- [ ] 6.3 Implement scroll sync between editor and preview
- [x] 6.4 Style preview to match rendered output

## 7. Frontend - State & Integration

- [x] 7.1 Create Wails bindings for Go backend methods
- [x] 7.2 Implement frontend state management (current note, notes list, tags)
- [x] 7.3 Implement auto-save with 500ms debounce
- [ ] 7.4 Wire up file watcher events to refresh note list
- [ ] 7.5 Implement external file change detection (prompt user)

## 8. Polish & Features

- [x] 8.1 Implement new note creation with frontmatter template
- [x] 8.2 Add note title from H1 or filename fallback
- [x] 8.3 Sort notes by last modified (newest first)
- [x] 8.4 Implement tag filtering (click tag to filter)
- [x] 8.5 Implement search results display
- [ ] 8.6 Add macOS menu bar integration
- [ ] 8.7 Handle window close/save prompts

## 9. Testing & Build

- [ ] 9.1 Test on macOS with sample notes
- [ ] 9.2 Test file watcher with external edits
- [ ] 9.3 Test search functionality
- [ ] 9.4 Test tag filtering
- [x] 9.5 Build production .app
- [ ] 9.6 Test production build