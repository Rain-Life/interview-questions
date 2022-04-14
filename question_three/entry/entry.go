package main

import (
	"fmt"
	"interview-questions/question_three"
)

func main() {
	fmt.Printf("rand_13实现rand_5：%d\n", question_three.Rand13ToRand5())
	fmt.Printf("rand_5实现rand_13：%d\n", question_three.Rand5ToRand13())
}
