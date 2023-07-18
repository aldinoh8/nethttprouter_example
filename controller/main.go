package controller

import (
	"encoding/json"
	"net/http"
	"p2httprouter/models"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type Conroller struct {
	movieRepository models.MovieRepository
}

func New(mr models.MovieRepository) Conroller {
	return Conroller{movieRepository: mr}
}

type ErrorMessage struct {
	Message string
}

func (c Conroller) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	newMovie := models.Movie{}
	err := decoder.Decode(&newMovie)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Message: "failed to decode request body"})
		return
	}

	err = c.movieRepository.Create(&newMovie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "success create movie"})
}

func (c Conroller) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	movies := c.movieRepository.FindAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func (c Conroller) Detail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idStr := p.ByName("id")
	id, _ := strconv.Atoi(idStr)
	movie, err := c.movieRepository.FindByPk(id)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorMessage{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func (c Conroller) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idStr := p.ByName("id")
	id, _ := strconv.Atoi(idStr)
	w.Header().Set("Content-Type", "application/json")

	err := c.movieRepository.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "success delete movie"})
}
