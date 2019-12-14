package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getIp() (string, error) {

	resp, err := http.Get("https://www.myexternalip.com/raw")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close() // very important for releasing OS sockets

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {

	ip, err := getIp()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ip)
}
