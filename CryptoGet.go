package cryptoping

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const api = "https://api.coinmarketcap.com/v1/ticker/?limit=50"

func Latest() []Record {

	resp, err := http.Get(api)
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

	result, err := toRecords([]byte(prettyBody))
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

func toRecords(body []byte) ([]Record, error) {
	var s []Record
	err := json.Unmarshal(body, &s)
	return s, err
}
