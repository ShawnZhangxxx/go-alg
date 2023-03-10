package main

import (
	"fmt"
	"time"
)

func main() {
	var m = make(map[int]int)

	// init memory structure of map
	for i := 0; i < 1000; i++ {
		m[i] = 0
	}

	// just update values for old keys concurrently
	for i := 0; i < 1000; i++ {
		go func() {
			m[i] = i
		}()
	}

	time.Sleep(time.Second * 5)
	fmt.Println(m)

}
