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
	for i := 0; i < 2; i++ {
		node := &Node{
			Item: i,
		}
		if !sl.Append(node) {
			t.Error("插入出错")
		}
	}
	sl.Tail.Next = sl.Head
	t.Log(sl.HasLoop())
}

func TestMerge(t *testing.T) {
	sl1 := &SingleList{}
	sl2 := &SingleList{}
	sl1.Init()
	sl2.Init()
	for i := 1; i < 13; i++ {
		node := &Node{
			Item: i,
		}
		if i%2 == 0 {
			if !sl2.Append(node) {
				t.Error("插入出错")
			}
		}
	}
	for i := 2; i < 14; i++ {
		node := &Node{
			Item: i,
		}
		if i%2 == 0 {
			if !sl1.Append(node) {
				t.Error("插入出错")
			}
		}
	}
	sl1.Print()
	fmt.Println("")
	//sl1.Merge(sl2, IsRecursive)
	sl1.Merge(sl2, -1)
	sl1.Print()
	fmt.Println("")
	fmt.Println(sl1)
}

func TestDelByReciprocal(t *testing.T) {
	sl := &SingleList{}
	sl.Init()
	for i := 0; i < 1; i++ {
		node := &Node{
			Item: i,
		}
		if !sl.Append(node) {
			t.Error("插入出错")
		}
	}
	sl.Print()
	fmt.Println(sl.DelByReciprocal(1))
	sl.Print()
}

func TestFindMiddle(t *testing.T) {
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
	sl.Print()
	fmt.Println("------")
	fmt.Println(sl.FindMiddle())
}
