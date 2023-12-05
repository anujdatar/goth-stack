package models

import "time"

type User struct {
	UserID            int       `json:"user_id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	Password          string    `json:"-"` // omit the password field
	Phone             string    `json:"phone"`
	Role              string    `json:"role"`
	AccountState      string    `json:"account_state"`
	IncorrectPwdCount int       `json:"incorrect_pwd_count"`
	PwdResetFlag      int       `json:"pwd_reset_flag"`
	PwdResetCode      string    `json:"pwd_reset_code"`
	InviteCode        string    `json:"invite_code"`
	PhoneVerification string    `json:"phone_verification"`
	EmailVerification string    `json:"email_verification"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
