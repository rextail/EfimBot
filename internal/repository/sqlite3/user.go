package sqlite3

import (
	"EfimBot/internal/models"
	"context"
	"database/sql"
)

type userRepo struct {
	userDB *sql.DB
}

func (u *userRepo) Create(ctx context.Context, user models.User) error {
	query := `INSERT INTO users VALUES(?,?,?,?)`

	args := []string{user.Name, user.Position, user.Department, user.SubDepartment}

	_, err := u.userDB.ExecContext(ctx, query, args)
	if err != nil {
		return err
	}

	return nil
}
func (u *userRepo) Delete(ctx context.Context, username string, subdepartment string) error {
	query := `DELETE FROM users WHERE user_name = ? AND sub_department = ?`

	args := []string{username, subdepartment}

	_, err := u.userDB.ExecContext(ctx, query, args)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepo) GetID(ctx context.Context, username string, subdepartment string) (ID int, err error) {
	query := `SELECT user_id FROM users WHERE user_name = ? AND sub_department = ?`

	args := []string{username, subdepartment}

	err = u.userDB.QueryRowContext(ctx, query, args).Scan(&ID)
	if err != nil {
		return -1, err
	}

	return ID, nil
}
