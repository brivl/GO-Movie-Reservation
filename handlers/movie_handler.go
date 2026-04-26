package handlers

import (
	"movies-app/models"
	"movies-app/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	MovieRepo *repositories.MovieRepository
}

func (handler *MovieHandler) GetList(c *gin.Context) {
	movies, err := handler.MovieRepo.GetList()

	if err != nil {
		movies = []*models.Movie{}
	}

	c.JSON(http.StatusOK, movies)
}

func (handler *MovieHandler) GetById(c *gin.Context) {
	id, _ := c.Params.Get("id")

	movie, err := handler.MovieRepo.GetById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie Not found."})
	}

	c.JSON(http.StatusOK, movie)
}

func (handler *MovieHandler) Create(c *gin.Context) {
	// check admin rights
}

func (handler *MovieHandler) Update(c *gin.Context) {
	// check admin rights
}

func (handler *MovieHandler) Delete(c *gin.Context) {
	// check admin rights
}
