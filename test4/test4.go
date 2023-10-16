package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// 난수 추출된 수의 소수 판정 프로그램 v0.6
func main() {
	seed := time.Now().Unix() //seed값에 math패키지를 써서 현재시간을 기준으로 seed값을 생성하겠다
	rand.Seed(seed)           //seed값을 랜덤으로 계속 바꿔주겠다

	isPrime := true //
	reflect.TypeOf(isPrime)
	number := rand.Intn(150) + 2 //정수형으로 만들건데 컴퓨터가 읽으면 0부터 시작이니 +2를 해서 1부터 시작하게 바꾸겠다
	//number = 21
	fmt.Println("임의로 추출된 수 : ", number) //number을 임의로 추출해서 150 숫자중 아무거나 랜덤으로 뽑아서 출력

	for i := 2; i < number; i++ {
		if number%i == 0 { //i로 number를 나눈 값이 0이면
			isPrime = false //false로 출력
			break           // 첫 번째 약수가 발견되면 반복문 즉시 종료
		}
		fmt.Print(i, " ") //0이면 false라서 출력이 안되는데 0이 아니면 0이 아닌 숫자들이 전부 출력
	}

	if isPrime { // 비교 연산자 제거
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다~")
	}
}

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // 난수 추출된 수의 소수 판정 프로그램 v0.4
// func main() {
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	isPrime := true
// 	number := rand.Intn(150) + 2
// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 		}
// 		fmt.Print(i, " ")
// 	}

// 	if isPrime { // 비교 연산자 제거
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다~")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // 난수 추출된 수의 소수 판정 프로그램 v0.3
// func main() {
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	isPrime := true
// 	number := rand.Intn(150) + 2
// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			//count = count + 1
// 		}
// 	}

// 	if isPrime == true {
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다~")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // 난수 추출된 수의 소수 판정 프로그램 v0.2
// func main() {
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	count := 0
// 	number := rand.Intn(150) + 2 // 0과 1제외, 2 ~ 151 사이의 수
// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 2; i < number; i++ { // 1과 number일때 loop 돌지 않음
// 		if number%i == 0 {
// 			//count++
// 			count = count + 1
// 		}
// 	}

// 	if count == 0 {
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다~")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // 난수 추출된 수의 소수 판정 프로그램 v0.1
// // 소수 : 1과 자기자신외에는 나누어 떨어지지 않는 수 (0과 1은 제외)
// func main() {
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	count := 0
// 	number := rand.Intn(150) + 2 // 0과 1제외, 2 ~ 151 사이의 수
// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 1; i <= number; i++ {
// 		if number%i == 0 {
// 			count++
// 		}
// 	}

// 	if count == 2 {
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다~")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time" // seed 생성용 패키지
// )

// func main() {
// 	// seed 설정
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	for i := 1; i < 6; i++ {
// 		dice := rand.Intn(6) + 1
// 		fmt.Println(dice)
// 	}
// }
