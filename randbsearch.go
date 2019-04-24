/* SID: 1301184423
Name: Diya Namira
This program is doing a searching that to use to found a data on the user input/array
*/
package main

import "fmt"
import "math/rand"

const nMax = 1000003

type intRecArray struct {
	number int
}

type intArray [nMax]intRecArray

/* I.S. : This is a main function to call the procedure
   F.S. : after call the procedurethe program will return if it's found or not
*/
func main() {
	var data intArray
	var x, i int
	fillArray(&data)
	x = nextValue()
	for x >= 0 {
		dt := binaryFind(data, x)
		for x > 0 && x <= nMax {
			if dt == -1 {
				fmt.Println("Not FOUND!")
			} else {
				fmt.Println("FOUND!")
				i++
			}
			x = nextValue()
		}
	}
}

/* I.S. : This is a procedure to input a number on the array
   F.S. : will input the number into the array
*/
func fillArray(T *intArray) {
	begin := rand.Intn(1013)
	for i := 0; i < nMax; i++ {
		if i == 0 {
			T[i].number = begin
		} else {
			T[i].number = T[i-1].number + begin/2
		}
	}
}

/* I.S. : this functiom it use to find the value with binary
   F.S. : will return the found value
*/
func binaryFind(T intArray, x int) int {
	right := nMax
	left := 0
	dt := -1
	for left <= right && dt == -1 {
		mid := (left + right) / 2
		if T[mid].number == x {
			dt = mid
		} else if T[mid].number > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return dt
}

/* I.S. : This is a function ask for the number to found
   F.S. : will return the result of the input
*/
func nextValue() int {
	var x int
	fmt.Print("Input an integer? ")
	fmt.Scanln(&x)
	return x
}