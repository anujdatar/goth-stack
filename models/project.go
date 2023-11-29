package models

import "time"

type Project struct {
	ProjectID   int       `json:"project_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerID     int       `json:"owner_id"`
	TeamID      int       `json:"team_id"`
	Deleted     int       `json:"deleted"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProjectUser struct {
	ProjectID int `json:"project_id"`
	UserID    int `json:"user_id"`
}

type ProjectAdmin struct {
	ProjectID int `json:"project_id"`
	UserID    int `json:"user_id"`
}
