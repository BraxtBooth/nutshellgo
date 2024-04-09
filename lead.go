package nutshellgo

import "time"

// Lead represents a lead or lead stub in Nutshell.
type Lead struct {
	ID                 int            `json:"id,omitempty"`
	Stub               bool           `json:"stub,omitempty"`
	EntityType         string         `json:"entityType,omitempty"`
	Rev                string         `json:"rev,omitempty"`
	Name               string         `json:"name,omitempty"`
	HtmlURL            string         `json:"htmlUrl,omitempty"`
	AvatarURL          string         `json:"avatarUrl,omitempty"`
	Tags               []string       `json:"tags,omitempty"`
	Description        string         `json:"description,omitempty"`
	CreatedTime        *time.Time     `json:"createdTime,omitempty"`
	Creator            User           `json:"creator,omitempty"`
	PrimaryAccount     Account        `json:"primaryAccount,omitempty"`
	Milestone          Milestone      `json:"milestone,omitempty"`
	Stageset           Stageset       `json:"stageset,omitempty"`
	ActivitiesCount    map[string]int `json:"activitiesCount,omitempty"`
	Status             int            `json:"status,omitempty"`
	Confidence         int            `json:"confidence,omitempty"`
	Completion         int            `json:"completion,omitempty"`
	Urgency            string         `json:"urgency,omitempty"`
	IsOverdue          bool           `json:"isOverdue,omitempty"`
	Market             Market         `json:"market,omitempty"`
	Assignee           User           `json:"assignee,omitempty"`
	Processes          []Process      `json:"processes,omitempty"`
	Sources            []StubEntity   `json:"sources,omitempty"`
	Channels           []string       `json:"channels,omitempty"`
	Competitors        []StubEntity   `json:"competitors,omitempty"`
	Products           []StubEntity   `json:"products,omitempty"`
	Contacts           []StubEntity   `json:"contacts,omitempty"`
	Accounts           []StubEntity   `json:"accounts,omitempty"`
	DueTime            *time.Time     `json:"dueTime,omitempty"`
	NextStepDueTime    *time.Time     `json:"nextStepDueTime,omitempty"`
	Value              Price          `json:"value,omitempty"`
	NormalizedValue    Price          `json:"normalizedValue,omitempty"`
	Notes              []Note         `json:"notes,omitempty"`
	LastContactedDate  *time.Time     `json:"lastContactedDate,omitempty"`
	DeletedTime        *time.Time     `json:"deletedTime,omitempty"`
	Priority           int            `json:"priority,omitempty"`
	PrimaryAccountName string         `json:"primaryAccountName,omitempty"` // Stub field
	PrimaryContactName string         `json:"primaryContactName,omitempty"` // Stub field
	ClosedTime         *time.Time     `json:"closedTime,omitempty"`         // Stub field
}

// Stageset represents a stageset in Nutshell.
type Stageset struct {
	ID         int    `json:"id,omitempty"`
	EntityType string `json:"entityType,omitempty"`
	Name       string `json:"name,omitempty"`
	Default    int    `json:"default,omitempty"`
	Position   int    `json:"position,omitempty"`
}
