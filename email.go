package nutshellgo

import "time"

// Email represents an email or email stub in Nutshell.
type Email struct {
	Stub        bool             `json:"stub,omitempty"`
	ID          int              `json:"id,omitempty"`
	EntityType  string           `json:"entityType,omitempty"`
	Zendesk     Zendesk          `json:"zendesk,omitempty"`
	SentTime    *time.Time       `json:"sentTime,omitempty"`
	CreatedTime *time.Time       `json:"createdTime,omitempty"`
	To          []EmailRecipient `json:"to,omitempty"`
	From        []EmailRecipient `json:"from,omitempty"`
	CC          []EmailRecipient `json:"cc,omitempty"`
	Subject     string           `json:"subject,omitempty"`
	Body        string           `json:"body,omitempty"`
	BodyExcerpt string           `json:"bodyExcerpt,omitempty"` // Stub field
	Accounts    []Account        `json:"accounts,omitempty"`
	Contacts    []Contact        `json:"contacts,omitempty"`
	Leads       []Lead           `json:"leads,omitempty"`
}

// Zendesk represents the Zendesk details in an email.
type Zendesk struct {
	ID   int    `json:"id"`
	Link string `json:"link"`
}

// EmailRecipient represents a recipient in an email.
type EmailRecipient struct {
	Address string      `json:"address"`
	Display string      `json:"display"`
	Entity  EmailEntity `json:"entity"`
}

// EmailEntity represents the entity details in an email.
type EmailEntity struct {
	EntityName string `json:"entityName"`
	ID         string `json:"id"`
}
