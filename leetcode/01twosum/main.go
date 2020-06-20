/*
@Time : 19-7-11 上午10:25
@Author : young
@File : twosum
@Software: GoLand
*/
package main

import "log"

func main() {
	log.Println("two sum")
}


// 耗时 40ms
//func TwoSum(nums[]int, target int) []int {
//
//	for i, iv := range nums {
//
//		for j, jv := range nums {
//
//			if i == j {
//				continue
//			}
//
//			if iv + jv == target{
//				return  []int{i,j}
//			}
//		}
//	}
//
//	return  nil
//}

func TwoSum(nums []int, target int) []int {

	// map[value]=index
	index := make(map[int]int)

	for i, num := range nums {

		tmp := target - num

		if j, ok := index[tmp]; ok {

			return []int{j, i}
		}

		index[num] = i
	}

	return nil
}