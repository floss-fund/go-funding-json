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
	URL       string `json:"url" db:"url"`
	WellKnown string `json:"wellKnown,omitempty" db:"wellKnown"`

	// Parsed URLs.
	URLobj       *url.URL `json:"-" db:"-"`
	WellKnownObj *url.URL `json:"-" db:"-"`
}

// Entity represents an entity in charge of a project: individual, organisation etc.
//
//easyjson:json
type Entity struct {
	Type        string `json:"type" db:"type"`
	Role        string `json:"role" db:"role"`
	Name        string `json:"name" db:"name"`
	Email       string `json:"email" db:"email"`
	Phone       string `json:"phone" db:"phone"`
	Description string `json:"description" db:"description"`
	WebpageURL  URL    `json:"webpageUrl" db:"webpageUrl"`
}

// Project represents a FOSS project.
//
//easyjson:json
type Project struct {
	GUID          string   `json:"guid" db:"guid"`
	Name          string   `json:"name" db:"name"`
	Description   string   `json:"description" db:"description"`
	WebpageURL    URL      `json:"webpageUrl" db:"webpageUrl"`
	RepositoryURL URL      `json:"repositoryUrl" db:"repositoryUrl"`
	Licenses      []string `json:"licenses" db:"licenses"`
	Tags          []string `json:"tags" db:"tags"`
}

//easyjson:json
type Projects []Project

// Channel is a loose representation of a payment channel. eg: bank, cash, or a processor like PayPal.
//
//easyjson:json
type Channel struct {
	GUID        string `json:"guid" db:"guid"`
	Type        string `json:"type" db:"type"`
	Address     string `json:"address" db:"address"`
	Description string `json:"description" db:"description"`
}

// easyjson:json
type Channels []Channel

// Plan represents a payment plan / ask for the project.
//
//easyjson:json
type Plan struct {
	GUID        string   `json:"guid" db:"guid"`
	Status      string   `json:"status" db:"status"`
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	Amount      float64  `json:"amount" db:"amount"`
	Currency    string   `json:"currency" db:"currency"`
	Frequency   string   `json:"frequency" db:"frequency"`
	Channels    []string `json:"channels" db:"channels"`
}

// easyjson:json
type Plans []Plan

// History represents a very course, high level income/expense statement.
//
//easyjson:json
type HistoryItem struct {
	Year        int     `json:"year" db:"year"`
	Income      float64 `json:"income" db:"income"`
	Expenses    float64 `json:"expenses" db:"expenses"`
	Taxes       float64 `json:"taxes" db:"taxes"`
	Currency    string  `json:"currency" db:"currency"`
	Description string  `json:"description" db:"description"`
}

// easyjson:json
type History []HistoryItem

// easyjson:json
type Funding struct {
	Channels Channels `json:"channels" db:"channels"`
	Plans    Plans    `json:"plans" db:"plans"`
	History  History  `json:"history" db:"history"`
}

//easyjson:json
type Manifest struct {
	// This is added internally and is not expected in the manifest itself.
	URL URL `json:"-" db:"-"`

	Version  string   `json:"version" db:"version"`
	Entity   Entity   `json:"entity" db:"entity"`
	Projects Projects `json:"projects" db:"projects"`
	Funding  Funding  `json:"funding" db:"funding"`
}
