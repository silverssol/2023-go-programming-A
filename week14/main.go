package main

import (
	"fmt"
)

func main() {
	// var games map[int]string
	// games = make(map[int]string)
	//map liter
	games := map[int]string{

		456: "성기훈",
		218: "박해수",
		067: "강새벽",
		001: "오일남",
		199: "알리",
		101: "아이오아이",
	}
	//fmt.Println(games[100])
	//name, ok := games[100]
	name, ok := games[101]
	fmt.Println(name, ok)

	for _, v := range games {
		fmt.Println(v)
	}
	games[101] = "장덕수" //update
	delete(games, 199) // del

	for k, v := range games {
		fmt.Println(k, v)
	}
}