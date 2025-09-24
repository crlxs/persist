package main

// https://dev.to/kengowada/go-routing-101-handling-and-grouping-routes-with-nethttp-4k0e
// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

func main() {

	fmt.Println("Persist!")

	mux := http.NewServeMux() //Request multiplexer

	mux.HandleFunc("GET /exercise/", getExercise)
	// mux.HandleFunc("GET /sets/", getSets)
	// mux.HandleFunc("GET /workouts/", getWorkouts)
	// mux.HandleFunc("GET /splits/", getSplits)

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

// Handlers
// A proper HTTP response has 3 key parts: a status code, header and a body.

func getExercise(w http.ResponseWriter, r *http.Request) {
	// response := "You called the getExercise function."
	// utils.WriteJSONResponse(w, http.StatusOK, response)
	// fmt.Println("getExercise")
	
	e := Exercise{Name: "Leg Press", Targetmuscles: []string{"Glutes", "Quadriceps"}, ID: 12}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(e)
}






