package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.Version())
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	runtime.GC()
}
