package main

import (
	"flag"
	"fmt"
)

var beers = map[string]string{
	"x": "4",
	"f": "3",
	"a": "1",
	"g": "2",
}

func main() {
	b := flag.Bool("beers", false, "print beers")
	flag.Parse()
	if *b {
		fmt.Println(beers)
	}
	fmt.Printf("hello")
}
