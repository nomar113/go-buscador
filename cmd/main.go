package main

import (
	"buscador/internal/fetcher"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	priceChannel := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		var totalPrice float64
		countPrices := 0.0
		for price := range priceChannel {
			totalPrice += price
			countPrices++
			avgPrice := totalPrice / countPrices
			fmt.Printf("Preco recebido: R$ %.2f | Preco medio até agora: R$ %.2f \n", price, avgPrice)
		}
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromSite1()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromSite2()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromSite3()
	}()

	wg.Wait()
	close(priceChannel)

	fmt.Printf("Tempo total: %s", time.Since(start))
}
