package main

import "fmt"

func main() {
	// var s []int
	// s = make([]int, 5)

	//s := make([]int,5)//make함수를 이용해 슬라이스 생성 후 메모리 할당, 제로 값 적용

	s := []int{0, 0, 0, 0, 0} //슬라이스 리터럴을 이용해 슬라이스 생성 및 메모리 항당, 초기화 진행

	s[4] = 99
	s[2] = 91

	for _, value := range s {
		fmt.Println(value)
	}
	copys := s[1:4]
	for _, value := range copys {
		fmt.Println(value)
	}
	test := [3]string{"inha", "go", "student"} //배열 리터럴을 이용해서 test배열 생성
	tests := test[0:2]
	fmt.Println(len(tests))
}
