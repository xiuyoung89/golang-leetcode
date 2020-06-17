/*
   @Time : 2020/6/17 上午10:54
   @Author : young
   @File : main.go
   @Software: GoLand
*/
package main

import "log"

func main() {
	log.Println("42 rain trap.")
	var test = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	log.Println(trap(test))
}

func trap(height []int) int {
	index, sum := 0, 0
	var stack = make([]int, 0)
	for index < len(height) {

		for len(stack) != 0 && height[index] > height[stack[len(stack)-1]] {
			// 计算水滴数量 并且 让数据出栈
			h := height[stack[len(stack)-1]]
			stack = stack[0 : len(stack)-1]

			if len(stack) == 0 {
				break
			}
			dis := index - stack[len(stack)-1] - 1
			min := min(height[index], height[stack[len(stack)-1]])
			sum += dis*(min - h)

		}
		// 进桟
		stack = append(stack, index)
		// 移动游标
		index++
	}
	return sum
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
