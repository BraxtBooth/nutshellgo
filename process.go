package nutshellgo

import "time"

// Process represents a process in Nutshell.
type Process struct {
	ID                 int        `json:"id,omitempty"`
	EntityType         string     `json:"entityType,omitempty"`
	Rev                string     `json:"rev,omitempty"`
	Name               string     `json:"name,omitempty"`
	Type               string     `json:"type,omitempty"`
	StartedTime        *time.Time `json:"startedTime,omitempty"`
	ClosedTime         *time.Time `json:"closedTime,omitempty"`
	CurrentMilestoneID int        `json:"currentMilestoneId,omitempty"`
	ProcessTemplateID  int        `json:"processTemplateId,omitempty"`
	Steps              []Step     `json:"steps,omitempty"`
}

// Step represents a step in a process.
type Step struct {
	Name               string           `json:"name,omitempty"`
	ID                 int              `json:"id,omitempty"`
	EntityType         string           `json:"entityType,omitempty"`
	Rev                string           `json:"rev,omitempty"`
	Description        string           `json:"description,omitempty"`
	MilestoneID        int              `json:"milestoneId,omitempty"`
	Sequence           int              `json:"sequence,omitempty"`
	Status             string           `json:"status,omitempty"`
	WorkableTime       *time.Time       `json:"workableTime,omitempty"`
	DueTime            *time.Time       `json:"dueTime,omitempty"`
	CompletedTime      *time.Time       `json:"completedTime,omitempty"`
	CompletedUser      User             `json:"completedUser,omitempty"`
	LogActivityTypeID  interface{}      `json:"logActivityTypeId,omitempty"`  // TODO add type
	Assignee           interface{}      `json:"assignee,omitempty"`           // TODO add type
	PrerequisiteStepID interface{}      `json:"prerequisiteStepId,omitempty"` // TODO add type
	AvailableDelayIDs  []string         `json:"availableDelayIds,omitempty"`
	AppliedDelayIDs    []string         `json:"appliedDelayIds,omitempty"`
	Requirement        StepRequirement  `json:"requirement,omitempty"`
	Choice             StepChosenEntity `json:"choice,omitempty"`
}

// StepRequirement represents the requirement for a step.
type StepRequirement struct {
	EntityType        string `json:"entityType,omitempty"`
	FieldName         string `json:"fieldName,omitempty"`
	RequireForLeadWin bool   `json:"requireForLeadWin,omitempty"`
	MustChoose        bool   `json:"mustChoose,omitempty"`
}

// StepChosenEntity represents an entity chosen for a step.
type StepChosenEntity struct {
	EntityType string         `json:"entityType,omitempty"`
	ID         int            `json:"id,omitempty"`
	Rev        int            `json:"rev,omitempty"`
	FieldName  string         `json:"fieldName,omitempty"`
	FieldValue StepFieldValue `json:"fieldValue,omitempty"`
}

// StepFieldValue represents the field value of a step chosen entity.
type StepFieldValue struct {
	Name             string `json:"name,omitempty"`
	Location         string `json:"location,omitempty"`
	LocationAccuracy string `json:"locationAccuracy,omitempty"`
	Address1         string `json:"address_1,omitempty"`
	Address2         string `json:"address_2,omitempty"`
	Address3         string `json:"address_3,omitempty"`
	City             string `json:"city,omitempty"`
	State            string `json:"state,omitempty"`
	PostalCode       string `json:"postalCode,omitempty"`
	Country          string `json:"country,omitempty"`
}
