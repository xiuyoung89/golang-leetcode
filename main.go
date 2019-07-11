/*
@Time : 19-7-11 上午10:22
@Author : young
@File : main.go
@Software: GoLand
*/
package main

import (
	"fmt"
	"golang-leetcode/leetcode/twosum"
)

func main()  {
	fmt.Print("This is leetcode project.")

	var nums =  []int {2, 7, 11, 15}
	var target int = 13
	ret := twosum.TwoSum(nums, target)

	fmt.Println("ret: \n", ret)

}