package models

type Team struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	OwnerID       int            `json:"owner_id"`
	Owner         User           `json:"owner"`
	Subscriptions []Subscription `json:"subscriptions"`
	Users         []User         `json:"users"`
	Admins        []User         `json:"admins"`
	Projects      []Project      `json:"projects"`
	Deleted       bool           `json:"deleted"`
	CreatedAt     string         `json:"created_at"`
	UpdatedAt     string         `json:"updated_at"`
}
