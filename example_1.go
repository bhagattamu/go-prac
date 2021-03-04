// map is a datastructure where key points to the value
// Keys must be of the same type and values must be of the same type

package main

import "fmt"

func main() {
	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 2321.3,
		"WRCW": 2145.2,
	}

	// Number of items
	fmt.Println(len(stocks))

	// Get a value
	fmt.Println(stocks["AMZN"])

	// Get 0 value if not found
	fmt.Println(stocks["AFDW"])

	// Use two value to see if found
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println("TSLA found", value)
	}

	// Set
	stocks["TSLA"] = 234
	fmt.Println(stocks)

	// Delete
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	// Single value for in on keys
	for key := range stocks {
		fmt.Println(key)
	}

	// Double value "for" is key, value
	for key, value := range stocks {
		fmt.Printf("%s -> %.2f \n", key, value) // .2 is to 2 decimal value
	}
}
