## Why

Add support for Mermaid diagrams in markdown preview. Currently, the app only renders plain markdown with syntax-highlighted code blocks. Users want to include diagrams (flowcharts, sequence diagrams, etc.) directly in their notes without needing external tools or screenshots.

Mermaid is a pure JavaScript solution that renders diagrams in the browser - no backend or Java installation required.

## What Changes

- Add Mermaid npm package to frontend dependencies
- Configure markdown-it to use Mermaid plugin for rendering ```mermaid code blocks
- Update preview to display rendered SVG diagrams
- Support common diagram types: flowchart, sequenceDiagram, classDiagram, stateDiagram, gantt

## Capabilities

### New Capabilities
- `mermaid-diagrams`: Render Mermaid diagram syntax in markdown preview as SVG images

### Modified Capabilities
- `note-editor`: Update live preview to handle mermaid code blocks (delta spec)

## Impact

- Frontend: Add @vscode/markdown-it-mermaid package, update App.svelte to configure
- Dependencies: One new npm package
- No backend changes required