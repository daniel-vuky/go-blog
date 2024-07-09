package admin

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Gender string

const (
	Gender1 Gender = "1"
	Gender2 Gender = "2"
	Gender3 Gender = "3"
)

func (e *Gender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Gender(s)
	case string:
		*e = Gender(s)
	default:
		return fmt.Errorf("unsupported scan type for Gender: %T", src)
	}
	return nil
}

type NullGender struct {
	Gender Gender `json:"gender"`
	Valid  bool   `json:"valid"` // Valid is true if Gender is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullGender) Scan(value interface{}) error {
	if value == nil {
		ns.Gender, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Gender.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullGender) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Gender), nil
}

type User struct {
	UserID            int64              `json:"user_id"`
	Email             string             `json:"email"`
	Firstname         string             `json:"firstname"`
	Lastname          string             `json:"lastname"`
	Subscribe         pgtype.Bool        `json:"subscribe"`
	Gender            NullGender         `json:"gender"`
	Dob               pgtype.Timestamptz `json:"dob"`
	HashedPassword    string             `json:"hashed_password"`
	PasswordChangedAt time.Time          `json:"password_changed_at"`
	CreatedAt         time.Time          `json:"created_at"`
}
