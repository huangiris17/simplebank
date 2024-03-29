// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"time"

	"github.com/google/uuid"
)

type Accounts struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
}

type Entries struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"accountId"`
	// can be negative or positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type Sessions struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refreshToken"`
	UserAgent    string    `json:"userAgent"`
	ClientIp     string    `json:"clientIp"`
	IsBlocked    bool      `json:"isBlocked"`
	ExpiresAt    time.Time `json:"expiresAt"`
	CreatedAt    time.Time `json:"createdAt"`
}

type Transfers struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"fromAccountId"`
	ToAccountID   int64 `json:"toAccountId"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type Users struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashedPassword"`
	FullName          string    `json:"fullName"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
	CreatedAt         time.Time `json:"createdAt"`
	IsEmailVerified   bool      `json:"isEmailVerified"`
}

type VerifyEmails struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	SecretCode string    `json:"secretCode"`
	IsUsed     bool      `json:"isUsed"`
	CreatedAt  time.Time `json:"createdAt"`
	ExpiredAt  time.Time `json:"expiredAt"`
}
