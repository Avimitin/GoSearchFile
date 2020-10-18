package main

import (
	"fmt"
	"time"
	"os"
	"github.com/Avimitin/SearchFile/internal/search"
)

func main() {
	if length := len(os.Args); length != 3 {
		fmt.Printf("search need 2 args but got only %d\n", length-1)
		os.Exit(-1)
	}

	target := os.Args[1]
	orderPath := os.Args[2]
	fmt.Printf("target: %s\norderpath:%s\nDo you want to continue? [y/N] ", target, orderPath)
	var confirm string
	fmt.Scanf("%s", confirm)
	if confirm == "n" || confirm == "N" {
		os.Exit(0)
	}

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
