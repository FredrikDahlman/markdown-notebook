## ADDED Requirements

### Requirement: Note list display
The system SHALL display a list of all notes in the notes directory.

#### Scenario: List notes on startup
- **WHEN** application starts
- **THEN** the system SHALL scan the notes directory and display all .md files in a sidebar

#### Scenario: Note shows title
- **WHEN** note is displayed in list
- **THEN** the system SHALL use the first H1 heading as the display title, or the filename if no H1 exists

#### Scenario: Sort notes
- **WHEN** notes are displayed
- **THEN** the system SHALL sort by last modified date (newest first) by default

### Requirement: Create new note
The system SHALL allow users to create new markdown notes.

#### Scenario: Create note
- **WHEN** user clicks "New Note" button
- **THEN** the system SHALL create a new .md file with a timestamp-based filename and frontmatter template

#### Scenario: Note has frontmatter template
- **WHEN** new note is created
- **THEN** the note SHALL include YAML frontmatter with created date and empty tags array

### Requirement: Read note
The system SHALL load and display note content.

#### Scenario: Open note
- **WHEN** user clicks a note in the sidebar
- **THEN** the system SHALL load the file content and display it in the editor

#### Scenario: Read frontmatter
- **WHEN** note with YAML frontmatter is loaded
- **THEN** the system SHALL parse the frontmatter and expose tags and metadata for display

### Requirement: Write note
The system SHALL save note changes to disk.

#### Scenario: Auto-save on edit
- **WHEN** user modifies note content
- **THEN** the system SHALL auto-save the file after 500ms of inactivity

#### Scenario: Manual save
- **WHEN** user presses Cmd+S
- **THEN** the system SHALL immediately save the current note to disk

### Requirement: Delete note
The system SHALL allow users to delete notes.

#### Scenario: Delete note
- **WHEN** user right-clicks a note and selects "Delete"
- **THEN** the system SHALL move the file to trash (not permanent delete)

### Requirement: Notes directory configuration
The system SHALL allow users to configure the notes directory.

#### Scenario: Default notes directory
- **WHEN** application runs for the first time
- **THEN** the system SHALL create ~/Documents/MarkdownNotebook/ as the default notes directory

#### Scenario: Custom notes directory
- **WHEN** user changes notes directory in settings
- **THEN** the system SHALL use the new directory for all note operations