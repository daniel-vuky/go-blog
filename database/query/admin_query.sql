-- name: GetAdmin :one
SELECT *
FROM admin
WHERE email = $1;

-- name: GetListAdmin :many
SELECT *
FROM admin
LIMIT $1 OFFSET $2;

-- name: CreateAdmin :one
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
RETURNING *;

-- name: UpdateAdmin :one
UPDATE admin
SET role_id = COALESCE(sqlc.narg(role_id), role_id),
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    firstname = COALESCE(sqlc.narg(firstname), firstname),
    lastname = COALESCE(sqlc.narg(lastname), lastname),
    active = COALESCE(sqlc.narg(active), active),
    lock_expires = COALESCE(sqlc.narg(lock_expires), lock_expires),
    password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at)
WHERE email = $1
RETURNING *;

-- name: DeleteAdmin :one
DELETE FROM admin
WHERE email = $1
RETURNING *;