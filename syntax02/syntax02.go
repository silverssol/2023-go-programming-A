package main

import (
	"fmt"
	//"reflect"
	"strings"
)

func main() {
	
	texts := "G@ M@ney~"
	fmt.Println(texts)
	r := strings.NewReplacer("@" ,"o")
	newTexts := r.Replace(texts)
	fmt.Println(newTexts)

// // 변수명은 영문자로 시작해야된다.
	// // 영문 대문자의 경우 다른 패키지에서 접근할 수 있다. 소문자로 시작하는 변수는 동일 패키지에서만 접근 가능
// 	var e string
// 	var d bool
// 	var b float64
// 	var c rune
// 	var a int
// 	//a := 7
// //	zero value
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Println(e)

// 	fmt.Printf("%T\n", d)
// 	fmt.Printf("%T\n", e)

// 	fmt.Println(reflect.TypeOf(d))
// 	fmt.Println(reflect.TypeOf(e))
}
