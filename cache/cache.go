package cache

import (
	"fmt"
	"github.com/toztemel/cryptoping"
	"math"
)

var cache map[string]cryptoping.Record

const log_precision float64 = 0.1
const notify_precision float64 = 0.5

const log_precision_rank int = 0
const notify_precision_rank int = 2

func Init(records []cryptoping.Record) {

	if records == nil {
		panic("No results")
	}

	fmt.Println("Retrieve initial results:")
	cache = make(map[string]cryptoping.Record)

	for _, record := range records {
		add(record)
	}

}

func add(record cryptoping.Record) {

	cache[record.Symbol] = record
	if (record.MarketCapUsd >= 1000000000) {
		fmt.Printf("\t%s\t%.2f\n", record.Symbol, record.PriceUSD)
	}

}

func Update(records []cryptoping.Record) {
	for _, ticker := range records {
		update(ticker)
	}
}

func update(record cryptoping.Record) {

	if record.LastUpdated == cache[record.Symbol].LastUpdated {
		return
	}

	symbol := record.Symbol

	var change float64

	change = float64((record.PriceUSD - cache[symbol].PriceUSD) / cache[symbol].PriceUSD)
	if math.Abs(change) >= log_precision {
		fmt.Printf("\t%s price changed:%.2f\n", record.Name, change)
		cache[symbol] = record
		if (math.Abs(change) >= notify_precision) {
			// notify
		}
	}

	change = float64((record.MarketCapUsd - cache[symbol].MarketCapUsd) / record.MarketCapUsd)
	if math.Abs(change) >= log_precision {
		fmt.Printf("\t%s market cap changed:%.2f\n", record.Name, change)
		cache[symbol] = record
		if (math.Abs(change) >= notify_precision) {
			// notify
		}
	}

	ranking := (record.Rank - cache[symbol].Rank)
	if math.Abs(float64(ranking)) > float64(log_precision_rank) {
		fmt.Printf("\t%s rank changed:%d\n", record.Name, ranking)
		cache[symbol] = record

		if math.Abs(float64(ranking)) > float64(notify_precision_rank) {
			// notify ranking
		}
	}

}
