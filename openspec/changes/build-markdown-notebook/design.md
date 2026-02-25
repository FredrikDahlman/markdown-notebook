## Context

Building a native macOS markdown notebook app using Wails (Go backend + web frontend). The application provides a split-pane editor with live preview, tag-based organization via YAML frontmatter, and search capabilities.

## Goals / Non-Goals

**Goals:**
- Native macOS .app using Wails framework
- Split-pane layout: CodeMirror editor (left) + live markdown preview (right)
- YAML frontmatter parsing for tags/metadata
- Tag sidebar with filtering
- Search across filenames and frontmatter
- File watching to detect external changes
- Fast, lightweight alternative to Electron-based markdown apps

**Non-Goals:**
- Built-in cloud sync (externalize to user's preferred solution)
- Real-time collaboration
- Mobile support
- Plugin system
- Full-text search initially (upgrade path to bleve later)

## Decisions

### 1. Wails over pure Go UI or Electron
**Decision:** Use Wails to build the app
**Rationale:** Wails provides native macOS .app output while allowing Go backend + web frontend. Better than gioui (non-native look) and lighter than Electron (Node.js overhead).

### 2. Svelte over React/Vue
**Decision:** Use Svelte + Vite for frontend
**Rationale:** Less boilerplate than React, excellent reactivity for live preview updates, smaller bundle size. Good fit for a single-pane application.

### 3. CodeMirror 6 over Monaco
**Decision:** Use CodeMirror 6 for the editor
**Rationale:** Monaco is ~2MB, CodeMirror is ~400KB. Monaco is designed for IDEs (VS Code), CodeMirror is designed for editors. Faster cold start time critical for good UX.

**Alternative considered:** Monaco - rejected due to bundle size and cold start performance

### 4. markdown-it + Shiki for rendering
**Decision:** Use markdown-it with Shiki for syntax highlighting
**Rationale:** Standard combination, Shiki produces better highlighting than highlight.js (uses TextMate grammars like VS Code)

### 5. Single notes directory
**Decision:** Notes stored in ~/Documents/MarkdownNotebook/ (configurable)
**Rationale:** Simple mental model, lets users use iCloud/Dropbox/other sync solutions without building sync ourselves.

**Alternative considered:** SQLite database - rejected in favor of plain markdown files for portability

### 6. YAML frontmatter
**Decision:** Use YAML for frontmatter with go-yaml library
**Rationale:** Most common format (Obsidian, Jekyll, etc.), good library support

### 7. Substring search initially
**Decision:** Start with filename + frontmatter substring search
**Rationale:** Simple to implement, covers 80% of use cases. Design path to upgrade to bleve later.

## Risks / Trade-offs

| Risk | Impact | Mitigation |
|------|--------|------------|
| Wails macOS WebView limitations | Some web APIs may not work | Stick to standard APIs, test on real macOS |
| CodeMirror markdown mode gaps | No live preview hooks built-in | Wire up change events to trigger preview render |
| File watching race conditions | External edits might conflict with in-app edits | Prompt user to reload or overwrite |
| Frontmatter edge cases | Malformed YAML could crash parser | Graceful error handling, show raw text |
| Search performance at scale | Substring search O(n) on content | Upgrade path to bleve for full-text indexing |

## Open Questions

- Should notes support subdirectories? (YAML: `folder: subdir/note`)
- How to handle note renaming? (file rename vs title change in frontmatter)
- Default markdown extension: .md or .markdown?
- Keyboard shortcuts: VS Code style or macOS native?
- Dark/light mode: follow system or manual toggle?