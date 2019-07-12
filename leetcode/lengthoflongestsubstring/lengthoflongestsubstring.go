/*
@Time : 19-7-12 上午11:03
@Author : young
@File : lengthoflongestsubstring
@Software: GoLand
*/
package lengthoflongestsubstring

import "strings"

func LengthOfLongestSubstring(s string) int {
	//record := make(map[string]int)
	//count := 0
	//var maxCount int = 0;
	//for i, v := range s {
	//
	//	if _, ok := record[string(v)]; ok {
	//
	//		count = i - record[string(v)]
	//		//更新index
	//		record[string(v)] = i
	//	} else {
	//		count += 1;
	//	}
	//
	//	if maxCount < count {
	//		maxCount = count
	//	}
	//
	//	record[string(v)] = i
	//}
	//return maxCount

	// 遇到不重复的字符时，跟之前的不重复字符数做大小对比并最大替换重复数，如此便可以得出最长字串数
	num, j, t := 0, 0, 0
	for i := 0; i < len(s); i++ {
		t = strings.Index(s[j:i], string(s[i]))
		if t == -1 {
			if num < (i - j + 1) {
				num = (i - j + 1)
			}
		} else {
			j = j + t + 1
		}
	}
	return num
}
