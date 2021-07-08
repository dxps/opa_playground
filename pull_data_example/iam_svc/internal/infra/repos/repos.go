package repos

import "database/sql"

type Repos struct {
	Subjects   SubjectRepo
	Attributes AttributeRepo
}

func New(db *sql.DB) Repos {
	return Repos{
		Subjects:   SubjectRepo{DB: db},
		Attributes: AttributeRepo{DB: db},
	}
}
