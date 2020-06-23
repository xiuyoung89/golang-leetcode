/*
   @Time : 2020/6/22 下午2:55
   @Author : young
   @File : main
   @Software: GoLand
*/
package main

import (
	"log"
	"math"
)

func main() {
	log.Println("offer 59")
	//nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	//k := 3
	nums := []int{-7,-8,7,5,7,1,6,0}
	k:=4
	log.Println(maxSlidingWindow(nums, k))
}

type helpQue []int

func (s *helpQue) PushTail(x int) {
	*s = append(*s, x)
	//if len(*s) == 0 {
	//	*s = append(*s, x)
	//} else {
	//	if x < (*s)[len(*s)-1] {
	//		*s = append(*s, x)
	//	}
	//}
}

func (s *helpQue) PushHead(x int) {
	if len(*s) == 0 {
		*s = append(*s, x)
	} else {
		tmp := []int{x}
		*s = append(tmp, *s...)
	}
}

func (s *helpQue) PopLess(x int) {
	for len(*s) != 0 {
		if x > (*s)[len(*s)-1] {
			*s = (*s)[:len(*s)-1]
		} else {
			//*s = append(*s, x)
			break
		}
	}

}

func (s *helpQue) PopTail() int {
	ret := math.MinInt32
	if len(*s) != 0 {
		ret = (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
	}
	return ret
}

func (s *helpQue) PopHead() {
	if len(*s) != 0 {
		*s = (*s)[1:len(*s)]
	}
}

func (s *helpQue) Head() int {
	if len(*s) != 0 {
		return (*s)[0]
	}
	return math.MaxInt32
}

func (s *helpQue) Tail() int {
	if len(*s) != 0 {
		return (*s)[len(*s)-1]
	} else {
		return math.MinInt32
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	var ret []int
	// 维护一张非严格单调队列
	helpQue := helpQue{}
	if len(nums) == 0 {
		return ret
	}

	// 初始化第一个
	for i := 0; i < k; i++ {
		for len(helpQue) > 0 && nums[i] > helpQue.Tail() {
			helpQue.PopTail()
		}
		helpQue.PushTail(nums[i])
	}
	ret = append(ret, helpQue.Head())

	if len(nums)< k {
		return ret
	}

	// i and j
	for i, j := 0, k; j < len(nums); {

		if nums[i] == helpQue.Head() {
			helpQue.PopHead()
		}

		// pop less
		helpQue.PopLess(nums[j])
		helpQue.PushTail(nums[j])
		ret = append(ret, helpQue.Head())
		i++
		j++
	}
	return ret
}
