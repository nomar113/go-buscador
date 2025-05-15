package main

import (
	"buscador/internal/fetcher"
	"buscador/internal/processor"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	priceChannel := make(chan float64)
	done := make(chan bool)

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAVG(priceChannel, done)

	<-done

	fmt.Printf("Tempo total: %s", time.Since(start))
}
