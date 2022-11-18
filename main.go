// warning: main.go가 아니라면 컴파일 할 수 없음
// 컴파일하지 않고 다른 사람과의 공유 등의 목적이라면 파일 이름 상관없음
package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 20 {
		fmt.Println("한국 나이: ", koreanAge)
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
