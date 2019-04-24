// program Average1
package main

import "fmt"

func main() {
	var n,t int

	fmt.Scanln(&n)

	for i := 1; i <= 8; i++ {
		if t:= n%i == 0 {
			fmt.Println("factors of ",t,i)
		} else {
			fmt.Println("false",t,i)
			i = i + 1

		}
	}
}
