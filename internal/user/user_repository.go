package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	// Methods
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

// Methods
func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastIserted int
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastIserted)
	if err != nil {
		return &User{}, err
	}

	user.ID = int64(lastIserted)
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}

	query := "SELECT id, username, email, password FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Username, &u.Email, &u.Password)
	if err != nil {
		return &User{}, err
	}

	return &u, nil
}
