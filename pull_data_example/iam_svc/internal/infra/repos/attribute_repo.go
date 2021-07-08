package repos

import (
	"context"
	"database/sql"
	"time"

	"github.com/dxps/opa_playground/pull_data_example/iam_svc/internal/domain"
)

// AttributeRepo is an implementation of a repository of attributes.
type AttributeRepo struct {
	DB *sql.DB
}

func (ar AttributeRepo) Add(attr domain.Attribute) error {

	query := `
		INSERT INTO attributes (owner_id, owner_type, name, value)
		VALUES ($1, $2, $3, $4)
	`
	args := []interface{}{
		attr.OwnerID, attr.OwnerType, attr.Name, attr.Value,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := ar.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (ar AttributeRepo) GetAllAttributesBySubjectID(id int64) ([]*domain.Attribute, error) {

	query := `
		SELECT name, value FROM attributes 
		WHERE owner_id=$1
	`
	args := []interface{}{id}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := ar.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	attrs := []*domain.Attribute{}

	for rows.Next() {
		var attr domain.Attribute
		err := rows.Scan(&attr.Name, &attr.Value)
		if err != nil {
			return nil, err
		}
		attrs = append(attrs, &attr)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return attrs, nil
}

func (ar AttributeRepo) GetAllAttributesBySubjectEID(eid string) ([]*domain.Attribute, error) {

	query := `
		SELECT name, value FROM attributes 
		WHERE owner_id IN (SELECT iid FROM subjects WHERE eid=$1)
	`
	args := []interface{}{eid}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := ar.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	attrs := []*domain.Attribute{}

	for rows.Next() {
		var attr domain.Attribute
		err := rows.Scan(&attr.Name, &attr.Value)
		if err != nil {
			return nil, err
		}
		attrs = append(attrs, &attr)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return attrs, nil
}
