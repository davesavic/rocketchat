package data

type CreateUser struct {
	baseReader
	baseValidator

	Name     string `json:"name" validate:"required,min=1,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
	Username string `json:"username" validate:"required,min=3,max=100"`

	Active                bool           `json:"active,omitzero" validate:"omitempty"`
	Nickname              string         `json:"nickname,omitzero" validate:"omitempty,min=1,max=100"`
	Bio                   string         `json:"bio,omitzero" validate:"omitempty,min=1,max=100"`
	JoinDefaultChannels   bool           `json:"joinDefaultChannels,omitzero" validate:"omitempty"`
	StatusText            string         `json:"statusText,omitzero" validate:"omitempty,min=1,max=100"`
	Roles                 []string       `json:"roles,omitzero" validate:"omitempty"`
	RequirePasswordChange bool           `json:"requirePasswordChange,omitzero" validate:"omitempty"`
	SetRandomPassword     bool           `json:"setRandomPassword,omitzero" validate:"omitempty"`
	SendWelcomeEmail      bool           `json:"sendWelcomeEmail,omitzero" validate:"omitempty"`
	Verified              bool           `json:"verified,omitzero" validate:"omitempty"`
	CustomFields          map[string]any `json:"customFields,omitzero" validate:"omitempty"`
}

type LoginUser struct {
	baseReader
	baseValidator

	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}
