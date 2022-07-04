package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/mdisprgm/goball/utils"
)

type Baller struct {
	a, b, c int
}

func (bl *Baller) Setup() {
	for bl.a = utils.Generate(); bl.a == 0; {
		bl.a = utils.Generate()
	}
	for bl.b = utils.Generate(); bl.a == bl.b; {
		bl.b = utils.Generate()
	}
	for bl.c = utils.Generate(); bl.a == bl.c || bl.b == bl.c; {
		bl.c = utils.Generate()
	}
}

func (bl *Baller) Check(s string) (int, int, error) {
	if len(s) != 3 {
		return -1, -1, errors.New("잘못된 입력입니다")
	}

	var strike, ball int

	if s[0] == s[1] || s[1] == s[2] || s[0] == s[2] {
		fmt.Println("중복된 숫자가 있습니다")
		return -1, -1, nil
	}

	user_number, _ := strconv.Atoi(s)

	ua := user_number / 100
	ub := (user_number / 10) % 10
	uc := user_number % 10

	if bl.a == ua {
		strike++
	} else if bl.a == ub || bl.a == uc {
		ball++
	}
	if bl.b == ub {
		strike++
	} else if bl.b == ua || bl.b == uc {
		ball++
	}
	if bl.c == uc {
		strike++
	} else if bl.c == ub || bl.c == ua {
		ball++
	}
	return strike, ball, nil
}

func main() {
	utils.Seed(time.Now().UnixNano())

	game := Baller{}

	game.Setup()
	fmt.Println("숫자를 생성했습니다")

	for {
		var user_str string

		fmt.Print("정답을 예측하세요 : ")
		for cnt, err := fmt.Scan(&user_str); cnt != 1 || len(user_str) != 3; {
			utils.IsSafe(err)
			fmt.Println("세 자리 수를 입력해주세요")
			utils.Flush()

			fmt.Print("정답을 예측하세요 : ")
			cnt, err = fmt.Scan(&user_str)
			utils.IsSafe(err)
		}

		s, b, err := game.Check(user_str)
		utils.IsSafe(err)

		if s == 3 {
			fmt.Println("정답입니다")
			return
		}

		fmt.Printf("%v -> S%v, B%v\n", user_str, s, b)
	}
}
