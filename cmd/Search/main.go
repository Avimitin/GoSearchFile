package main

import (
	"fmt"
	"time"

	"github.com/Avimitin/SearchFile/internal/search"
)

func main() {
	fmt.Println("===Using one thread===")
	start := time.Now()
	search.Search("test", "/home/avimitin/")
	fmt.Println(search.Matches, "matches")
	fmt.Println("During", time.Since(start))

	search.Matches = 0
	search.Target = "test"
	search.MaxThread = 32	

	fmt.Println("===Using multi thread===")
	start = time.Now()
	go search.CompSearch("/home/avimitin/", true)
	search.WaitForResult()
	fmt.Println(search.Matches, "matches")
	fmt.Println("During", time.Since(start))
}
