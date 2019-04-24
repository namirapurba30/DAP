/*
SID:1301184423
Name: diya namira
the program is doing calenderdays
*/

package main

import "fmt"

func main() {
	var year int

	fmt.Print("enter a year :")
	fmt.Scan(&year)
	if year%400 == 0 {
		fmt.Print("True")
	} else if year%100 == 0 {
		fmt.Print("False")
	} else if year%4 == 0 {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}
}
