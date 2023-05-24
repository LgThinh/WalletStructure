package model

import "pkg/mod/github.com/google/uuid@v1.3.0"

type User struct {
	BaseModel
	ID       uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Wallets  []Wallet  `json:"wallets"`
}
