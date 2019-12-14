package main

import (
	"fmt"
)

func job() {
	defer println("cleanUp done 2")
	defer println("cleanUp done 1")
	fmt.Println("thing 1 ..")
	fmt.Println("thing 2 ..")
}

func main() {
	job()
}
