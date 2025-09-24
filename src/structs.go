package main

// import "time"  //https://pkg.go.dev/time#Date
// import "civil"  //https://pkg.go.dev/cloud.google.com/civil#Date

// Exercise
type Exercise struct {
	Name string
	Targetmuscles []string // A slice of strings
	ID int // DB PrimaryKey
}

// Set
// type Set struct {
// 	Exercise int // ForeignKey of the exercise
// 	Weight float64
// 	WeightUnits string
// 	Repetitions int
// 	Comment string
// 	Rir string //Reps In Reserve, how close you brought the exercise to failure.
// 	ID int // DB PrimaryKey
// }

// Workout
// type Workout struct {
// 	Name string
// 	Sets []Set // A slice of ints, for the foreignkeys of the exercises.
// 	Duration int // In seconds, simply transform them to a more readable format when displyaed on the frontend.
// 	Date time.Date
// 	ID int // DB PrimaryKey
// }

// Split
// type Split struct {
// 	Name string
// 	Workouts []SplitEntry // To tie workouts to days
// 	Comment string
// 	ID int // DB PrimaryKey
// }

// SplitEntry
// type SplitEntry struct {
// 	WorkoutID int
// 	Day int
// 	ID int // DB PrimaryKey
// }
