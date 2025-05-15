package processor

import (
	"buscador/internal/models"
	"fmt"
)

func ShowPriceAVG(priceChannel <-chan models.PriceDetail, done chan<- bool) {
	var totalPrice float64
	countPrices := 0.0
	for priceDetail := range priceChannel {
		totalPrice += priceDetail.Value
		countPrices++
		avgPrice := totalPrice / countPrices

		fmt.Printf("[%s] Preco recebido de %s: R$ %.2f | Preco medio até agora: R$ %.2f \n", priceDetail.Timestamp.Format("02-01 15:04:05"), priceDetail.StoreName, priceDetail.Value, avgPrice)
	}
	done <- true
}
