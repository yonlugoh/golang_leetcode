package _323__Maximum_69_Number

import "math"

func maximum69Number(num int) int {

	temp := num
	indexSix := -1
	i := 0
	for temp > 0 {
		if temp%10 == 6 {
			indexSix = i
		}
		temp = temp / 10
		i++
	}

	if indexSix == -1 {
		return num
	} else {
		return num + 3*int(math.Pow10(indexSix))
	}

}
