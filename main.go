// warning: main.go가 아니라면 컴파일 할 수 없음
// 컴파일하지 않고 다른 사람과의 공유 등의 목적이라면 파일 이름 상관없음
package main

import (
	"fmt"
)

func main() {
	harry := map[string]string{"name": "SCD", "age": "18"} // map의 key와 value 값을 string으로 지정

	for key, value := range harry {
		fmt.Println(key, value)
	}
}
