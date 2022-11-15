// warning: main.go가 아니라면 컴파일 할 수 없음
// 컴파일하지 않고 다른 사람과의 공유 등의 목적이라면 파일 이름 상관없음
package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	totalLength, upperName := lenAndUpper("changduck")
	fmt.Println(totalLength, upperName)

	repeatMe("cd", "harry", "SCD", "duck")
}
