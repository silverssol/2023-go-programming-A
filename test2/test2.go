package main

import (
	"fmt"
	"reflect" //Type of함수를 사용해서 입력한 것에 타입을 알아내는 패키지
	"strings"
)

func main() {
	texts := "G@ M@ney~~" //var texts str = "G@ M@ney~~"
	fmt.Println(texts)
	r := strings.NewReplacer("@", "o") //strings 패키지에서 NewReplaser라는 새로운 문자로 바꿔주는 함수를 가져와 바꿀 문자@를 바꿔줄 문자o로 바꿔줌
	newTexts := r.Replace(texts)       //위에 새로 정의 된 r을 가져와 tests의 새로 바꿔주고 newTexts의 담아주고
	fmt.Println(newTexts)              //출력하면 바뀐 문구를 저장한 newTexts를 출력하면 @가 아닌o로 바뀐 Go Money로 출력

	// 변수명은 영문자로 시작해야된다.
	// 영문 대문자의 경우 다른 패키지에서 접근할 수 있다. 소문자로 시작하는 변수는 동일 패키지에서만 접근 가능
	var e string
	var d bool
	var c rune
	var b float64
	var a int

	// naming convention
	var studentId string
	var i int32

	//a := 7
	fmt.Println(studentId)
	fmt.Println(i)

	// zero value
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	//fmt.Println(e)

	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(e))
}
