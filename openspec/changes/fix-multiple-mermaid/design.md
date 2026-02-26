## Context

The current markdown rendering process extracts ````mermaid``` code blocks and replaces them with `<div id="mermaid-0"></div>` placeholders. It then loops over the extracted blocks and calls `mermaid.render()` for each. However, there's a bug in how multiple blocks are parsed or rendered, causing only the first one to appear.

## Goals / Non-Goals

**Goals:**
- Fix the bug preventing multiple mermaid diagrams from rendering
- Ensure independent error handling so one failed diagram doesn't stop others from rendering
- Make sure IDs generated are strictly unique across renders

**Non-Goals:**
- Upgrading or changing the mermaid version (sticking to the current working v9 setup)
- Supporting other diagram types

## Decisions

### 1. Fix regex replacement loop
**Decision:** Review and fix the `html.replace` logic or the subsequent rendering loop. 
**Rationale:** The issue likely stems from how `mermaidRegex.exec()` or `String.replace()` is matching multiple instances. If `replace` with a global regex string is used, the function should correctly iterate, but we need to ensure the ID incrementing works properly.

### 2. Independent try/catch blocks
**Decision:** Ensure `mermaid.render()` is wrapped in individual try/catch blocks for each diagram.
**Rationale:** If diagram 2 fails, it shouldn't stop diagram 3 from rendering.

## Risks / Trade-offs

| Risk | Impact | Mitigation |
|------|--------|------------|
| ID collision across re-renders | Render fails or wrong element | Reset counter on each `updatePreview` call |