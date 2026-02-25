## ADDED Requirements

### Requirement: Tag indexing
The system SHALL index tags from note frontmatter.

#### Scenario: Parse tags from frontmatter
- **WHEN** note is loaded
- **THEN** the system SHALL extract the tags array from YAML frontmatter

#### Scenario: Tags are case-insensitive
- **WHEN** tags are processed
- **THEN** the system SHALL treat "Work" and "work" as the same tag

#### Scenario: Rebuild tag index
- **WHEN** notes directory changes or notes are modified
- **THEN** the system SHALL rebuild the tag index

### Requirement: Tag sidebar display
The system SHALL display available tags in a sidebar.

#### Scenario: Show all tags
- **WHEN** sidebar displays tags
- **THEN** the system SHALL show all unique tags found across all notes with a count of notes per tag

#### Scenario: Tags sorted alphabetically
- **WHEN** tags are displayed
- **THEN** the system SHALL sort tags alphabetically

### Requirement: Filter by tag
The system SHALL allow users to filter notes by tag.

#### Scenario: Click tag to filter
- **WHEN** user clicks a tag in the sidebar
- **THEN** the system SHALL filter the note list to show only notes with that tag

#### Scenario: Clear tag filter
- **WHEN** user clicks "Clear filter" or selects all notes
- **THEN** the system SHALL show all notes again

#### Scenario: Multiple tag selection
- **WHEN** user holds Cmd and clicks additional tags
- **THEN** the system SHALL show notes that have ANY of the selected tags (OR logic)