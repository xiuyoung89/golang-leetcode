/*
   @Time : 2020/6/15 下午5:42
   @Author : young
   @File : main.go
   @Software: GoLand
*/
package main

import "log"

func main() {
	var num1 []int = []int{1, 2, 3}
	var num2 []int = []int{3, 6, 7, 8}

	ret := findMedianSortedArrays(num1, num2)
	log.Printf("get ret: %f", ret)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var ret float64 = 0.0
	var ioa int = 0
	var iob int = 0
	var ioc int = 0
	var newArray = []int{0}
	for {
		if ioa < len(nums1) {
			// num1 开始
			if iob < len(nums2) {
				//compare
				if nums1[ioa] < nums2[iob] {
					newArray = append(newArray, nums1[ioa])
					ioa++
				} else {
					newArray = append(newArray, nums2[iob])
					iob++
				}
				ioc++

			} else {
				//copy(newArray[ioc:], nums1[ioa:])
				newArray = append(newArray[:], nums1[ioa:]...)
				break
			}
		} else {
			if iob < len(nums2) {
				newArray = append(newArray[:], nums2[iob:]...)
				//log.Printf("copy end n: %d", nn)
				break
			}
		}
	}

	if len(newArray) > 0 {
		if len(newArray)%2 == 0 {
			ret = float64(newArray[len(newArray)/2-1]+newArray[len(newArray)/2]) / 2
		} else {
			ret = float64(newArray[(len(newArray)-1)/2])
		}
	}
	log.Printf("ttt:%v", newArray)
	return ret
}
