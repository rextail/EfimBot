package sqlite3

import (
	"context"
	"database/sql"
)

type taskRepo struct {
	taskDB *sql.DB
}

func (t *taskRepo) Done(ctx context.Context) {

}
func (t *taskRepo) CreateExternal(ctx context.Context) {

}
func (t *taskRepo) AdressExternalTo(ctx context.Context) {

}
func (t *taskRepo) CreateInternalAttachedToUser(ctx context.Context) {

}
