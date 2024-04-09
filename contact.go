package nutshellgo

import "time"

// Contact represents a contact in Nutshell.
type Contact struct {
	ID                int               `json:"id,omitempty"`
	EntityType        string            `json:"entityType,omitempty"`
	Rev               string            `json:"rev,omitempty"`
	Name              interface{}       `json:"name,omitempty"` // Can be a string or DetailedName struct
	HTMLURL           string            `json:"htmlUrl,omitempty"`
	Avatar            Avatar            `json:"avatar,omitempty"`
	AvatarURL         string            `json:"avatarUrl,omitempty"`
	Description       string            `json:"description,omitempty"`
	LegacyID          string            `json:"legacyId,omitempty"`
	Tags              []string          `json:"tags,omitempty"`
	Creator           User              `json:"creator,omitempty"`
	Leads             []Lead            `json:"leads,omitempty"`
	Accounts          []Account         `json:"accounts,omitempty"`
	ContactedCount    int               `json:"contactedCount,omitempty"`
	Address           Address           `json:"address,omitempty"`
	Phone             Phone             `json:"phone,omitempty"`
	URL               map[string]string `json:"url,omitempty"`
	Email             map[string]string `json:"email,omitempty"`
	Notes             []Note            `json:"notes,omitempty"`
	Territory         StubEntity        `json:"territory,omitempty"`
	MergedInto        []Contact         `json:"mergedInto,omitempty"`
	LastContactedDate *time.Time        `json:"lastContactedDate,omitempty"`
	DeletedTime       *time.Time        `json:"deletedTime,omitempty"`
	ModifiedTime      *time.Time        `json:"modifiedTime,omitempty"`
	CreatedTime       *time.Time        `json:"createdTime,omitempty"`
}
