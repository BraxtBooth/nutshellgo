package nutshellgo

// Address represents an address in the system.
type Address struct {
	Addresses map[string]AddressDetails `json:"address"`
}

// AddressDetails represents the details of an address.
type AddressDetails struct {
	Name             interface{} `json:"name"`
	Location         Location    `json:"location"`
	LocationAccuracy string      `json:"locationAccuracy"`
	Address1         string      `json:"address_1"`
	Address2         string      `json:"address_2"`
	Address3         string      `json:"address_3"`
	City             string      `json:"city"`
	State            string      `json:"state"`
	PostalCode       string      `json:"postalCode"`
	Country          string      `json:"country"`
	Timezone         interface{} `json:"timezone"`
}

// Assignee represents the user or team assigned to a task.
type Assignee struct {
	ID         int    `json:"id,omitempty"`
	EntityType string `json:"entityType,omitempty"`
	Rev        string `json:"rev,omitempty"`
	Name       string `json:"name,omitempty"`
}

// Avatar represents an avatar url.
type Avatar struct {
	URL string `json:"url"`
}

// Delay represents a delay in the system.
type Delay struct {
	ID         string `json:"id"`
	EntityType string `json:"entityType"`
	Rev        string `json:"rev"`
	Name       string `json:"name"`
}

// DetailedName can be used as a more detailed name model.
type DetailedName struct {
	GivenName   string `json:"givenName"`
	FamilyName  string `json:"familyName"`
	Salutation  string `json:"salutation"`
	DisplayName string `json:"displayName"`
}

// Industry represents an industry in the system.
type Industry struct {
	ID         string `json:"id"`
	EntityType string `json:"entityType"`
	Rev        string `json:"rev"`
	Name       string `json:"name"`
}

// Location represents the geographic location of an address.
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// Market represents a market in the system.
type Market struct {
	ID                string `json:"id"`
	EntityType        string `json:"entityType"`
	Rev               string `json:"rev"`
	Name              string `json:"name"`
	CurrencyShortname string `json:"currencyShortname"`
}

// Milestone represents a milestone in the system.
type Milestone struct {
	ID         int    `json:"id"`
	EntityType string `json:"entityType"`
	Rev        string `json:"rev"`
	Name       string `json:"name"`
	Position   int    `json:"position"`
	StagesetID int    `json:"stagesetId"`
}

// Origin represents an origin in the system.
type Origin struct {
	ID         string `json:"id"`
	EntityType string `json:"entityType"`
	Name       string `json:"name"`
}

// Phone represents multiple phone numbers.
type Phone struct {
	Numbers map[string]PhoneDetails `json:"phone"`
}

// PhoneDetails represents the details of a phone number.
type PhoneDetails struct {
	CountryCode string `json:"countryCode"`
	Number      string `json:"number"`
	Extension   string `json:"extension"`
}

// Price represents the currency and amount of a price.
type Price struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

// Source represents a source in the system.
type Source struct {
	ID         string `json:"id"`
	EntityType string `json:"entityType"`
	Rev        string `json:"rev"`
	Name       string `json:"name"`
}

// StubEntity represents a stub for an entity in the Nutshell API.
type StubEntity struct {
	ID           string `json:"id,omitempty"`
	EntityType   string `json:"entityType,omitempty"`
	Rev          string `json:"rev,omitempty"`
	Name         string `json:"name,omitempty"`         // Lead_ProductMap
	Relationship string `json:"relationship,omitempty"` // Lead_CompetitorMap
	Status       string `json:"status,omitempty"`       // Lead_CompetitorMap
	Quanitity    string `json:"quantity,omitempty"`     // Lead_ProductMap
	Price        Price  `json:"price,omitempty"`        // Lead_ProductMap
	Description  string `json:"description,omitempty"`  // Lead_ContactMap, Lead_AccountMap
}
