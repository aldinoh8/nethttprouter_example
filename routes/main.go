package routes

import (
	"database/sql"
	"p2httprouter/controller"
	"p2httprouter/models"

	"github.com/julienschmidt/httprouter"
)

func New(db *sql.DB) *httprouter.Router {
	router := httprouter.New()

	movieRepository := models.NewMovieRepository(db)
	controller := controller.New(movieRepository)

	router.POST("/movies", controller.Create)
	router.GET("/movies", controller.Index)
	router.GET("/movies/:id", controller.Detail)
	router.DELETE("/movies/:id", controller.Delete)

	return router
}
