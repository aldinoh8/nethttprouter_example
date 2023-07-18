package main

import (
	"fmt"
	"net/http"
	"p2httprouter/config"
	"p2httprouter/routes"
)

func main() {
	db := config.InitializeDb("mysql:mysql@tcp(localhost:3306)/ftgo_p2_week1_db")
	router := routes.New(db)

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprintf(w, "Some error %v", i)
	}

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
