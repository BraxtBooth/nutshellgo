package nutshellgo

import "time"

// User represents a user or user stub in Nutshell.
type User struct {
	ID              int        `json:"id,omitempty"`
	Stub            bool       `json:"stub,omitempty"`
	EntityType      string     `json:"entityType,omitempty"`
	Rev             string     `json:"rev,omitempty"`
	Name            string     `json:"name,omitempty"`
	FirstName       string     `json:"firstName,omitempty"`
	LastName        string     `json:"lastName,omitempty"`
	IsEnabled       bool       `json:"isEnabled,omitempty"`
	IsAdministrator bool       `json:"isAdministrator,omitempty"`
	Emails          []string   `json:"emails,omitempty"`
	Teams           Teams      `json:"teams,omitempty"`
	Email           string     `json:"email,omitempty"`    // DEPRECATED (according to Nutshell Class Reference)
	Username        string     `json:"username,omitempty"` // DEPRECATED (according to Nutshell Class Reference)
	ModifiedTime    *time.Time `json:"modifiedTime,omitempty"`
	CreatedTime     *time.Time `json:"createdTime,omitempty"`
}

// Team represents a team in Nutshell.
type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Teams represents the user's teams in different access levels.
type Teams struct {
	ViewOwn  []Team `json:"viewOwn,omitempty"`
	ViewTeam []Team `json:"viewTeam,omitempty"`
	ViewAll  []Team `json:"viewAll,omitempty"`
}
