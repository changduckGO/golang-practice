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
		fmt.Println(<-channel) // <-channel은 일종의 blocking operation
		// 다시 말해서 main 함수에서 무언가를 받기 전까지는 동작을 멈춘다는 의미
	}
}

func sexyCount(person string, channel chan string) {
	time.Sleep(time.Second * 5)
	channel <- person + " is sexy"
}
