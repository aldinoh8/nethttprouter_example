package models

import (
	"context"
	"database/sql"
	"errors"
)

type MovieRepository struct {
	DB *sql.DB
}

func NewMovieRepository(db *sql.DB) MovieRepository {
	return MovieRepository{DB: db}
}

func (mr MovieRepository) FindAll() []Movie {
	ctx := context.Background()
	query := `
		SELECT id, title, rating FROM movies
	`
	movies := []Movie{}
	rows, err := mr.DB.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		m := Movie{}
		err := rows.Scan(&m.Id, &m.Title, &m.Rating)
		if err != nil {
			panic(err)
		}
		movies = append(movies, m)
	}

	return movies
}

func (mr MovieRepository) FindByPk(id int) (Movie, error) {
	ctx := context.Background()
	query := `
		SELECT id, title, rating FROM movies
		WHERE id = ?
	`
	movie := Movie{}
	rows, err := mr.DB.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Rating)
		if err != nil {
			panic(err)
		}
	} else {
		return movie, errors.New("data not found")
	}

	return movie, nil
}

func (mr MovieRepository) Create(newMovie *Movie) error {
	ctx := context.Background()
	query := `
		INSERT INTO movies (title, rating)
		VALUES (?, ?)
	`
	result, err := mr.DB.ExecContext(ctx, query, newMovie.Title, newMovie.Rating)
	if err != nil {
		return err
	}

	newId, _ := result.LastInsertId()
	newMovie.Id = int(newId)

	return nil
}

func (mr MovieRepository) Delete(id int) error {
	ctx := context.Background()
	query := `
		DELETE from movies
		WHERE id = ?
	`

	result, err := mr.DB.ExecContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	affectedRows, err := result.RowsAffected()
	if affectedRows == 0 {
		return errors.New("data not found")
	}

	return nil
}
