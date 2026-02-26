## Context

The markdown notebook app currently uses markdown-it for rendering markdown and Shiki for code block syntax highlighting. Mermaid diagrams will be rendered in the browser using JavaScript, requiring no backend changes.

## Goals / Non-Goals

**Goals:**
- Render Mermaid diagrams in the preview pane
- Support flowchart, sequence diagram, class diagram, state diagram, and Gantt chart types
- Maintain offline functionality (no external API calls)
- Minimal performance impact on preview rendering

**Non-Goals:**
- PlantUML support (future enhancement)
- Server-side rendering (not needed for this use case)
- Diagram editing (view only)

## Decisions

### 1. Use @vscode/markdown-it-mermaid plugin
**Decision:** Use the official VS Code mermaid plugin for markdown-it
**Rationale:** Well-maintained, handles the integration between markdown-it and mermaid seamlessly

### 2. Render on preview update
**Decision:** Re-render mermaid diagrams when preview updates
**Rationale:** Simple approach, diagrams are re-rendered when content changes anyway

### 3. Mermaid initialization
**Decision:** Initialize mermaid in the onMount function alongside Shiki
**Rationale:** Single initialization point, consistent with existing architecture

## Risks / Trade-offs

| Risk | Impact | Mitigation |
|------|--------|------------|
| Large diagrams slow down preview | Performance degradation | Mermaid is efficient; unlikely to be an issue |
| Mermaid syntax errors crash preview | Diagram shows error | Mermaid handles errors gracefully |
| Browser compatibility | Some older browsers may struggle | Modern browsers only (matches Electron WebView) |

## Open Questions

- Should we add a loading indicator while diagrams render?
- Should we cache rendered diagrams for performance?