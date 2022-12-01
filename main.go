package main

import (
	"errors"
	"fmt"
	"time"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	go sexyCount("harry")
	go sexyCount("SCD")

	time.Sleep(time.Second * 5)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
