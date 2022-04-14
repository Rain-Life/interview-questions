package question_three

import (
	"math"
	"math/rand"
	"time"
)

func Rand13ToRand5() int {
	randVal := math.MaxInt64
	for randVal > 5 {
		randVal = randNumFun(1, 13)
	}

	return randVal
}

func Rand5ToRand13() int {
	return randNumFun(1, 5) + 2*(randNumFun(1, 5)-1)
}

//生成指定区间的随机数
func randNumFun(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max) + min
}