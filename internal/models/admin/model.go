package admin

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Admin struct {
	AdminID           int32              `json:"admin_id"`
	RoleID            int64              `json:"role_id"`
	Email             string             `json:"email"`
	HashedPassword    string             `json:"hashed_password"`
	Firstname         string             `json:"firstname"`
	Lastname          pgtype.Text        `json:"lastname"`
	Active            pgtype.Bool        `json:"active"`
	LockExpires       pgtype.Timestamptz `json:"lock_expires"`
	PasswordChangedAt time.Time          `json:"password_changed_at"`
	CreatedAt         time.Time          `json:"created_at"`
}

type AuthorizationRole struct {
	RoleID          int32     `json:"role_id"`
	RoleName        string    `json:"role_name"`
	IsAdministrator bool      `json:"is_administrator"`
	CreatedAt       time.Time `json:"created_at"`
}

type AuthorizationRule struct {
	RuleID         int64     `json:"rule_id"`
	RoleID         int64     `json:"role_id"`
	PermissionCode string    `json:"permission_code"`
	IsAllowed      bool      `json:"is_allowed"`
	CreatedAt      time.Time `json:"created_at"`
}

type CreateAdminParams struct {
	RoleID            int64              `json:"role_id"`
	Email             string             `json:"email"`
	HashedPassword    string             `json:"hashed_password"`
	Firstname         string             `json:"firstname"`
	Lastname          pgtype.Text        `json:"lastname"`
	Active            pgtype.Bool        `json:"active"`
	LockExpires       pgtype.Timestamptz `json:"lock_expires"`
	PasswordChangedAt time.Time          `json:"password_changed_at"`
}

type GetListAdminParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type UpdateAdminParams struct {
	Email             string             `json:"email"`
	RoleID            pgtype.Int8        `json:"role_id"`
	HashedPassword    pgtype.Text        `json:"hashed_password"`
	Firstname         pgtype.Text        `json:"firstname"`
	Lastname          pgtype.Text        `json:"lastname"`
	Active            pgtype.Bool        `json:"active"`
	LockExpires       pgtype.Timestamptz `json:"lock_expires"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
}
