package repos

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dxps/opa_playground/pull_data_example/iam_svc/internal/app"
	"github.com/dxps/opa_playground/pull_data_example/iam_svc/internal/domain"
	"github.com/gofrs/uuid"
)

// SubjectRepo is an implementation of a repository of subjects.
type SubjectRepo struct {
	DB *sql.DB
}

// Add method inserts the subject into the repository, updated with the persistence details
// that is the internal and external IDs (IID, EID), CreatedAt, and Version.
func (sr SubjectRepo) Add(subj *domain.Subject) error {

	eid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	query := `
		INSERT INTO subjects (eid, name, email, password_hash, active)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING iid, created_at, version
	`
	args := []interface{}{eid.String(), subj.Name, subj.Email, subj.Password.Hash, subj.Active}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = sr.DB.QueryRowContext(ctx, query, args...).Scan(&subj.IID, &subj.CreatedAt, &subj.Version)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "subjects_email_key"`:
			return app.ErrDuplicateEmail
		default:
			return err
		}
	}
	subj.EID = eid.String()
	return nil
}

func (sr SubjectRepo) GetByEmail(email string) (*domain.Subject, error) {
	query := `
		SELECT iid, eid, created_at, name, email, password_hash, active, version
		FROM subjects
		WHERE email = $1
	`
	var subj domain.Subject

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := sr.DB.QueryRowContext(ctx, query, email).Scan(
		&subj.IID, &subj.EID, &subj.CreatedAt, &subj.Name, &subj.Email, &subj.Password.Hash, &subj.Active, &subj.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, app.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &subj, nil
}

func (sr SubjectRepo) GetSubjectIDByEID(eid string) (*int64, error) {
	query := `
		SELECT iid FROM subjects WHERE eid = $1
	`
	var id int64

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := sr.DB.QueryRowContext(ctx, query, eid).Scan(&id)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, app.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &id, nil
}
