package main

import (
	"buscador/internal/fetcher"
	"buscador/internal/models"
	"buscador/internal/processor"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	priceChannel := make(chan models.PriceDetail, 4)
	done := make(chan bool)

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAVG(priceChannel, done)

	<-done

	fmt.Printf("Tempo total: %s", time.Since(start))
}
