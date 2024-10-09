package v1

import (
	"net/url"
)

var (
	EntityTypes     = []string{"individual", "group", "organisation", "other"}
	EntityRoles     = []string{"owner", "steward", "maintainer", "contributor", "other"}
	ChannelTypes    = []string{"bank", "payment-provider", "cheque", "cash", "other"}
	PlanFrequencies = []string{"one-time", "weekly", "fortnightly", "monthly", "yearly", "other"}
	PlanStatuses    = []string{"active", "inactive"}
)

//easyjson:json
type URL struct {
	URL       string `json:"url"`
	WellKnown string `json:"wellKnown,omitempty"`

	// Parsed URLs.
	URLobj       *url.URL `json:"-" db:"-"`
	WellKnownObj *url.URL `json:"-" db:"-"`
}

// Entity represents an entity in charge of a project: individual, organisation etc.
//
//easyjson:json
type Entity struct {
	Type        string `json:"type"`
	Role        string `json:"role"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	WebpageURL  URL    `json:"webpageUrl"`
}

// Project represents a FOSS project.
//
//easyjson:json
type Project struct {
	GUID          string   `json:"guid"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	WebpageURL    URL      `json:"webpageUrl"`
	RepositoryURL URL      `json:"repositoryUrl"`
	Licenses      []string `json:"licenses"`
	Tags          []string `json:"tags"`
}

//easyjson:json
type Projects []Project

// Channel is a loose representation of a payment channel. eg: bank, cash, or a processor like PayPal.
//
//easyjson:json
type Channel struct {
	GUID        string `json:"guid"`
	Type        string `json:"type"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

// easyjson:json
type Channels []Channel

// Plan represents a payment plan / ask for the project.
//
//easyjson:json
type Plan struct {
	GUID        string   `json:"guid"`
	Status      string   `json:"status"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Amount      float64  `json:"amount"`
	Currency    string   `json:"currency"`
	Frequency   string   `json:"frequency"`
	Channels    []string `json:"channels"`
}

// easyjson:json
type Plans []Plan

// History represents a very course, high level income/expense statement.
//
//easyjson:json
type HistoryItem struct {
	Year        int     `json:"year"`
	Income      float64 `json:"income"`
	Expenses    float64 `json:"expenses"`
	Description string  `json:"description"`
}

// easyjson:json
type History []HistoryItem

// easyjson:json
type Funding struct {
	Channels Channels `json:"channels"`
	Plans    Plans    `json:"plans"`
	History  History  `json:"history"`
}

//easyjson:json
type Manifest struct {
	// This is added internally and is not expected in the manifest itself.
	URL URL `json:"-" db:"-"`

	Version  string   `json:"version"`
	Entity   Entity   `json:"entity"`
	Projects Projects `json:"projects"`
	Funding  Funding  `json:"funding"`
}
