package main

import (
	"fmt"
	"math/rand"
)

func main() {
	distance := 62100000

	fmt.Println("Spaceline       Days  Trip  type  Price")
	fmt.Println("=======================================")

	for i := 10; i > 0; i-- {

		speed := rand.Intn(15) + 16
		days := distance / speed
		price := rand.Intn(15) + 36

		keys := rand.Intn(3)
		trips := rand.Intn(2)

		spacekeys := ""
		triptype := ""

		switch keys {
		case 0:
			spacekeys = "virgin Galactic"
		case 1:
			spacekeys = "Spacex"
		case 2:
			spacekeys = "Space Adventure"
		}

		switch trips {
		case 0:
			triptype = "Round-trip"
		case 1:
			triptype = "One-way"
		}

		fmt.Printf("%-15v  %-6v %10v  $ %2v \n", spacekeys, days, triptype, price)
	}

}
