package main

import "fmt"

// shadowing
func main() {
	var test1 float64 = 9.1  //test1을 float64방식으로 9.1을 저장하겠다
	var test2 float64 = 7.9  //test2을 float방식으로 7.9를 저장하겠다
	var univ string = "inha" //univdmf 문자열 방식으로 inha저장

	var f1 string = "functions"
	var f2 = append([]string{}, "함수") //배열이다.append로 strigs라는 배열을 f2에 담을거다 그 f2에 함수라는 문자를 넣겠다

	fmt.Println(test1 + test2) //9.1 + 7.9 =17
	fmt.Println(univ)          //inha
	fmt.Println(f1)            //functions
	fmt.Println(f2)            //[함수]

	// 자료타입을 변수명으로 사용
	// var float64 float64 = 9.1
	// var test float64 = 7.9
	// fmt.Println(float64)

	// 패키지를 변수명으로 사용
	// var fmt string = "inha"
	// fmt.Println()

	// 함수를 변수명으로 사용
	// var append string = "functions"
	// var f = append([]string{}, "함수")f의 strings타입의 배열을 추가하겠다
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// // 입력(0과 1처리)된 수의 소수 판정 프로그램 v0.9
// func main() {
// 	var number int

// 	fmt.Print("정수 입력 : ")
// 	_, err := fmt.Scanln(&number) 변수를 저장하지 않고 scanln으로 바로 떄려박는다.

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for number < 2 {
// 		fmt.Println(number, "는(은) 소수가 아닙니다~")
// 		os.Exit(0)
// 	}

// 	isPrime := true
// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			break
// 		}
// 	}

// 	if isPrime {
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다~")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"log"
// )

// // 입력(fmt 패키지의 Scanln())된 수의 소수 판정 프로그램 v0.8
// func main() {
// 	var number int
// 	fmt.Print("정수 입력 : ")
// 	_, err := fmt.Scanln(&number)

// 	//fmt.Println(n)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	isPrime := true
// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			break
// 		}
// 	}

// 	if isPrime {
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다~")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	//"math/rand"
// 	//"time"
// 	"bufio"
// 	"log"
// 	"os"
// )

// // 입력된 수의 소수 판정 프로그램 v0.7
// func main() {
// 	fmt.Print("정수 입력 : ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, err := rd.ReadString('\n')

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	in = strings.TrimSpace(in)
// 	number, err := strconv.Atoi(in)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	isPrime := true
// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			break
// 		}
// 	}

// 	if isPrime {
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다~")
// 	}
// }
