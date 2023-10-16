package main

import (
	"fmt"  //출력함수를 포함하는 패키지
	"math" //수학함수
	"strings"
)

func main() {

	var c rune = '가'
	//var a int16 = 7
	//var a = 7
	//a := 7
	a := 7
	// var b float64 = 5.3
	b := 5.3
	a = int(b) //Type Conversion, Casting
	d := false

	fmt.Println(d)
	fmt.Printf("%T\n", d)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(c)        //유니코드utf-8 값출력
	fmt.Printf("%c\n", c) //한 글자 출력
	fmt.Printf("%T\n", c) //rune은 결국 int32의 별명

	fmt.Println(math.Round(2.71))
	fmt.Println("hello Go~~")
	fmt.Println(strings.Title("go git github java"))
}
