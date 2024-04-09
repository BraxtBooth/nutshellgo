package nutshellgo

import "time"

// Note represents a note or note stub in Nutshell.
type Note struct {
	Stub               bool       `json:"stub,omitempty"`
	EntityType         string     `json:"entityType,omitempty"`
	ID                 int        `json:"id,omitempty"`
	Rev                string     `json:"rev,omitempty"`
	User               User       `json:"user,omitempty"`
	OriginID           int        `json:"originId,omitempty"`
	Note               string     `json:"note,omitempty"`
	NoteMarkup         string     `json:"noteMarkup,omitempty"`
	NoteHTML           string     `json:"noteHtml,omitempty"`
	Date               *time.Time `json:"date,omitempty"`
	Files              []File     `json:"files,omitempty"`
	Mentions           []Mention  `json:"mentions,omitempty"`
	Lead               Lead       `json:"lead,omitempty"`
	Contact            Contact    `json:"contact,omitempty"`
	Account            Account    `json:"account,omitempty"`
	PrimaryAccountName string     `json:"primaryAccountName,omitempty"`
	DueTime            *time.Time `json:"dueTime,omitempty"`
	NextStepDueTime    *time.Time `json:"nextStepDueTime,omitempty"`
	CreatedTime        *time.Time `json:"createdTime,omitempty"`
}

// File represents a file attached to a note.
type File struct {
	EntityType string `json:"entityType,omitempty"`
	ID         int    `json:"id,omitempty"`
	URI        string `json:"uri,omitempty"`
	Name       string `json:"name,omitempty"`
	MIME       string `json:"mime,omitempty"`
	Rev        string `json:"rev,omitempty"`
	Size       string `json:"size,omitempty"`
	ClientType string `json:"clientType,omitempty"`
}

// Mention represents a mention in a note.
type Mention struct {
	EntityType string `json:"entityType,omitempty"`
	ID         int    `json:"id,omitempty"`
	Rev        int    `json:"rev,omitempty"`
	JobTitle   string `json:"jobTitle,omitempty"`
	Name       string `json:"name,omitempty"`
	Stub       bool   `json:"stub,omitempty"`
}
