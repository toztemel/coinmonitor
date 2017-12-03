package main

import (
	"github.com/toztemel/coinmonitor/market"
	"github.com/toztemel/coinmonitor/cache"
	"time"
	"fmt"
)

func main() {

	fmt.Println(time.Now())
	result := market.Latest()

	cache.Init(result)

	ticker := time.NewTicker(time.Minute * 5)
	go func() {
		for t := range ticker.C {
			fmt.Println()
			fmt.Println("Tick at",t)
			fmt.Println()
			result = market.Latest()

			cache.Update(result)

		}
	}()

	time.Sleep(time.Minute * 30)
	ticker.Stop()
	fmt.Println("Ticker stopped")

}
