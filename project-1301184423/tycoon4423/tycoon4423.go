/* SID:1301184423
Name: Diya Namira Purba
This program is doing the tradig or as known as tycoon, that will let user will play some trade game
*/
package main

import "fmt"
import "math/rand"

var balance int
var material [4]string = [4]string{"silk", "silver", "gold", "diamond"}
var sum [4]int = [4]int{0, 0, 0, 0}
var buyerPrice [4]int = [4]int{0, 0, 0, 0}
var priceNow [4]int = [4]int{100, 200, 400, 600}
var name, s1, s2 string
var num, n int
var array [100]int
var day = 1
var warehouse [4]string
var warprice = 50
var rent bool = false
var logmoney [1000]float64
var longday [1000]int
var initialaccount int
var j int = 1

/* I.S. : Accept the prompt from the user
   F.S. : will output what the user prompt
*/

func prompt(s1 *string, i *int, s2 *string) {
	fmt.Print("Day[", day, "]", "What's your command?")
	fmt.Scanln(s1, i, s2)
	if *s1 == "show" {
		fmt.Print("account or catalog? ")
		fmt.Scanf("%s\n", s2)
	}

}

/* I.S. : the procedure will work when user prompt "buy"
   F.S. : will output the program if it prompt to buy
*/

func buy(i2 int, c string) {
	var y string
	switch c {
	case "silk":
		if balance > i2*priceNow[0] {
			fmt.Print("Buying ", i2, " ", c, " bought at ", i2*priceNow[0], "?")
			fmt.Scanln(&y)
			if y == "yes" {
				sum[0] = sum[0] + i2
				buyerPrice[0] = priceNow[0]
			}
		} else {
			fmt.Print("I am sorry,", name, ". You don’t have enough money.")
		}

	case "silver":
		if balance > i2*priceNow[1] {
			fmt.Print("Buying ", i2, " ", c, " bought at ", i2*priceNow[1], "?")
			fmt.Scanln(&y)
			if y == "yes" {
				sum[1] = sum[1] + i2
				buyerPrice[1] = priceNow[1]
				balance = balance - i2*priceNow[1]
			}
		} else {
			fmt.Print("I am sorry,", name, ". You don’t have enough money.")
		}

	case "gold":
		if balance > i2*priceNow[2] {
			fmt.Print("Buying ", i2, " ", c, " bought at ", i2*priceNow[2], "?")
			fmt.Scanln(&y)
			if y == "yes" {
				sum[2] = sum[2] + i2
				buyerPrice[2] = priceNow[2]
				balance = balance - i2*priceNow[2]
			}
		} else {
			fmt.Print("I am sorry,", name, ". You don’t have enough money.")
		}

	case "diamond":
		if balance > i2*priceNow[3] {
			fmt.Print("Buying ", i2, " ", c, " bought at ", i2*priceNow[3], "?")
			fmt.Scanln(&y)
			if y == "yes" {
				sum[3] = sum[3] + i2
				buyerPrice[3] = priceNow[3]
				balance = balance - i2*priceNow[3]
			}
		} else {
			fmt.Print("I am sorry,", name, ". You don’t have enough money.")
		}

	}
}

/* I.S. : this procedure it for the sell prompt
   F.S. : will save the item that we buy to the inventory/truck
*/
func sell(i2 int, c string) {
	switch c {
	case "silk":
		sum[0] = sum[0] - i2
		buyerPrice[0] = priceNow[0]
		balance = balance + (i2 * buyerPrice[0])
		fmt.Println("You sell ", i2, c, " for ", i2*buyerPrice[0])
	case "silver":
		sum[1] = sum[1] - i2
		buyerPrice[0] = priceNow[1]
		balance = balance + (i2 * buyerPrice[1])
		fmt.Println("You sell ", i2, c, " for ", i2*buyerPrice[1])
	case "gold":
		sum[2] = sum[2] - i2
		buyerPrice[2] = priceNow[2]
		balance = balance + (i2 * buyerPrice[2])
		fmt.Println("You sell ", i2, c, " for ", i2*buyerPrice[2])
	case "diamond":
		sum[3] = sum[3] - i2
		buyerPrice[3] = priceNow[3]
		balance = balance + (i2 * buyerPrice[3])
		fmt.Println("You sell ", i2, c, " for ", i2*buyerPrice[3])
	}
}

