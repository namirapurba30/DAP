package main

import "fmt", "math"
func main () {
	var co1, co2, co3, co4 float64

	fmt.Print("co1")
	fmt.Scan(&co1)
	fmt.Print("co2")
	fmt.Scan(&co2)
	fmt.Print("co3")
	fmt.Scan(&co3)
	fmt.Print("co4")
	fmt.Scan("&co4")

distance:= sqrt((co2-co1)**2 + (co4-co3)**2)
wide := (co2 - co1) * (co3-co2)
around := (co2-co1) + (co3-co2) + (co4-co3) + (co4-co1)

fmt.Println("wide")
fmt.Println("around")
}