package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

const API_COINMARKETCAP_COM = "https://api.coinmarketcap.com/v1/ticker/?limit=3"

type Ticker struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Rank         string `json:"rank"`
	PriceUSD     string `json:"price_usd"`
	PriceBTC     string `json:"price_btc"`
	VolumeUSD    string `json:"34h_volume_usd"`
	MarketCapUsd string `json:"market_cap_usd"`
	Supply       string `json:"available_supply"`
	Change1h     string `json:"percent_change_1h"`
	Change24h    string `json:"percent_change_24h"`
	Change7d     string `json:"percent_change_7d"`
	LastUpdated  string `json:"last_updated"`
}

func main() {

	result := latest()
	fmt.Println(result[0].Id)

}

func latest() ([]Ticker) {

	resp, err := http.Get(API_COINMARKETCAP_COM)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	prettyBody, err := pretty(body)
	if err != nil {
		panic(err.Error())
	}

	result, err := getTickers([]byte(prettyBody))
	if err != nil {
		panic(err.Error())
	}

	return result
}

func pretty(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "   ")
	return out.Bytes(), err
}

func getTickers(body []byte) ([]Ticker, error) {
	var s []Ticker
	err := json.Unmarshal(body, &s)
	if (err != nil) {
		fmt.Println("errror:", err)
	}
	return s, err
}
