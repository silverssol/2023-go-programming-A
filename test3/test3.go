package main

import (
	"bufio"   //입출력을 도와주는 장치
	"fmt"     //입출력하게 해주는 패키지
	"log"     // Fatal function //에러발생시 중지시키는 패키지
	"os"      //사용자에게 입력을 받는 패키지
	"strconv" // TrimSpace  문자열을 다양한 데이터 형태로 바꿔주는 패키지
	"strings" // ParseInt //문자열을 정수로 바꿔주는 패키지로 Atoi는 10진수로 고정되어 있지만 ParseInt는 변수, 진수, 비트 값을 내가 지정할 수 있어
	//더 유연하다
)

func main() {
	fmt.Print("단 입력 : ")
	rd := bufio.NewReader(os.Stdin) //사용자의 입력을 받는 함수를 bufio.NewReader이라는 함수로 더 유연하게 읽어주고 rd의 저장
	in, err := rd.ReadString('\n')  //\n이 나오기 전 rd한 줄까지를 읽고 err발생시 if문으로 들어감

	if err != nil { // 에러 발생시 nil이 아니면
		log.Fatal(err) //err출력 후 종료
	}

	in = strings.TrimSpace(in)   //in의 저장된 문자열 양옆 공백을 없애는 함수를 써주고 in에다가 다시 저장
	dan, err := strconv.Atoi(in) //in문자열을 문자형을 정수로 바꿔주는 Atoi함수로 바꿔주고 dan에 저장
	if err != nil {              //err발생시 nil이 아니면
		log.Fatal(err) //오류 출력 후 종료
	}

	// 다른 언어의 while문 구현
	i := 1
	for i < 10 {
		fmt.Printf("%d * %d = %d\n", dan, i, (dan * i)) //Printf는 c의 scanf역할을 함
		i++
	}
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log" // Fatal function
// 	"os"
// 	"strconv" // TrimSpace
// 	"strings" // ParseInt
// )

// func main() {
// 	fmt.Print("단 입력 : ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, err := rd.ReadString('\n')

// 	if err != nil { // 에러가 발생하면
// 		log.Fatal(err)
// 	}

// 	in = strings.TrimSpace(in)
// 	//dan, err := strconv.ParseInt(in, 10, 32)
// 	dan, err := strconv.Atoi(in)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for i := 1; i < 10; i++ {
// 		//fmt.Println(dan, " * ", i, " = ", (int(dan) * i))
// 		//fmt.Println(dan, " * ", i, " = ", (dan * i))
// 		fmt.Printf("%d * %d = %d\n", dan, i, (dan * i))
// 	}
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Print("숫자 입력 : ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, _ := rd.ReadString('\n')
// 	fmt.Println(in)
// }
