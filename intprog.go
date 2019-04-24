/* SID: 1301184423
   Name: Diya namira Purba
   thhis program it is for loookingsome data on the array
*/

package main

import "fmt"

const nMax = 21

type intArray struct {
	tabInt [nMax]int
	N      int
}

func main() {
	var T, R intArray
	var sort, search string
	var x int
	var found bool
	fillArray(&T, &R)
	printData(T, R)
	fmt.Scanln(&sort)
	switch sort {
	case "insertsionsort":
		insertsionsort(&T)
	case "selectionSort":
		selectionSort(&T)
	}
	printData(T, R)
	fmt.Scanln(&search, &x)
	switch search {
	case "SearchA":
		found = SearchA(T, x)
	case "SearchB":
		found = SearchB(T, x)
	case "binarysearch":
		found = binarysearch(T, x)
	}
	if found {
		fmt.Println("FOUND!")
	} else {
		fmt.Println("NOT FOUND!")
	}
}

/*
Process: Accept N from user, then get N integers from user
  F.S. tableInt contains N data, 1 < N < nMax
*/
func fillArray(T1, T2 *intArray) {
	i := 0
	fmt.Scanln(&T1.N)
	for i < T1.N {
		fmt.Scanln(&T1.tabInt[i])
		i = i + 1
	}
	T2 = T1
}

/*
I.S. tableInt contains N data, 1 < N < nMax
  F.S. all elements in tableInt are printed
*/
func printData(T1, T2 intArray) {
	i := 0
	for i < T1.N {
		fmt.Print(T1.tabInt[i])
		fmt.Print(T2.tabInt[i])
		i = i + 1
	}
}

/*
I.S. tab contains N elements.
  F.S. will return true if X in the table tab, false if not
*/
func SearchA(tableInt intArray, x int) bool {
	i := 0
	found := false
	for i < tableInt.N && !found {
		if tableInt.tabInt[i] == x {
			found = true
		} else {
			i = i + 1
		}
	}
	return found
}

/*
I.S. tab contains N elements.
  F.S. will return true if X in the table tab, false if not
*/
func SearchB(tableInt intArray, x int) bool {
	i := 0
	for i < tableInt.N && tableInt.tabInt[i] != x {
		i = i + 1
	}
	return tableInt.tabInt[i] == x
}

/*
I.S looK for the data using  the binary search
F.S will output the if it is found
*/

func binarysearch(tableInt intArray, x int) bool {
	left := 0
	right := tableInt.N
	found := false
	for left <= right && !found {
		mid := (left + right) / 2
		if tableInt.tabInt[mid] == x {
			found = true
		} else if tableInt.tabInt[mid] > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return found
}

/*
I.S will sort the data
F.S will output the data if it sorted
*/
func selectionSort(tableInt *intArray) {
	j := tableInt.N - 1
	for j > 1 {
		max := 0
		i := 1
		for i <= j {
			if tableInt.tabInt[i] < tableInt.tabInt[max] {
				max = i
			}
			i = i + 1
		}
		temp := tableInt.tabInt[max]
		tableInt.tabInt[max] = tableInt.tabInt[j]
		tableInt.tabInt[j] = temp
		j = j - 1
	}
}

/*
I.S. tableInt contains N elements, 1 < N < nMax
  F.S. The content of tableInt should be in order from
*/

func insertsionsort(tableInt *intArray) {
	j := 1
	for j <= tableInt.N-1 {
		temp := tableInt.tabInt[j]
		i := j - 1
		for i >= 0 && tableInt.tabInt[i] > temp {
			tableInt.tabInt[i+1] = tableInt.tabInt[i]
			i = i - 1
		}
		tableInt.tabInt[i+1] = temp
		j = j + 1
	}
}

/*
I.S The range from the data
F.S output the calculate data
*/

func Range(T1, T2 *intArray) bool {
	var sumA, sumB int
	i := 0
	for i < T1.N {
		sumA = sumA + T1.tabInt[i]
		sumB = sumB + T2.tabInt[i]
		i = i + 1
	}
	return sumA == sumB
}

/*
I.S Search the data if it found or not
F.S if the data found will be return true
*/
func itFound(T1, T2 *intArray) bool {
	i := 0
	isitFound := true
	for i < T1.N && isitFound {
		if T1.tabInt[i] != T2.tabInt[i] {
			isitFound = false
		}
		i = i + 1
	}
	return isitFound
}
