package repositories

import (
	"movies-app/models"
	"time"

	"gorm.io/gorm"
)

type MovieRepository struct {
	DB *gorm.DB
}

func (r *MovieRepository) GetList() ([]*models.Movie, error) {
	var movies []*models.Movie

	err := r.DB.Where("deleted_at IS NULL").Find(&movies).Error

	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *MovieRepository) GetById(id string) (*models.Movie, error) {
	var movie *models.Movie
	err := r.DB.Where("id = ?", id).First(&movie).Error

	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (repo *MovieRepository) Create(title string, description string, duration time.Duration) error {
	movie := models.Movie{
		Title:       title,
		Description: description,
		Duration:    &duration,
	}

	return repo.DB.Create(&movie).Error
}
