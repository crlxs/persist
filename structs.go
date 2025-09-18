// import "time"

import "civil"  //https://pkg.go.dev/cloud.google.com/civil#Date

// Exercise
type Exercise struct {
	name string
	targetmuscles []string // A slice of strings
}

// Set
type Set struct {
	exercise string
	repetitions int
	comment string
	rir string //Reps In Reserve, how close you brought the exercise to failure.
}

// Workout
type Workout struct {
	name string
	exercises []string // A slice of strings
	duration int
	date 
}

// Split
type Split struct {
	name string
}
