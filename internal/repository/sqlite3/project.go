package sqlite3

import (
	"EfimBot/internal/models"
	"EfimBot/pkg/e"
	"context"
	"database/sql"
)

type projectRepo struct {
	projectDB *sql.DB
}

func (p *projectRepo) Create(ctx context.Context, project models.Project) error {
	query := `INSERT INTO projects(code_id,project_name,project_manager) VALUES(?,?,?)`

	_, err := p.projectDB.ExecContext(ctx, query, project.Code, project.Name, project.ProjectManager)
	if err != nil {
		return err
	}

	return nil
}

func (p *projectRepo) GetID(ctx context.Context, project models.Project) (ID int, err error) {
	query := `SELECT project_id FROM projects WHERE code_id = ?`

	err = p.projectDB.QueryRowContext(ctx, query, project.Code).Scan(&ID)
	if err != nil {
		return -1, err
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

func (p *projectRepo) Complete(ctx context.Context, projectID int) error {
	query := `DELETE FROM projects WHERE project_id = ?`

	_, err := p.projectDB.ExecContext(ctx, query, projectID)
	if err != nil {
		return e.Wrap("can't make project completed", err)
	}

	return nil
}