/* I.S. : this procedure it for check the account balance
   F.S. : the function will be output the balance after that
*/
func checkAccount() {
	maintence := (sum[0] + sum[1] + sum[2] + sum[3]) * 1 / 100
	balance = balance - maintence
	fmt.Println("Your balance is", balance)
	fmt.Println(" You have in your storage:")
	for i := 0; i < 4; i++ {
		fmt.Println(sum[i], " ", material[i], " at", buyerPrice[i])
	}
}

/* I.S. intialize the prrice into the array
   F.S. : will show the pricelist if there is "catalog"
*/
func checkcatalogA() {

	fmt.Println("Current commodity prices:")
	for i := 0; i < 4; i++ {
		fmt.Println(material[i], " ", priceNow[i])
	}
}

/* I.S. intialize the prrice into the array
   F.S. : will show the pricelist if there is "catalog"
*/

func checkcatalogB() {
	fmt.Println("Current commodity prices:")
	for i := 0; i < 4; i++ {
		fmt.Println(material[i], " ", priceNow[i])
	}
}

/* I.S. : inzizaltize and calculate the raising price with random
   F.S. : will show the output of the price randomly
*/

func showPriceA() {
	randchange := rand.Intn(100)
	if randchange < 50 {
		randup := rand.Intn(50)
		//randup = chg
		if randup < 35 {
			for i := 0; i < 4; i++ {
				priceNow[i] = priceNow[i] + (priceNow[i] * 25 / 100)
			}
		} else {
			for i := 0; i < 4; i++ {
				priceNow[i] = priceNow[i] - (priceNow[i] * 15 / 100)
			}

		}
	}
}

/* I.S. : initialize and calculate the raising price with random
   F.S. : will show the output of the price randomly
*/

func showPriceB() {
	randchange := rand.Intn(100)
	if randchange < 50 {
		randup := rand.Intn(50)
		//randup = chg
		if randup < 35 {
			for i := 0; i < 4; i++ {
				priceNow[i] = priceNow[i] + (priceNow[i] * 25 / 100)
			}
		} else {
			for i := 0; i < 4; i++ {
				priceNow[i] = priceNow[i] - (priceNow[i] * 15 / 100)
			}

		}
	}
}

/* I.S. : will sorting the balance
   F.S. : will show the sorting balance
*/
func sorting(logmoney *[1000]float64, longday *[1000]int) {
	var pos, i, max, daytempo int = 0, 0, 0, 0
	var tempmoney float64
	for pos <= 1000-1 {
		max = pos
		i = max + 1
		for i < 1000 {
			if logmoney[i] < logmoney[max] {
				max = i
			}
			i = i + 1
		}
		tempmoney = logmoney[pos]
		daytempo = longday[pos]

		logmoney[pos] = logmoney[max]
		longday[pos] = longday[max]

		logmoney[max] = tempmoney
		longday[max] = daytempo

		pos++
	}
}

