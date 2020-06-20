/*
   @Time : 2020/6/16 上午10:48
   @Author : young
   @File : main.go
   @Software: GoLand
*/
package main

import "log"

func main() {
	log.Printf("this is zigzag conversion.")

	var s = "AB"
	ret := convert(s, 1)
	log.Printf("get ret: %s", ret)
}

func convert(s string, numRows int) string {

	if len(s) < numRows || numRows ==1	{
		return s
	}

	var tmp = make([]string, numRows)
	var ret string
	revSign := 1
	rowStep := 0
	for _, v := range s {

		tmp[rowStep] += string(v)
		rowStep += revSign
		if rowStep == 0 || rowStep == numRows-1 {
			revSign = -revSign
		}
	}

	//合并数据
	for j, _ := range tmp {
		ret += tmp[j]
	}
	return ret
}
