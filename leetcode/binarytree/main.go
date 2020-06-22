/*
   @Time : 2020/6/19 下午2:30
   @Author : young
   @File : main
   @Software: GoLand
*/
package main

import "log"

func main()  {
	log.Println("binary tree")
}

type TreeNode struct {
	Var int
	Left *TreeNode
	Right *TreeNode
}

// 前序遍历, root->left->right
// 递归遍历
func preOrderTraversal(root * TreeNode)  {
	if root == nil {
		return
	}

	log.Println(root.Var)
	preOrderTraversal(root.Left)
	preOrderTraversal(root.Right)
}

// 非递归遍历
func preOrderTraversal1(root *TreeNode)  {
	for root != nil {
		log.Println(root.Var)

		for root.Left != nil{
			root.Left =
		}
	}
}