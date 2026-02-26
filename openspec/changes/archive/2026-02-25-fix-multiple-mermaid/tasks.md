## 1. Implementation

- [x] 1.1 Fix regex parsing to properly extract all mermaid code blocks in a note
- [x] 1.2 Ensure unique ID generation works correctly for multiple diagrams
- [x] 1.3 Add independent error handling so one failed diagram doesn't block others

## 2. Testing

- [x] 2.1 Verify a single markdown note with three different mermaid diagrams renders all of them
- [x] 2.2 Verify that one invalid diagram doesn't break the rendering of the valid ones