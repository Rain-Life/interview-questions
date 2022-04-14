package question_one

// Perm() 对 a 形成的每⼀排列调⽤ f().
func Perm(strRune []rune, print func([]rune)) {
	perm(strRune, print, 0)
}

// 对索引 i 从 0 到 len(a) - 1，实现递归函数 perm().
func perm(strRune []rune, print func([]rune), i int) {
	if i == len(strRune)-1 {
		print(strRune)
	}

	hashMap := map[rune]bool{}
	for idx := i; idx < len(strRune); idx++ {
		if !hashMap[strRune[idx]] {
			strRune[i], strRune[idx] = strRune[idx], strRune[i]
			hashMap[strRune[i]] = true
			perm(strRune, print, i + 1)
			strRune[i], strRune[idx] = strRune[idx], strRune[i]
		}
	}
}


