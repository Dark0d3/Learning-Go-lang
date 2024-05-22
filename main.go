package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIDs() (ids []int) {
	resp, err := http.Get("https://localhost:7129/WeatherForecast/ids")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	json.Unmarshal(body, &ids)
	return
}

func main() {
	ids := getIDs()
	fmt.Println(ids)
}
