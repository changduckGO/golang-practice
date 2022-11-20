// warning: main.go가 아니라면 컴파일 할 수 없음
// 컴파일하지 않고 다른 사람과의 공유 등의 목적이라면 파일 이름 상관없음
package main

import (
	"fmt"
)

func main() {
	names := []string{"HARRY", "SCD", "CHANGDUCK"} // length를 빈칸으로 해줌으로써 Slice 객체 생성
	names = append(names, "HARRY4DUCK")            // append arguments: Slice, 추가할 요소
	fmt.Println(names)
}
