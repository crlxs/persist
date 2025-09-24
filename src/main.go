package main

// https://dev.to/kengowada/go-routing-101-handling-and-grouping-routes-with-nethttp-4k0e
// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/

import (
	"fmt"
	"net/http"
	"log"
)

func main() {

	fmt.Println("Persist!")

	mux := http.NewServeMux() //Request multiplexer

	mux.HandleFunc("GET /exercises/", getExercises)
	// mux.HandleFunc("GET /sets/", getSets)
	// mux.HandleFunc("GET /workouts/", getWorkouts)
	// mux.HandleFunc("GET /splits/", getSplits)

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

// Handlers

func getExercises(w http.ResponseWriter, r *http.Request) {
	response := "You called the getExercises function."
	// utils.WriteJSONResponse(w, http.StatusOK, response)
}
