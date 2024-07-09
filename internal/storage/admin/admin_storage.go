package admin

import (
	"context"
	model "github.com/daniel-vuky/go-blog/internal/models/admin"
	"github.com/daniel-vuky/go-blog/internal/storage"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository
// Wraps the Queries struct from the storage package and a connection pool.
type Repository struct {
	*storage.Queries
	connPool *pgxpool.Pool
}

// NewAdminRepository
// Returns a new instance of Repository.
// @param connPool *pgxpool.Pool
// @return *Repository
func NewAdminRepository(connPool *pgxpool.Pool) *Repository {
	return &Repository{
		Queries:  storage.New(connPool),
		connPool: connPool,
	}
}

const createAdmin = `-- name: CreateAdmin :one
INSERT INTO admin
    (
         role_id,
         email,
         hashed_password,
         firstname,
         lastname,
         active,
         lock_expires,
         password_changed_at
     )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING admin_id, role_id, email, hashed_password, firstname, lastname, active, lock_expires, password_changed_at, created_at
`

// Create
// Creates a new admin.
// @param ctx context.Context
// @param arg *model.CreateAdminParams
// @return model.Admin
func (repo *Repository) Create(
	ctx context.Context,
	arg *model.CreateAdminParams,
) (model.Admin, error) {
	row := repo.connPool.QueryRow(ctx, createAdmin,
		arg.RoleID,
		arg.Email,
		arg.HashedPassword,
		arg.Firstname,
		arg.Lastname,
		arg.Active,
		arg.LockExpires,
		arg.PasswordChangedAt,
	)
	var i model.Admin
	err := row.Scan(
		&i.AdminID,
		&i.RoleID,
		&i.Email,
		&i.HashedPassword,
		&i.Firstname,
		&i.Lastname,
		&i.Active,
		&i.LockExpires,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAdmin = `-- name: DeleteAdmin :one
DELETE FROM admin
WHERE email = $1
RETURNING admin_id, role_id, email, hashed_password, firstname, lastname, active, lock_expires, password_changed_at, created_at
`

// Delete
// Deletes an admin.
// @param ctx context.Context
// @param email string
// @return model.Admin
func (repo *Repository) Delete(
	ctx context.Context,
	email string,
) (model.Admin, error) {
	row := repo.connPool.QueryRow(ctx, deleteAdmin, email)
	var i model.Admin
	err := row.Scan(
		&i.AdminID,
		&i.RoleID,
		&i.Email,
		&i.HashedPassword,
		&i.Firstname,
		&i.Lastname,
		&i.Active,
		&i.LockExpires,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getAdmin = `-- name: GetAdmin :one
SELECT admin_id, role_id, email, hashed_password, firstname, lastname, active, lock_expires, password_changed_at, created_at
FROM admin
WHERE email = $1
`

// Get
// Returns an admin by email.
// @param ctx context.Context
// @param email string
// @return model.Admin
func (repo *Repository) Get(
	ctx context.Context,
	email string,
) (model.Admin, error) {
	row := repo.connPool.QueryRow(ctx, getAdmin, email)
	var i model.Admin
	err := row.Scan(
		&i.AdminID,
		&i.RoleID,
		&i.Email,
		&i.HashedPassword,
		&i.Firstname,
		&i.Lastname,
		&i.Active,
		&i.LockExpires,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getListAdmin = `-- name: GetListAdmin :many
SELECT admin_id, role_id, email, hashed_password, firstname, lastname, active, lock_expires, password_changed_at, created_at
FROM admin
LIMIT $1 OFFSET $2
`

// GetList
// Returns a list of admins.
// @param ctx context.Context
// @param arg *model.GetListAdminParams
// @return []model.Admin
func (repo *Repository) GetList(
	ctx context.Context,
	arg *model.GetListAdminParams,
) ([]model.Admin, error) {
	rows, err := repo.connPool.Query(ctx, getListAdmin, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.Admin
	for rows.Next() {
		var i model.Admin
		if err := rows.Scan(
			&i.AdminID,
			&i.RoleID,
			&i.Email,
			&i.HashedPassword,
			&i.Firstname,
			&i.Lastname,
			&i.Active,
			&i.LockExpires,
			&i.PasswordChangedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAdmin = `-- name: UpdateAdmin :one
UPDATE admin
SET role_id = COALESCE($2, role_id),
    hashed_password = COALESCE($3, hashed_password),
    firstname = COALESCE($4, firstname),
    lastname = COALESCE($5, lastname),
    active = COALESCE($6, active),
    lock_expires = COALESCE($7, lock_expires),
    password_changed_at = COALESCE($8, password_changed_at)
WHERE email = $1
RETURNING admin_id, role_id, email, hashed_password, firstname, lastname, active, lock_expires, password_changed_at, created_at
`

// Update
// Updates an admin.
// @param ctx context.Context
// @param arg *model.UpdateAdminParams
// @return model.Admin
func (repo *Repository) Update(
	ctx context.Context,
	arg *model.UpdateAdminParams,
) (model.Admin, error) {
	row := repo.connPool.QueryRow(ctx, updateAdmin,
		arg.Email,
		arg.RoleID,
		arg.HashedPassword,
		arg.Firstname,
		arg.Lastname,
		arg.Active,
		arg.LockExpires,
		arg.PasswordChangedAt,
	)
	var i model.Admin
	err := row.Scan(
		&i.AdminID,
		&i.RoleID,
		&i.Email,
		&i.HashedPassword,
		&i.Firstname,
		&i.Lastname,
		&i.Active,
		&i.LockExpires,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
