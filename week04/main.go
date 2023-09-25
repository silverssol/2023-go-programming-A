package main

import (
	"fmt"
	"reflect"
)

func main() {

	// var a int
	// a = 9
	// fmt.Println(a)
	// var a = 9
	var b float32 = 2.71
	a := 9
	// bm := 2.71
	c := 'Z'
	d, e := 4.10, 3.99

	var g int
	var h float32
	var i bool
	var j string
	fmt.Println(j, g, h, i)
	fmt.Printf("%s %d %f\n", j, g, h)
	// k := "변수명이 대문자로 시작하면 다른 패키지에서 이 변수를 사용할 수 있음"
	fmt.Println(a, reflect.TypeOf(a), b, reflect.TypeOf(b), c, reflect.TypeOf(c), d, reflect.TypeOf(e), e, reflect.TypeOf(e))
	fmt.Println(float32(a) * b)
	fmt.Println(a * int(b))
	// fmt.Println(a > b)

	// fmt.Println(a)
	// fmt.Println(reflect.TypeOf('z'))
	// fmt.Println('A', 'a', '0')
	// fmt.Println(math.Ceil(3.17))
	// fmt.Println(math.Floor(3.17))
	// fmt.Println(strings.Title("오픈소스 프로그래밍"))
	// fmt.Println(strings.Title("OpenSource programming~"))
}
