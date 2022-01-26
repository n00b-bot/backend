// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_passwd,
    full_name,
    email
) VALUES (
    $1,$2,$3,$4
) RETURNING username, hashed_passwd, full_name, email, password_changed_at, created_at
`

type CreateUserParams struct {
	Username     string `json:"username"`
	HashedPasswd string `json:"hashed_passwd"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.Username,
		arg.HashedPasswd,
		arg.FullName,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPasswd,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, hashed_passwd, full_name, email, password_changed_at, created_at FROM users WHERE username= $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPasswd,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}