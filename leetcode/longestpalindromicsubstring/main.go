/*
   @Time : 20-6-15 下午11:06
   @Author : young
   @File : main.go
   @Software: GoLand
*/
package main

import (
	"log"
)

func main() {
	log.Printf("this is longets palindromic substring.")

	srcString := ""
	ret := longestPalindrome(srcString)

	log.Printf("src string : %s, get sub string: %s.", srcString, ret)
}

func longestPalindrome(s string) string {

	begin, end := 0, 0
	if len(s) <= 0 {
		return s
	}

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindromic(s, i, j) {
				if j-i > end-begin {
					begin = i
					end = j
				}
			}
		}
	}
	return s[begin : end+1]
}

func isPalindromic(s string, begin int, end int) (ok bool) {
	ok = true
	for {
		if begin >= end {
			break
		}

		if s[begin] == s[end] {
			begin++
			end--
		} else {
			ok = false
			break
		}
	}
	return
}
