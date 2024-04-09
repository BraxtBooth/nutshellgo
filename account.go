package nutshellgo

import "time"

// Account models a Nutshell account. Accounts have been renamed to 'Companies' on Nutshell frontend.
type Account struct {
	ID                int               `json:"id,omitempty"`
	EntityType        string            `json:"entityType,omitempty"`
	Rev               string            `json:"rev,omitempty"`
	Name              string            `json:"name,omitempty"`
	Description       string            `json:"description,omitempty"`
	LegacyID          string            `json:"legacyId,omitempty"`
	HTMLURL           string            `json:"htmlUrl,omitempty"`
	Avatar            Avatar            `json:"avatar,omitempty"`
	AvatarURL         string            `json:"avatarUrl,omitempty"`
	Tags              []string          `json:"tags,omitempty"`
	Industry          Industry          `json:"industry,omitempty"`
	AccountType       AccountType       `json:"accountType,omitempty"`
	Creator           User              `json:"creator,omitempty"`
	Leads             []Lead            `json:"leads,omitempty"`
	Contacts          []Contact         `json:"contacts,omitempty"`
	Phone             Phone             `json:"phone,omitempty"`
	URL               map[string]string `json:"url,omitempty"`
	Email             map[string]string `json:"email,omitempty"`
	Territory         StubEntity        `json:"territory,omitempty"`
	MergedInto        []Account         `json:"mergedInto,omitempty"`
	LastContactedDate *time.Time        `json:"lastContactedDate,omitempty"`
	DeletedTime       *time.Time        `json:"deletedTime,omitempty"`
	ModifiedTime      *time.Time        `json:"modifiedTime,omitempty"`
	CreatedTime       *time.Time        `json:"createdTime,omitempty"`
}

type AccountType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
