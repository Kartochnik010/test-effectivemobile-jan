package repository

import (
	"github.com/Kartochnik010/test-effectivemobile-jan/internal/models"
	"github.com/Kartochnik010/test-effectivemobile-jan/internal/repository/postgres"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Person
}

type Person interface {
	InsertPerson(person models.Person) (models.Person, error)
	FindPersonById(id int) (models.Person, error)
	DeletePersonById(id int) error
	UpdatePersonById(id int, person models.Person) (models.Person, error)
	SearchPerson(filters models.Filters) ([]models.Person, models.Metadata, error)
}

func NewPostgresRepo(db *pgx.Conn) *Repository {
	return &Repository{
		Person: postgres.NewPersonRepo(db),
	}
}
