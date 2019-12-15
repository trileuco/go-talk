package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type IfconfigResponse struct {
	IP          string `json:"IP"`
	Country     string `json:"Country"`
	CountryCode string `json:"Country Code"`
	City        string `json:"City"`
	Latitude    string `json:"Latitude"`
	Longitude   string `json:"Longitude"`
	Accept      string `json:"Accept"`
	Host        string `json:"Host"`
	UserAgent   string `json:"User-Agent"`
}

func main() {
	resp, err := http.Get("https://ifconfig.es/json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	icr := IfconfigResponse{}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(data, &icr); err != nil {
		log.Fatal(err)
	}
	fmt.Println(icr)
}
