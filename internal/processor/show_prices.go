package processor

import "fmt"

func ShowPriceAVG(priceChannel <-chan float64, done chan<- bool) {
	var totalPrice float64
	countPrices := 0.0
	for price := range priceChannel {
		totalPrice += price
		countPrices++
		avgPrice := totalPrice / countPrices
		fmt.Printf("Preco recebido: R$ %.2f | Preco medio até agora: R$ %.2f \n", price, avgPrice)
	}
	done <- true
}
