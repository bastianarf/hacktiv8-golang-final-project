package responses

import "time"

type UserRegister struct {
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type UserUpdate struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       uint      `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLogin struct {
	Token string `json:"token"`
}
