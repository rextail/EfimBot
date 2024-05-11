package sqlite3

import (
	"EfimBot/internal/models"
	"context"
	"database/sql"
)

type projectRepo struct {
	projectDB *sql.DB
}

func (p *projectRepo) Create(ctx context.Context, project models.Project) error {
	query := `INSERT INTO projects(code_id,project_name,project_manager) VALUES(?,?,?)`

	args := []string{project.Code, project.Name, project.ProjectManager}

	_, err := p.projectDB.ExecContext(ctx, query, args)
	if err != nil {
		return err
	}

	return nil
}

func (p *projectRepo) GetID(ctx context.Context, project models.Project) (ID int, err error) {
	query := `SELECT project_id FROM projects WHERE code_id = ?`
	args := project.Code

	row := p.projectDB.QueryRowContext(ctx, query, args)

	err = row.Scan(&ID)
	if err != nil {
		return -1, sql.ErrNoRows
	}

	return ID, nil
}

func (p *projectRepo) GetByID(ctx context.Context, ID int) (project models.Project, err error) {
	query := `SELECT code_id, project_name, project_manager FROM projects WHERE project_id = ?`

	err = p.projectDB.QueryRowContext(ctx, query, ID).Scan(
		&project.Code,
		&project.Name,
		&project.ProjectManager,
	)
	if err != nil {
		return project, err
	}

	return project, nil
}
