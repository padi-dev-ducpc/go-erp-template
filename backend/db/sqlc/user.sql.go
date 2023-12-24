// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username,
  first_name,
  last_name,
  email,
  password,
  avatar,
  last_login,
  user_status_id,
  role_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING id, username, first_name, last_name, email, password, avatar, last_login, user_status_id, role_id, created_at, updated_at
`

type CreateUserParams struct {
	Username     string           `json:"username"`
	FirstName    string           `json:"first_name"`
	LastName     string           `json:"last_name"`
	Email        string           `json:"email"`
	Password     string           `json:"password"`
	Avatar       pgtype.Text      `json:"avatar"`
	LastLogin    pgtype.Timestamp `json:"last_login"`
	UserStatusID int32            `json:"user_status_id"`
	RoleID       int32            `json:"role_id"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
		arg.Avatar,
		arg.LastLogin,
		arg.UserStatusID,
		arg.RoleID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.LastLogin,
		&i.UserStatusID,
		&i.RoleID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, username, first_name, last_name, email, password, avatar, last_login, user_status_id, role_id, created_at, updated_at FROM users 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.LastLogin,
		&i.UserStatusID,
		&i.RoleID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT 
    users.id,
    users.username,
    users.first_name,
    users.last_name,
    users.email,
    users.password,
    users.avatar,
    users.last_login,
    users.user_status_id,
    users.role_id,
    users.created_at,
    users.updated_at,
    user_roles.role_name AS role_name,
    user_statuses.status_name AS status_name
FROM users 
JOIN user_roles ON users.role_id = user_roles.id
JOIN user_statuses ON users.user_status_id = user_statuses.id
WHERE users.email = $1 
LIMIT 1
`

type GetUserByEmailRow struct {
	ID           int32            `json:"id"`
	Username     string           `json:"username"`
	FirstName    string           `json:"first_name"`
	LastName     string           `json:"last_name"`
	Email        string           `json:"email"`
	Password     string           `json:"password"`
	Avatar       pgtype.Text      `json:"avatar"`
	LastLogin    pgtype.Timestamp `json:"last_login"`
	UserStatusID int32            `json:"user_status_id"`
	RoleID       int32            `json:"role_id"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
	RoleName     string           `json:"role_name"`
	StatusName   string           `json:"status_name"`
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.LastLogin,
		&i.UserStatusID,
		&i.RoleID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RoleName,
		&i.StatusName,
	)
	return i, err
}

const getUserForUpdate = `-- name: GetUserForUpdate :one
SELECT id, username, first_name, last_name, email, password, avatar, last_login, user_status_id, role_id, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetUserForUpdate(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUserForUpdate, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.LastLogin,
		&i.UserStatusID,
		&i.RoleID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, username, first_name, last_name, email, password, avatar, last_login, user_status_id, role_id, created_at, updated_at FROM users
LIMIT $2
OFFSET $1
`

type GetUsersParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsers, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Password,
			&i.Avatar,
			&i.LastLogin,
			&i.UserStatusID,
			&i.RoleID,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateLastLogin = `-- name: UpdateLastLogin :exec
UPDATE users
SET
  last_login = NOW()
WHERE id = $1
`

func (q *Queries) UpdateLastLogin(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, updateLastLogin, id)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET
  username = $2,
  first_name = $3,
  last_name = $4,
  email = $5,
  password = $6,
  user_status_id = $7,
  role_id = $8,
  avatar = $9,
  updated_at = NOW()
WHERE id = $1
`

type UpdateUserParams struct {
	ID           int32       `json:"id"`
	Username     string      `json:"username"`
	FirstName    string      `json:"first_name"`
	LastName     string      `json:"last_name"`
	Email        string      `json:"email"`
	Password     string      `json:"password"`
	UserStatusID int32       `json:"user_status_id"`
	RoleID       int32       `json:"role_id"`
	Avatar       pgtype.Text `json:"avatar"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.ID,
		arg.Username,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
		arg.UserStatusID,
		arg.RoleID,
		arg.Avatar,
	)
	return err
}
