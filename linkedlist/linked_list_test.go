package linkedlist

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	//1.构建初始单链表数据
	sl := &SingleList{}
	sl.Init()
	for i := 0; i < 3; i++ {
		node := &Node{
			Item: i,
		}
		if !sl.Append(node) {
			t.Error("插入出错")
		}
	}
	//打印初始链表
	fmt.Println(sl)
	sl.Print()
	fmt.Println("")
	//单链表反转 -- 递归
	sl.Reverse(IsRecursive)
	//单链表反转 -- 非递归
	//sl.Reverse(1)
	//打印反转后链表
	fmt.Println("")
	sl.Print()
	fmt.Println(sl)
}

func TestHasLoop(t *testing.T) {
	//构建有环链表
	sl := &SingleList{}
	sl.Init()
	//全环
	for i := 0; i < 5; i++ {
		node := &Node{
			Item: i,
		}
		if !sl.Append(node) {
			t.Error("插入出错")
		}
	}
	//sl.Tail.Next = sl.Head
	t.Log(sl.HasLoop())
}
