/*
   @Time : 20-6-20 下午8:05
   @Author : young
   @File : main.go
   @Software: GoLand
*/
package main

import "log"

func main() {
	log.Println("offer 04")
	var matrix = [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	var target = 5

	m := len(matrix)
	n := len(matrix[0])
	log.Println(m, n)
	log.Println(findNumberIn2DArray(matrix, target))
}

func findNumberIn2DArray1(matrix [][]int, target int) bool {
	ret := false

	for _, iv := range matrix {
		for _, jv := range iv {
			if jv == target {
				ret = true
				return ret
			}
		}
	}
	return ret
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	ret := false

	m := len(matrix)
	n := 0
	if m > 0 {
		n = len(matrix[0])
	}

	// 右上角开始
	for i, j := 0, n-1; i < m && j >=0 ; {
		if matrix[i][j] == target {
			ret = true
			break
		}
		if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return ret
}
