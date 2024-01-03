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
	checkPriceWithCommaOk(prices, "Apple")
	checkPriceWithCommaOk(prices, "Banana")
	fmt.Println("")
}

func checkPriceWithCommaOk(prices map[string]int, product string) {
	price, ok := prices[product]
	fmt.Printf("The price of %s is $%d | ok flag: %v\n", product, price, ok)
}
