package main

// variable declaration
	var pi float64 = 3.14159


// short variable declaration

// example := ""

var isnumCars = 10 // inferred to be an integer

// Same declarations 
age, name := 86, "Albert"

// is the same as

age := 86
name := "Albert"

// Constants
const firstName = "Mohammed Firdous"	
const integer = 40

// Using the %v variant that prints the Go syntax representation of a value
s := fmt.Sprintf("I am %v years old", 50)
// I am 50 years old


// %s string format
s := fmt.Sprintf("We are going to be %s years old in a few years time", "fifty and fifty-one")
// We are going to be fifty and fifty-one years old in a few years time


// %d integer in decimal format
s := fmt.Sprintf("I am %d years of age", 50)
// I am 50 years of age


// %f decimal format
s := fmt.Sprintf("I am %f years old", 50.890)
// I am 50.890 years old


