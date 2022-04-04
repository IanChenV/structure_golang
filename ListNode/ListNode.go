package structure

import (
	"fmt"
)

// 链表的数据结构
type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表转数组
func List2Ints(head *ListNode) []int {
	// 链条深度限制，链条深度超出此限制，会panic
	limit := 100
	time := 0
	res := []int{}

	for head != nil {
		time++
		if time > limit {
			msg := fmt.Sprintf("链条深度超过%d。可能出现环状链条。请检查错误，或者放宽l2i函数的limit限制", limit)
			panic(msg)
		}
		res = append(res, head.Val)
		head = head.Next
	}
	return res

}

// 数组转链表
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	dump := &ListNode{Val: -1}
	cur := dump

	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dump.Next
}

// 链表节点查找
func (l *ListNode) GetNodeWith(val int) *ListNode {
	cur := l
	for cur != nil {
		if cur.Val == val {
			break
		}
		cur = cur.Next
	}
	return cur
}

// 构建环状链表
func Ints2ListWithCycle(nums []int, pos int) *ListNode {
	head := Ints2List(nums)
	if pos == -1 {
		return head
	}

	// 环的位置
	cur := head
	for pos > 0 {
		cur = cur.Next
		pos--
	}

	// 尾部节点
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}

	// 尾部指向了环所在的节点
	tail.Next = cur
	return head
}
