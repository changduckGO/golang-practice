package main

import (
	"errors"
	"fmt"
	"time"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	channel := make(chan string)
	people := [5]string{"harry", "SCD", "CHAN", "DUCK", "HARRY POTTER"}

	for _, person := range people {
		go sexyCount(person, channel)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(<-channel)
	}
}

func sexyCount(person string, channel chan string) {
	time.Sleep(time.Second * 5)
	channel <- person + " is sexy"
}
