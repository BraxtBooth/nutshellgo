package nutshellgo

import "time"

// Task represents a task in Nutshell.
type Task struct {
	ID            int          `json:"id,omitempty"`
	Title         string       `json:"title,omitempty"`
	EntityType    string       `json:"entityType,omitempty"`
	Rev           string       `json:"rev,omitempty"`
	Description   string       `json:"description,omitempty"`
	CreatedTime   *time.Time   `json:"createdTime,omitempty"`
	Creator       User         `json:"creator,omitempty"`
	DueTime       *time.Time   `json:"dueTime,omitempty"`
	CompletedTime *time.Time   `json:"completedTime,omitempty"`
	CompletedUser User         `json:"completedUser,omitempty"`
	Assignee      []Assignee   `json:"assignee,omitempty"`
	Entity        []StubEntity `json:"entity,omitempty"`
}
