package mapMissingItem

import "fmt"

func Run() {
	mapMissingItem()
}

func mapMissingItem() {
	prices := map[string]int{
		"Apple": 0, // Apples are free!
		// We don't have Bananas
	}

	fmt.Printf("Our menu: %v\n\n", prices)

	fmt.Printf("The price of Apple is $%d\n", prices["Apple"])
	fmt.Printf("The price of Banana is $%d\n", prices["Banana"])

	fmt.Println("")
	fmt.Println("Checking with 'CommaOK':")
	priceA, okA := prices["Apple"]
	priceB, okB := prices["Banana"]
	fmt.Printf("The price of Apple is $%d | ok: %v\n", priceA, okA)
	fmt.Printf("The price of Banana is $%d | ok: %v\n", priceB, okB)
	fmt.Println("")
}
