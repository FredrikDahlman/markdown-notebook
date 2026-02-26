## ADDED Requirements

### Requirement: Mermaid diagram rendering
The system SHALL render Mermaid diagram syntax in markdown preview as SVG images.

#### Scenario: Render flowchart
- **WHEN** user has a ```mermaid code block with flowchart syntax in their note
- **THEN** the preview SHALL display the rendered flowchart as an SVG image

#### Scenario: Render sequence diagram
- **WHEN** user has a ```mermaid code block with sequenceDiagram syntax in their note
- **THEN** the preview SHALL display the rendered sequence diagram as an SVG image

#### Scenario: Render class diagram
- **WHEN** user has a ```mermaid code block with classDiagram syntax in their note
- **THEN** the preview SHALL display the rendered class diagram as an SVG image

#### Scenario: Render state diagram
- **WHEN** user has a ```mermaid code block with stateDiagram syntax in their note
- **THEN** the preview SHALL display the rendered state diagram as an SVG image

#### Scenario: Render Gantt chart
- **WHEN** user has a ```mermaid code block with gantt syntax in their note
- **THEN** the preview SHALL display the rendered Gantt chart as an SVG image

#### Scenario: Invalid Mermaid syntax
- **WHEN** user has a ```mermaid code block with invalid syntax
- **THEN** the preview SHALL display an error message instead of crashing

#### Scenario: Multiple diagrams in one note
- **WHEN** user has multiple ```mermaid code blocks in their note
- **THEN** the preview SHALL render all diagrams in order