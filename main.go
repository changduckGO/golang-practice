// warning: main.go가 아니라면 컴파일 할 수 없음
// 컴파일하지 않고 다른 사람과의 공유 등의 목적이라면 파일 이름 상관없음
package main

import "fmt"

func main() {
	const e_name string = "HARRY"
	fmt.Println(e_name)

	k_name := "CD" // var k_name string = "CD"
	k_name = "SCD"
	fmt.Println(k_name)
}
