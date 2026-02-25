## Why

Build a native macOS markdown notebook application using Wails (Go + web frontend) that provides a split-pane editor with live preview, tag-based organization via frontmatter, and search capabilities. The goal is a lightweight, fast alternative to Electron-based markdown apps like Obsidian, leveraging Go for the backend while using standard web technologies for the UI.

## What Changes

- Create new Wails-based macOS application
- Implement split-pane layout: editor (left) + live markdown preview (right)
- Add frontmatter parsing for YAML metadata (tags, created date, etc.)
- Build tag sidebar with filtering by tags derived from frontmatter
- Implement file watching to detect external changes to notes
- Add basic search (filename + frontmatter substring) with path to upgrade
- Use CodeMirror 6 instead of Monaco for lighter weight
- Externalize sync: users point notes folder at iCloud/Dropbox/whatever they use

## Capabilities

### New Capabilities
- `note-editor`: Split-pane markdown editor with CodeMirror 6 and live markdown-it + Shiki preview
- `note-management`: CRUD operations for notes (list, read, write, delete, watch file system)
- `note-tags`: Tag indexing derived from YAML frontmatter, sidebar filtering
- `note-search`: Substring search across filenames and frontmatter content

### Modified Capabilities
- (none - this is a new application)

## Impact

- New Go backend service handling file operations, frontmatter parsing, search indexing
- New Svelte/Vite frontend with CodeMirror editor and live preview panel
- Wails runtime bridging Go and web frontend, producing native macOS .app
- Notes stored as markdown files in user-chosen directory (default: ~/Documents/MarkdownNotebook/)