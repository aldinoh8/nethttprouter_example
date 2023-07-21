package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"p2httprouter/controller"
	"p2httprouter/models"

	"github.com/julienschmidt/httprouter"
)

func Middleware1(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("middleware 1 executed")
		n(w, r, p)
	}
}

func Middleware2(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("middleware 2 executed")
		n(w, r, p)
	}
}

func New(db *sql.DB) *httprouter.Router {
	router := httprouter.New()

	movieRepository := models.NewMovieRepository(db)
	controller := controller.New(movieRepository)

	router.POST("/movies", controller.Create)
	router.GET("/movies", Middleware1(Middleware2(controller.Index)))
	router.GET("/movies/:id", controller.Detail)
	router.DELETE("/movies/:id", controller.Delete)

	return router
}
