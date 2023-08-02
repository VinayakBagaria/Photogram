package repository

import (
	"database/sql"

	"github.com/VinayakBagaria/go-cat-pictures/db"
)

type PicturesRepository interface {
	Save(*db.Picture) error
	GetById(string) (*db.Picture, error)
	GetAll() ([]*db.Picture, error)
	Update(*db.Picture) error
	Delete(id string) error
}

type picturesRepository struct {
	db *sql.DB
}

func NewPicturesRepository(dbHandler *sql.DB) PicturesRepository {
	return &picturesRepository{db: dbHandler}
}

func (p *picturesRepository) Save(picture *db.Picture) error {
	return nil
}

func (p *picturesRepository) GetById(id string) (*db.Picture, error) {
	return nil, nil
}

func (p *picturesRepository) GetAll() ([]*db.Picture, error) {
	rows, err := p.db.Query("SELECT ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pictures := []*db.Picture{}
	for rows.Next() {
		var p *db.Picture
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, err
		}

		pictures = append(pictures, p)
	}

	return pictures, nil
}

func (p *picturesRepository) Update(picture *db.Picture) error {
	return nil
}

func (p *picturesRepository) Delete(id string) error {
	return nil
}
