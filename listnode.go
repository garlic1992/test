package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// PrintList 打印ListNode
func PrintList(head *ListNode) {
	ans := head
	for ans != nil {
		fmt.Println(ans)
		ans = ans.Next
	}
}

// InitListNode 初始化ListNode
func InitListNode(nums []int) *ListNode {
	dummy := &ListNode{0, nil}
	ans := dummy
	for _, num := range nums {
		dummy.Next = &ListNode{num, nil}
		dummy = dummy.Next
	}
	return ans.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	ans := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			dummy.Next = list1
			list1 = list1.Next
		} else {
			dummy.Next = list2
			list2 = list2.Next
		}
		dummy = dummy.Next
	}
	for list1 != nil {
		dummy.Next = list1
		list1 = list1.Next
		dummy = dummy.Next
	}
	for list2 != nil {
		dummy.Next = list2
		list2 = list2.Next
		dummy = dummy.Next
	}
	return ans.Next
}
