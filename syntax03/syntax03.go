package main

import (
	"bufio"
	"fmt"
	"log" //Fatal function 합수  err메세지 출력 후 프로그램 종료
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("단 입력: ")
	rd := bufio.NewReader(os.Stdin)
	in, err := rd.ReadString('\n')

	if err != nil { //에러가 발생하면
		log.Fatal(err)
	}

	in = strings.TrimSpace(in)
	dan, err := strconv.ParseInt(in, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i < 10; i++ {
		fmt.Println(dan, " * ", i, " = ", (dan * i))
	}
	fmt.Println(dan * 2)
}
