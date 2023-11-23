package models

type User struct {
	ID                int               `json:"id"`
	FirstName         string            `json:"first_name"`
	LastName          string            `json:"last_name"`
	UserName          string            `json:"user_name"`
	Email             string            `json:"email"`
	Password          string            `json:"-"` // omit the password field
	Phone             string            `json:"phone"`
	Role              string            `json:"role"`
	AccountState      string            `json:"account_state"`
	Teams             []Team            `json:"teams"`
	Projects          []Project         `json:"projects"`
	IncorrectPwdCount int               `json:"incorrect_pwd_count"`
	PwdResetFlag      bool              `json:"pwd_reset_flag"`
	PwdResetCode      string            `json:"pwd_reset_code"`
	InviteCode        string            `json:"invite_code"`
	PhoneVerification phoneVerification `json:"phone_verification"`
	EmailVerification emailVerification `json:"email_verification"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
}

type phoneVerification struct {
	Code     string `json:"code"`
	Verified bool   `json:"verified"`
}

type emailVerification struct {
	Code     string `json:"code"`
	Verified bool   `json:"verified"`
}
