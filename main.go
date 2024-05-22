package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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

func getPeoples(id int) person {
	url := fmt.Sprintf("https://localhost:7129/People/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return person{}
	}
	defer resp.Body.Close()
	var p person
	json.NewDecoder(resp.Body).Decode(&p)
	return p
}

type person struct {
	ID           int
	GivenName    string
	FamilyName   string
	StartDate    time.Time
	Rating       int
	FormatString string
}

func (p person) String() string {
	if p.FormatString != "" {
		return fmt.Sprintf("%s %s", p.FamilyName, p.GivenName)
	} else {
		return fmt.Sprintf("%s %s", p.GivenName, p.FamilyName)
	}
}

func main() {
	ids := getIDs()
	fmt.Println(ids)

	if len(ids) > 0 {
		for i := 0; i < len(ids); i++ {
			people := getPeoples(ids[i])
			fmt.Printf("%d: %v\n", people.ID, people)
		}

		for _, n := range ids {
			people := getPeoples(n)
			fmt.Printf("%d: %v\n", people.ID, people)
		}

	}

}
