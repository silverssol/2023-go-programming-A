package main

import (
	"fmt"
)

func main() {
	//var balls map[string]int
	balls := make(map[string]int)
	fmt.Printf("%#v\n", balls)
	balls["성기훈"] = 20
	fmt.Println("성기훈")
	fmt.Println("오일남")
}
