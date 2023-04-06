package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	d, err := time.ParseDuration(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d", d/time.Second)
}
