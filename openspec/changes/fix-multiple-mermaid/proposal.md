## Why

Currently, when a markdown note contains multiple Mermaid diagrams, only the first one renders correctly as an SVG. The subsequent diagrams fail to render and just show their source code or an error message. This limits users from documenting complex systems or processes that require multiple diagrams in a single note.

## What Changes

- Fix the Mermaid diagram rendering logic in the frontend to correctly process and render multiple `mermaid` code blocks in a single document
- Ensure each diagram gets a unique ID and is rendered independently

## Capabilities

### New Capabilities
None

### Modified Capabilities
- `mermaid-diagrams`: Update requirement to ensure multiple diagrams render correctly in the same note

## Impact

- Frontend: `App.svelte` updatePreview function will be modified to handle multiple diagram instances correctly.
- No backend changes required.