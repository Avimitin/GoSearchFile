package search

import "io/ioutil"

// Matches is number
var Matches int
var Target string
var MaxThread int
var done = make(chan bool)
var founded = make(chan bool)
var request = make(chan string)
var currentThread int = 1

// Search is using for search
func Search(query string, path string) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			fileName := file.Name()
			if fileName == query {
				Matches++
			}

			if file.IsDir() {
				Search(query, path+fileName+"/")
			}
		}
	}
}

// CompSearch use gorouting to search.
func CompSearch(path string, master bool) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			fileName := file.Name()

			if fileName == Target {
				founded <- true
			}

			if file.IsDir() {
				newPath := path + fileName + "/"
				if currentThread < MaxThread {
					request <- newPath
				} else {
					CompSearch(newPath, false)
				}
			}
		}

		if master {
			done <- true
		}
	}
}

func WaitForResult() {
	for {
		select {
		case path := <-request:
			currentThread++
			go CompSearch(path, true)
		case <- done:
			currentThread--
			if currentThread == 0 {
				return
			}
		case <- founded:
			Matches++
		}
	}
}

// Example method for testing
func Example() int {
	return 1
}
