package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("https://ifconfig.es/json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	df1, err := os.OpenFile("destinationA.txt", os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal(err)
	}
	df2, err := os.OpenFile("destinationB.txt", os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal(err)
	}

	// We are going to calculate the md5Sum without knowing the entire data
	sum := md5.New()

	// Prepare writers fan out!
	mw := io.MultiWriter(df1, df2, os.Stdout, sum)

	// Send the body to all writers !
	_, err = io.Copy(mw, resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	// Closing resources
	if err := df1.Close(); err != nil {
		log.Fatal(err)
	}
	if err := df2.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", sum.Sum(nil))
}
