// Methods

package main

import (
	"fmt"
	"os"
)

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("Symbol cannot be empty")
	}
	if volume <= 0 {
		return nil, fmt.Errorf("Volume must be >= 0 (was %d)", volume)
	}
	if price <= 0.0 {
		return nil, fmt.Errorf("Price must be >= 0 (was %.2f)", price)
	}

	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}
	return trade, nil
}

// Value return the trade value
// receiver(t *Trade)
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

func main() {
	t, err := NewTrade("MSFT", 10, 99.98, true)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(t.Value())
}
