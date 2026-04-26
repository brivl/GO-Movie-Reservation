package main

import (
	"log"
	"movies-app/handlers"
	"movies-app/models"
	"movies-app/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=myuser password=mypassword dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.Hall{}, &models.User{}, &models.Movie{}, &models.Showtime{}, &models.Reservation{}, &models.Seat{})

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Movie Reservation System!",
		})
	})

	userRepo := &repositories.UserRepository{DB: db}
	movieRepo := &repositories.MovieRepository{DB: db}

	userHandler := &handlers.UserHandler{UserRepo: userRepo}
	movieHandler := &handlers.MovieHandler{MovieRepo: movieRepo}

	router.POST("/register", userHandler.Register)
	router.POST("/admin/set", userHandler.SetAdmin)

	router.GET("/movies", movieHandler.GetList)
	router.GET("/movies/{id}", movieHandler.GetById)
	router.POST("/movies", movieHandler.Create)
	router.PUT("/movies/{id}", movieHandler.Update)
	router.DELETE("/movies/{id}", movieHandler.Delete)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
