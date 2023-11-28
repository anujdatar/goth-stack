package models

import "time"

type Subscription struct {
	SubscriptionID int       `json:"subscription_id"`
	Name           string    `json:"name"`
	Price          int       `json:"price"`
	Deleted        int       `json:"deleted"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
