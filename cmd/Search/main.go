package main

import (
	"fmt"
	"time"

	"github.com/Avimitin/SearchFile/internal/search"
)

func main() {
	target := "test"
	orderPath := "/home/avimitin"

	fmt.Println("===Using one thread===")
	start := time.Now()
	search.Search(target, orderPath)
	fmt.Println(search.Matches, "matches")
	fmt.Println("During", time.Since(start))

	search.Matches = 0
	search.Target = target
	search.MaxThread = 32	

	fmt.Println("===Using multi thread===")
	start = time.Now()
	go search.CompSearch(orderPath, true)
	search.WaitForResult()
	fmt.Println(search.Matches, "matches")
	fmt.Println("During", time.Since(start))
}
