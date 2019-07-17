/*
@Time : 19-7-16 下午7:00
@Author : young
@File : combinations.go
@Software: GoLand
*/
package combinations

//func comb(begin, end, k int) [][]int {
//	var r [][]int
//	for i := begin; i <= end; i++ {
//		if 1 == k {
//			r = append(r, []int{i})
//			continue
//		}
//		suf := comb(i+1, end, k-1)
//		for _, j := range suf {
//			r = append(r, append([]int{i}, j...))
//		}
//	}
//	return r
//}
//
//func Combine(n int, k int) [][]int {
//	return comb(1, n, k)
//}

func comb(begin, end, last int) [][]int {
	var ret [][]int

	for i := begin; i <= end; i++ {
		if 1 == last {
			ret = append(ret, []int{i})
			continue
		}

		sub := comb(i+1, end, last-1)
		for _, j := range sub {

			ret = append(ret, append([]int{i}, j...))
		}
	}
	return ret
}

func Combine(n, k int) [][]int {
	return comb(1, n, k)
}
