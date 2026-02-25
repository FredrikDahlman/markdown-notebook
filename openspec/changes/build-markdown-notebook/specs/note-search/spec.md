## ADDED Requirements

### Requirement: Search interface
The system SHALL provide a search input to find notes.

#### Scenario: Search input displayed
- **WHEN** user focuses search input
- **THEN** the system SHALL display a text input for entering search queries

#### Scenario: Search executes on input
- **WHEN** user types in search input
- **THEN** the system SHALL search as the user types (debounced 200ms)

### Requirement: Search scope
The search SHALL find notes matching the query in filenames and frontmatter.

#### Scenario: Search filename
- **WHEN** user searches for "meeting"
- **THEN** the system SHALL return notes with "meeting" in the filename

#### Scenario: Search frontmatter title
- **WHEN** user searches for text
- **THEN** the system SHALL search in the frontmatter title field

#### Scenario: Search frontmatter tags
- **WHEN** user searches for a tag name
- **THEN** the system SHALL return notes that have that tag

### Requirement: Search results display
The system SHALL display search results with context.

#### Scenario: Results show matching notes
- **WHEN** search returns results
- **THEN** the system SHALL display matching notes in the sidebar, sorted by relevance

#### Scenario: No results
- **WHEN** search has no matches
- **THEN** the system SHALL display "No notes found" message

#### Scenario: Clear search
- **WHEN** user clears the search input
- **THEN** the system SHALL return to showing all notes