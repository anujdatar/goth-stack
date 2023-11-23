package models

type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TeamID      int    `json:"team_id"`
	Team        Team   `json:"team"`
	Users       []User `json:"users"`
	Admins      []User `json:"admins"`
	Deleted     bool   `json:"deleted"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
