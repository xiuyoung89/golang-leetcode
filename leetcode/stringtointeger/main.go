/*
   @Time : 2020/6/16 下午1:57
   @Author : young
   @File : main.go
   @Software: GoLand
*/
package main

import (
	"log"
	"math"
	"strings"
)

func main() {
	log.Printf("this is string to integer")
	var s = "-9223372036854775808"
	log.Println(myAtoi(s))
}

func myAtoi(str string) int {
	return convert(clean(str))
	//if len(str) < 0 {
	//	return 0
	//}

	//log.Println(str[0])
	//if str[0] < {
	//	return 0
	//}

	//var sign int = 1
	//var ret int = 0
	//for i, _ := range str {
	//	if i == 0 && str[i] == '-' {
	//		sign = -1
	//	}
	//	if str[i] > '0' && str[i] < '9' {
	//		ret = ret*10 + int(str[i])
	//	}
	//}
	//if sign > 0 && ret > (1<<31-1) {
	//	return int(1<<31 -1)
	//}
	//
	//if sign < 0 && ret > (1 << 31) {
	//	return sign * int(1<<31)
	//}
	//return ret * sign
}

func clean(s string) (sign int, abs string) {
	//
	s = strings.TrimSpace(s)
	if len(s) <= 0 {
		return 1, ""
	}

	switch s[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, abs = 1, s[:]
	case '+':
		sign, abs = 1, s[1:]
	case '-':
		sign, abs = -1, s[1:]
	default:
		sign, abs = 1, ""
	}

	for j, v := range abs {
		if v < '0' || v > '9' {
			abs = abs[0:j]
			break
		}
	}

	return
}

func convert(sign int, s string) (ret int) {
	//var ret int
	for _, v := range s {
		ret = ret*10 + int(v-'0')

		switch  {
		case sign ==1 && ret > math.MaxInt32:
			return  math.MaxInt32
		case sign == -1 && ret *sign < math.MinInt32:
			return math.MinInt32
		}
	}
	ret = sign * ret

	return
}