/* I.S. : to search  using binary
   F.S. : show if it found or not
*/
func binaryFind(logmoney [1000]float64, longday [1000]int, x float64) {
	right := 1000 - 1
	left := 0
	pas := -1
	for left <= right && pas == -1 {
		mid := (left + right) / 2
		if logmoney[mid] == x {
			pas = mid
		} else if logmoney[mid] > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if pas == -1 {
		fmt.Println("NOT FOUND")
	} else {
		fmt.Println("FOUND", logmoney[pas], " in day ", longday[pas])
	}

}

/*
I.S all day and log money assign to array
f.S will update the array
*/
func log(longday *[1000]int, logmoney *[1000]float64, day int, balance float64, i int) {
	longday[i] = day
	logmoney[i] = balance
}

/*
I.S the two array contain day and money data inside the array
f.S will update the data according to the data
*/
func checkdata(longday [1000]int, logmoney [1000]float64) {
	i := 989
	for i < 1000 {
		fmt.Printf("%.2f", logmoney[i])
		fmt.Print(" ", longday[i])
		fmt.Println()
		i++
	}
}

/*
I.S this array it to initializw the balance into the maximum and minimum account baalnce
f.S will output the data accordung to the condition
*/

func Maxmin() {
	var max, min int
	if balance >= 500 {
		max = balance
		fmt.Println("Max account value:", max)
	} else if balance < 500 {
		min = balance
		fmt.Println("Minimum account value:", min)
	}
}

/* I.S. : this procedure it is for the "rent" prompt
   F.S. : will do the rent
*/
func rentWarehouse() {
	var ans string
	fmt.Print("Rent a warehouse? ")
	fmt.Scanln(&ans)
	if ans == "yes" {
		rent = true
	}
}

/* I.S. : intialize the material into the warehouse
   F.S. : and will output that the material save into the truck
*/
func store() {
	var amtI string
	var itemName string
	fmt.Print("How many item and what item? ")
	fmt.Scanln(&amtI, &itemName)
	switch itemName {
	case "silk":
		warehouse[0] += amtI
		material[0] = amtI
	case "silver":
		warehouse[1] += amtI
		material[1] = amtI

	case "gold":
		warehouse[2] += amtI
		material[2] = amtI

	case "diamond":
		warehouse[3] += amtI
		material[3] = amtI
	}
}

/* I.S. it is intialize the data  "unload"
F.S. : Upon return, will unload the warehoue to truck
*/
func unload() {
	var res, itemName string
	var amtI string
	fmt.Print("Unload what item and amount? ")
	fmt.Scanln(&amtI, &itemName)
	fmt.Println("Are you sure want to unload?")
	fmt.Scanln(&res)
	if res == "yes" {
		switch itemName {
		case "silk":
			material[0] += amtI
			warehouse[0] = amtI
		case "silver":
			material[1] += amtI
			warehouse[1] = amtI

		case "gold":
			material[2] += amtI
			warehouse[2] = amtI

		case "diamond":
			material[3] += amtI
			warehouse[3] = amtI
		}
	}
}

/* I.S. : To call the function and the procedure based on the prompt
   F.S. : will ouput the calling function or procedure
*/
func main() {
	var s1op, s2op, name, com string
	var iop, i int
	var search float64
	balance = 1000
	initialaccount = 1000
	var item int
	longday[0] = 1
	logmoney[0] = float64(balance)
	fmt.Print("Welcome sir/mam, your name is?")
	fmt.Scanln(&name)
	fmt.Println(name, "as a startup, you’ve been given account", balance, "Trade wisely.")
	prompt(&s1op, &iop, &s2op)
	for s1op != "quit" {
		showPriceA()
		showPriceB()
		if s1op == "buy" {
			buy(iop, s2op)
			log(&longday, &logmoney, day, float64(balance), j)
			sorting(&logmoney, &longday)
			checkdata(longday, logmoney)
		}
		if s1op == "sell" {
			sell(iop, s2op)
			log(&longday, &logmoney, day, float64(balance), j)
			sorting(&logmoney, &longday)
			checkdata(longday, logmoney)
		}
		if s1op == "show" && s2op == "catalog" {
			fmt.Print("Visit A or Visit B? ")
			fmt.Scanln(&com)
			if com == "A" {
				checkcatalogA()
			} else if com == "B" {
				checkcatalogB()
			}

		} else if s1op == "show" && s2op == "account" {
			checkAccount()

		}
		if s1op == "Rent" {
			balance = balance - warprice
			rentWarehouse()
		}
		if s1op == "Store" {
			store()
		}
		if s1op == "Unload" {
			unload()
		}
		if s1 != "show" {
			warprice = warprice + 1
			day = day + rand.Intn(5)
		}
		i = i + 1
		prompt(&s1op, &iop, &s2op)
	}
	item = 0
	i = 0
	for i <= 3 {
		item = item + (sum[i] * buyerPrice[i])
		i = i + 1
	}
	fmt.Println(name, "we liquidate all your asset.")
	fmt.Print("Search?")
	fmt.Scan(&search)
	binaryFind(logmoney, longday, search)
	Maxmin()
	nowaccount := float64(balance) + float64(item)
	diffacc := nowaccount - float64(initialaccount)
	fmt.Println("Your final asset value is ", balance)
	fmt.Println("Difference final:", diffacc)
	if diffacc >= 500 {
		fmt.Println("Apparently you win.")
	} else {
		fmt.Println("Apparently you Lost.")
	}
}
