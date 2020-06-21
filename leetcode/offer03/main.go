/*
   @Time : 20-6-20 下午7:07
   @Author : young
   @File : main.go
   @Software: GoLand
*/
package main

import "log"

func main() {
	log.Println("offer 03")
	var nums = []int{1,1,1}
	log.Println(findRepeatNumber(nums))
}

func findRepeatNumber01(nums []int) int {

	var ret = -1
	if len(nums) < 2 || len(nums) > 100000 {
		return ret
	}

	for i := 0; i < len(nums); i++ {

		

		for j := i + 1; j < len(nums); j++ {

			if nums[j] == nums[i] {
				ret = nums[i]
				break
			}
		}
		if ret != -1 {
			break
		}
	}

	return ret
}


func findRepeatNumber(nums []int) int {
	var numKeys = make(map[int]bool,0)

	for i , v := range nums{
		if _, ok := numKeys[v]; ok {
			return v
		}else {
			numKeys[v] = true
		}
	}
}