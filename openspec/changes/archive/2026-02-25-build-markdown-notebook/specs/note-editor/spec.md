## ADDED Requirements

### Requirement: Split-pane editor layout
The system SHALL display a horizontal split-pane layout with the editor on the left and live preview on the right. The split position SHALL be adjustable by dragging the divider.

#### Scenario: Editor displays markdown source
- **WHEN** user opens a note
- **THEN** the left pane SHALL display the raw markdown content with syntax highlighting

#### Scenario: Live preview updates on edit
- **WHEN** user modifies content in the editor
- **THEN** the preview pane SHALL update within 100ms to show rendered HTML

#### Scenario: Preview renders markdown correctly
- **WHEN** user types valid markdown syntax
- **THEN** the preview SHALL render the corresponding HTML (headings, bold, italic, lists, code blocks, links, images)

#### Scenario: Code blocks use syntax highlighting
- **WHEN** user types a fenced code block with a language identifier
- **THEN** the preview SHALL use Shiki to highlight the code with appropriate syntax colors

#### Scenario: Scroll positions sync
- **WHEN** user scrolls the editor pane
- **THEN** the preview pane SHALL scroll to maintain approximate vertical alignment

### Requirement: Editor interactions
The editor SHALL provide standard markdown editing capabilities.

#### Scenario: Cursor and selection
- **WHEN** user places cursor or selects text in the editor
- **THEN** the system SHALL display the cursor/selection with standard visual styling

#### Scenario: Undo and redo
- **WHEN** user presses Cmd+Z or Cmd+Shift+Z
- **THEN** the system SHALL undo or redo the last edit operation

#### Scenario: Find in note
- **WHEN** user presses Cmd+F
- **THEN** the system SHALL display a find dialog to search within the current note

### Requirement: File change detection
The system SHALL detect when notes are modified externally.

#### Scenario: External file modification
- **WHEN** a note file is modified outside the application
- **THEN** the system SHALL prompt the user to either reload the file or keep their current version

#### Scenario: File deletion
- **WHEN** a note file is deleted externally
- **THEN** the system SHALL remove the note from the list and notify the user