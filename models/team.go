package models

import "time"

type Team struct {
	TeamID         int       `json:"team_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	OwnerID        int       `json:"owner_id"`
	SubscriptionID int       `json:"subscription_id"`
	Deleted        int       `json:"deleted"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type TeamUser struct {
	TeamID int `json:"team_id"`
	UserID int `json:"user_id"`
}

type TeamAdmin struct {
	TeamID int `json:"team_id"`
	UserID int `json:"user_id"`
}

type TeamProject struct {
	TeamID    int `json:"team_id"`
	ProjectID int `json:"project_id"`
}

// type TeamSubscription struct {
// 	TeamID         int `json:"team_id"`
// 	SubscriptionID int `json:"subscription_id"`
// }
