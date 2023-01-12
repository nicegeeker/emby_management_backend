// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          int64     `json:"id"`
	UserName    string    `json:"userName"`
	OrderTypeID int64     `json:"orderTypeID"`
	Discount    float64   `json:"discount"`
	CreatedAt   time.Time `json:"createdAt"`
}

type OrderType struct {
	ID int64 `json:"id"`
	// must be positive
	Days  int64 `json:"days"`
	Price int64 `json:"price"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refreshToken"`
	UserAgent    string    `json:"userAgent"`
	ClientIp     string    `json:"clientIp"`
	IsBlocked    bool      `json:"isBlocked"`
	ExpiaresAt   time.Time `json:"expiaresAt"`
	CreatedAt    time.Time `json:"createdAt"`
}

type User struct {
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	HashedPassword    string    `json:"hashedPassword"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
	CreatedAt         time.Time `json:"createdAt"`
	ExpireAt          time.Time `json:"expireAt"`
}
