package model

import "time"

type User struct {
	ID                    string         `json:"_id"`
	Username              string         `json:"username"`
	Emails                []Email        `json:"emails"`
	Type                  string         `json:"type"`
	Roles                 []string       `json:"roles"`
	Status                string         `json:"status"`
	Active                bool           `json:"active"`
	CreatedAt             time.Time      `json:"createdAt"`
	UpdatedAt             time.Time      `json:"_updatedAt"`
	Bio                   string         `json:"bio"`
	Name                  string         `json:"name"`
	Nickname              string         `json:"nickname"`
	RequirePasswordChange bool           `json:"requirePasswordChange"`
	Settings              map[string]any `json:"settings"`
}

type Email struct {
	Address  string `json:"address"`
	Verified bool   `json:"verified"`
}

type AuthenticatedUser struct {
	AuthToken string `json:"authToken"`
	UserID    string `json:"userId"`
	User      User   `json:"me"`
}
