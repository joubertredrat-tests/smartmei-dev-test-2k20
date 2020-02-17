package main

import "encoding/json"
import "io/ioutil"
import "net/http"

func getExchange() (float64, float64) {
	resp, err := http.Get("https://api.exchangeratesapi.io/latest?base=BRL&symbols=USD,EUR")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	rates := result["rates"].(map[string]interface{})
	
	return rates["USD"].(float64), rates["EUR"].(float64)
}