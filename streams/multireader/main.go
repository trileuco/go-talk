package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	f1, err := os.Open("sourceA.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()
	f2, err := os.Open("sourceB.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()
	m := io.MultiReader(f1, f2)
	d, err := ioutil.ReadAll(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(d))
}
