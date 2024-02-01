package service

import (
	"github.com/Kartochnik010/test-effectivemobile-jan/internal/models"
	repository "github.com/Kartochnik010/test-effectivemobile-jan/internal/repository"
)

type PersonService struct {
	repo repository.Person // пока что эти методы никак не отличаются
}

func NewPersonService(repo repository.Person) *PersonService {
	return &PersonService{repo: repo}
}

func (p *PersonService) InsertPerson(person models.Person) (models.Person, error) {
	person, err := p.repo.InsertPerson(person)
	if err != nil {
		return models.Person{}, err
	}

	return person, nil
}
func (p *PersonService) FindPersonById(id int) (models.Person, error) {
	person, err := p.repo.FindPersonById(id)
	if err != nil {
		return models.Person{}, err
	}

	return person, nil
}
func (p *PersonService) DeletePersonById(id int) error {
	return p.repo.DeletePersonById(id)
}
func (p *PersonService) UpdatePersonById(id int, person models.Person) (models.Person, error) {
	updatedPerson, err := p.repo.UpdatePersonById(id, person)
	if err != nil {
		return models.Person{}, err
	}

	// Return the updated person
	return updatedPerson, nil
}
func (p *PersonService) SearchPerson(filters models.Filters) ([]models.Person, models.Metadata, error) {
	return nil, models.Metadata{}, nil
}
