/*
   @Time : 20-6-20 下午9:10
   @Author : young
   @File : main.go
   @Software: GoLand
*/
package main

import "log"

func main() {
	log.Printf("offer 05")
	var ss = "We are happy."
	log.Println(replaceSpace(ss))
}

func replaceSpace(s string) string {
	repStr := ""
	for i, v := range s {
		if v == ' ' {
			repStr += "%20"
		} else {
			repStr += string(s[i])
		}
	}
	return repStr
}
