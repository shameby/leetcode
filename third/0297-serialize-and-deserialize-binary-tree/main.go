package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	t1 := &TreeNode{Val: 4}
	t2 := &TreeNode{Val: 2}
	t3 := &TreeNode{Val: 7}
	// t4 := &TreeNode{Val: 1}
	t5 := &TreeNode{Val: 3}
	t6 := &TreeNode{Val: 6}
	t7 := &TreeNode{Val: 9}
	t1.Left = t2
	t1.Right = t3
	// t2.Left = t4
	t2.Right = t5
	t3.Left = t6
	t3.Right = t7
	c1 := Constructor()
	fmt.Println(c1.serialize(t1))
	fmt.Println(c1.deserialize("4,2,7,#,3,6,9,#,#,#,#,#,#"))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	l []string
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	q := []*TreeNode{root}
	res := make([]string, 0)
	for len(q) != 0 {
		n := q[0]
		q = q[1:]
		if n == nil {
			res = append(res, "#")
		} else {
			res = append(res, strconv.Itoa(n.Val))
			q = append(q, n.Left)
			q = append(q, n.Right)
		}
	}
	return strings.Join(res, ",")
}

func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	dataL := strings.Split(data, ",")
	rootVal, _ := strconv.Atoi(dataL[0])
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}
	idx := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		idx++
		if dataL[idx] != "#" {
			val, _ := strconv.Atoi(dataL[idx])
			node.Left = &TreeNode{Val:val}
			queue = append(queue, node.Left)
		}
		idx++
		if dataL[idx] != "#" {
			val, _ := strconv.Atoi(dataL[idx])
			node.Right = &TreeNode{Val:val}
			queue = append(queue, node.Right)
		}
	}
	return root
}
