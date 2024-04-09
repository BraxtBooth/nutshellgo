package nutshellgo

import "time"

// Activity represents an activity in Nutshell.
type Activity struct {
	ID             int           `json:"id,omitempty"`
	EntityType     string        `json:"entityType,omitempty"`
	Rev            string        `json:"rev,omitempty"`
	Name           interface{}   `json:"name,omitempty"` // Can be a string or DetailedName struct
	Description    string        `json:"description,omitempty"`
	ActivityType   ActivityType  `json:"activityType,omitempty"`
	Leads          []Lead        `json:"leads,omitempty"` // DEPRECATED (according to Nutshell Class Reference)
	StartTime      *time.Time    `json:"startTime,omitempty"`
	EndTime        *time.Time    `json:"endTime,omitempty"`
	IsAllDay       bool          `json:"isAllDay,omitempty"`
	IsFlagged      bool          `json:"isFlagged,omitempty"`
	Status         int           `json:"status,omitempty"`
	LogDescription string        `json:"logDescription,omitempty"` // DEPRECATED (according to Nutshell Class Reference)
	LogNote        Note          `json:"logNote,omitempty"`
	LoggedBy       User          `json:"loggedBy,omitempty"`
	Participants   []interface{} `json:"participants,omitempty"` // User, Contact, and Account can be participants
	Followup       *Activity     `json:"followup,omitempty"`
	FollowupTo     *Activity     `json:"followupTo,omitempty"`
	DeletedTime    *time.Time    `json:"deletedTime,omitempty"`
	ModifiedTime   *time.Time    `json:"modifiedTime,omitempty"`
	CreatedTime    *time.Time    `json:"createdTime,omitempty"`
}

// ActivityType represents the type of an activity.
type ActivityType struct {
	ID          string `json:"id"`
	EntityType  string `json:"entityType"`
	Rev         int    `json:"rev"`
	Name        string `json:"name"`
	QuotaWeight int    `json:"quotaWeight"`
}
