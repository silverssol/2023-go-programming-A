package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	//var a int16 = 7
	//var a = 7
	//a := 7
	a := 7 //var a int =7
	var b float64 = 5.3
	//b := 5.3
	var c rune = '가' //rune은 ''로 표시
	//c := '가'
	a = int(b) // Type Conversion, Casting
	d := false

	fmt.Println(d)
	fmt.Printf("%T\n", d) //d의 타입은 어떻게 되는지

	fmt.Println(a)
	fmt.Printf("%T\n", a) // a의 타입은 어떻게 되는지

	fmt.Println(c)        // 유니코드(UTF-8) 값 출력
	fmt.Printf("%c\n", c) // 한 글자 출력
	fmt.Printf("%T\n", c) // rune은 결국 int32의 별명

	fmt.Println(math.Round(2.71)) //3으로 출력 왜냐하면 math.Round 함수는 올림해주는 함수라서
	fmt.Println("Hello Go~")
	fmt.Println(strings.Title("go git github java")) //단어 앞부분 대문자로 변경
}
