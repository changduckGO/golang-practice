package main

import (
	"errors"
	"fmt"
	"time"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	channel := make(chan bool)
	people := [2]string{"harry", "SCD"}

	for _, person := range people {
		go sexyCount(person, channel)
	}

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

func sexyCount(person string, channel chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	channel <- true
}
