package homework1013

import "testing"

// middleNode函数返回链表中间结点 
// 先遍历链表求出链表长度,再次遍历链表，遍历到len/2次时返回当前节点记为中间节点
func middleNode(head *ListNode) *ListNode {
	p := head
	len := 0
	for p != nil {
		len++
		p = p.Next
	}

	for i := 0; i < len/2; i++ {
		head = head.Next
	}
	return head
}

// 定义链表结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

//数组插入链表中
func newList(args ...int) *ListNode {
	var head ListNode
	var p *ListNode = &head
	for _, v := range args {
		node := new(ListNode) //结构体使用new来开辟内存空间  （是一个地址）
		p.Next = node
		node.Val = v
		p = p.Next

	}
	return head.Next
}

func TestMiddleNode(t *testing.T) {
	data := []struct {
		l      *ListNode
		answer int
	}{

		{l: newList(1, 2, 3, 4, 5), answer: 3},
		{l: newList(1, 2, 3, 4, 5, 6), answer: 4},
	}

	for _, v := range data {
		result := middleNode(v.l)
		if result.Val != v.answer { //判断值是否相等
			t.Fatalf("expect:[%v] != result[%v]", result, v.answer)
		}
	}
}
