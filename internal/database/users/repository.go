package users

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"gitlab.com/robotomize/gb-golang/homework/03-01-umanager/internal/database"
)

func New(userDB *pgx.Conn, timeout time.Duration) *Repository {
	return &Repository{userDB: userDB, timeout: timeout}
}

type Repository struct {
	userDB  *pgx.Conn
	timeout time.Duration
}

func (r *Repository) Create(ctx context.Context, req CreateUserReq) (database.User, error) {
	var u database.User

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `INSERT INTO users (id, username, password, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW())
        ON CONFLICT (username) DO UPDATE
        SET password = EXCLUDED.password, updated_at = NOW()
        RETURNING id, username, password, created_at, updated_at;`

	resp := r.userDB.QueryRow(ctx, query, req.ID, req.Username, req.Password)
	if err := resp.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return database.User{}, err
	}

	return u, nil
}

func (r *Repository) FindByID(ctx context.Context, userID uuid.UUID) (database.User, error) {
	var u database.User

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `SELECT * FROM users WHERE id = $1;`
	resp := r.userDB.QueryRow(ctx, query, userID)
	if err := resp.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return database.User{}, err
	}

	return u, nil
}

func (r *Repository) FindByUsername(ctx context.Context, username string) (database.User, error) {
	var u database.User

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := `SELECT * FROM users WHERE username = $1;`
	resp := r.userDB.QueryRow(ctx, query, username)
	if err := resp.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return database.User{}, err
	}

	return u, nil
}
