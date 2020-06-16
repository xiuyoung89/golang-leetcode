/*
   @Time : 2020/6/16 下午1:09
   @Author : young
   @File : main
   @Software: GoLand
*/
package main

import "log"

func main() {
	log.Printf("this reverse interger.")
	x := 1534236469
	log.Println(reverse(x))
}

func reverse(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10

		if !( -(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}

		x /= 10
	}
	return y
}
