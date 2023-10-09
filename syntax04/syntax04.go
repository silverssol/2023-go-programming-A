package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 난수 추출된 수의 소수 판정 프로그램
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	number := rand.Intn(10) + 2 // 0과 1제외, 2 ~ 151 사이의 수
	fmt.Println("임의로 추출된 수 : ", number)
}

// // import (
// // 	"fmt"
// // 	"math/rand"
// // 	"time"
// // )

// // func main(){

// // 	seed := time.Now().Unix()
// // 	rand.Seed(seed)

// // 	for i := 1; i<6; i++{
// // 		dice := rand.Intn(6) + 1
// // 		fmt.Println(dice)
// // 	}
// // 	number := rand.Intn(3)
// // 	fmt.Println(number)
// // }
