// warning: main.go가 아니라면 컴파일 할 수 없음
// 컴파일하지 않고 다른 사람과의 공유 등의 목적이라면 파일 이름 상관없음
package main

import (
	"fmt"
)

func main() {
	// &: 메모리 주소를 확인
	// *: 해당 주소에 있는 값을 살펴보기
	a := 97
	b := &a
	*b = 200
	fmt.Println(a)
}
