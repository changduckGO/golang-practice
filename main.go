// warning: main.go가 아니라면 컴파일 할 수 없음
// 컴파일하지 않고 다른 사람과의 공유 등의 목적이라면 파일 이름 상관없음
package main

import (
	"fmt"
)

func superADD(numbers ...int) int {
	total := 0
	for _, number := range numbers { // go에서 range는 index와 element를 반환해줌(python에서 enumerate와 같은 역할)
		total += number
	}

	return total
}

func main() {
	total_sum := superADD(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(total_sum)
}
