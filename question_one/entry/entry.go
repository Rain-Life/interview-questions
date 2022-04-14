package main

import (
	"fmt"
	"interview-questions/question_one"
)

func main() {
	question_one.Perm([]rune("ABC"), func(strRune []rune) {
		fmt.Println(string(strRune))
	})
}
