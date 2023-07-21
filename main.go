package main

import (
	"fmt"
	"log"
	"net/http"
	"p2httprouter/config"
	"p2httprouter/routes"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("HTTP request sent to %s from %s", r.URL.Path, r.Method)
		next.ServeHTTP(w, r)
	})
}

func main() {
	db := config.InitializeDb("mysql:mysql@tcp(localhost:3306)/ftgo_p2_week1_db")
	router := routes.New(db)

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprintf(w, "Some error %v", i)
	}

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: Logger(router),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
