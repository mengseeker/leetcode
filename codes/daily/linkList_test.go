/*
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。


示例：

MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
linkedList.get(1);            //返回2
linkedList.deleteAtIndex(1);  //现在链表是1-> 3
linkedList.get(1);            //返回3


提示：

所有val值都在 [1, 1000] 之内。
操作次数将在  [1, 1000] 之内。
请不要使用内置的 LinkedList 库。

*/
package daily

import (
	"testing"
)

type MyLinkedList struct {
	Head *Node
	Tail *Node
	Len  int
}

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (l *MyLinkedList) At(index int) *Node {
	// fmt.Println("len: ", l.Len, "index: ", index)
	if index < 0 || index >= l.Len {
		return nil
	}
	var cur = l.Head
	if index <= l.Len/2 {
		for i := 0; i < index; i++ {
			// fmt.Println("len: ", l.Len, "index: ", index, "i: ", i)
			cur = cur.Next
		}
		return cur
	}
	cur = l.Tail
	for i := index + 1; i < l.Len; i++ {
		cur = cur.Prev
	}
	return cur
}

func (l *MyLinkedList) Get(index int) int {
	n := l.At(index)
	if n != nil {
		// fmt.Println("Get: ", index, "->", n.Val)
		return n.Val
	}
	// fmt.Println("AddAtHead: ", index, "->", -1)
	return -1
}

func (l *MyLinkedList) AddAtHead(val int) {
	// fmt.Println("AddAtHead: ", val)
	n := Node{Val: val, Next: l.Head}
	if l.Head != nil {
		l.Head.Prev = &n
	} else {
		l.Tail = &n
	}
	l.Head = &n
	l.Len++
}

func (l *MyLinkedList) AddAtTail(val int) {
	// fmt.Println("AddAtTail: ", val)
	n := Node{Val: val, Prev: l.Tail}
	if l.Tail != nil {
		l.Tail.Next = &n
	} else {
		l.Head = &n
	}
	l.Tail = &n
	l.Len++
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		l.AddAtHead(val)
	} else if index == l.Len {
		l.AddAtTail(val)
	} else if index < l.Len {
		// fmt.Println("AddAtIndex: ", index)
		n := l.At(index)
		nn := Node{Val: val, Next: n, Prev: n.Prev}
		n.Prev.Next = &nn
		n.Prev = &nn
		l.Len++
	}
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	// fmt.Println("DeleteAtIndex: ", index)
	n := l.At(index)
	if n != nil {
		if n.Prev != nil {
			n.Prev.Next = n.Next
		} else {
			// head
			l.Head = n.Next
		}
		if n.Next != nil {
			n.Next.Prev = n.Prev
		} else {
			l.Tail = n.Prev
		}
		l.Len--
	}
}

func (l *MyLinkedList) Print() {
	i := 1
	cur := l.Head
	for {
		print(cur.Val, ",")
		if cur.Next == nil {
			break
		}
		cur = cur.Next
		i++
	}
	print("LLen: ", l.Len, "/", i, "\n")

	i = 1
	cur = l.Tail
	vals := []int{}
	for {
		vals = append(vals, cur.Val)
		if cur.Prev == nil {
			break
		}
		cur = cur.Prev
		i++
	}
	for i := len(vals) - 1; i >= 0; i-- {
		print(vals[i], ",")
	}
	print("Rlen: ", l.Len, "/", i, "\n")
}

func TestList(t *testing.T) {
	var methods = []string{
		"addAtHead", "addAtTail", "addAtTail", "addAtTail", "addAtTail", "addAtTail", "addAtTail", "deleteAtIndex", "addAtHead", "addAtHead", "get", "addAtTail", "addAtHead", "get", "addAtTail", "addAtIndex", "addAtTail", "addAtHead", "addAtHead", "addAtHead", "get", "addAtIndex", "addAtHead", "get", "addAtHead", "deleteAtIndex", "addAtHead", "addAtTail", "addAtTail", "addAtIndex", "addAtTail", "addAtHead", "get", "addAtTail", "deleteAtIndex", "addAtIndex", "deleteAtIndex", "addAtHead", "addAtTail", "addAtHead", "addAtHead", "addAtTail", "addAtTail", "get", "get", "addAtHead", "addAtTail", "addAtTail", "addAtTail", "addAtIndex", "get", "addAtHead", "addAtIndex", "addAtHead", "addAtTail", "addAtTail", "addAtIndex", "deleteAtIndex", "addAtIndex", "addAtHead", "addAtHead", "deleteAtIndex", "addAtTail", "deleteAtIndex", "addAtIndex", "addAtTail", "addAtHead", "get", "addAtIndex", "addAtTail", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "deleteAtIndex", "get", "get", "addAtHead", "get", "addAtTail", "addAtTail", "addAtIndex", "addAtIndex", "addAtHead", "addAtTail", "addAtTail", "get", "addAtIndex", "addAtHead", "deleteAtIndex", "addAtTail", "get", "addAtHead", "get", "addAtHead", "deleteAtIndex", "get", "addAtTail", "addAtTail",
	}
	var params = [][]int{
		{38}, {66}, {61}, {76}, {26}, {37}, {8}, {5}, {4}, {45}, {4}, {85}, {37}, {5}, {93}, {10, 23}, {21}, {52}, {15}, {47}, {12}, {6, 24}, {64}, {4}, {31}, {6}, {40}, {17}, {15}, {19, 2}, {11}, {86}, {17}, {55}, {15}, {14, 95}, {22}, {66}, {95}, {8}, {47}, {23}, {39}, {30}, {27}, {0}, {99}, {45}, {4}, {9, 11}, {6}, {81}, {18, 32}, {20}, {13}, {42}, {37, 91}, {36}, {10, 37}, {96}, {57}, {20}, {89}, {18}, {41, 5}, {23}, {75}, {7}, {25, 51}, {48}, {46}, {29}, {85}, {82}, {6}, {38}, {14}, {1}, {12}, {42}, {42}, {83}, {13}, {14, 20}, {17, 34}, {36}, {58}, {2}, {38}, {33, 59}, {37}, {15}, {64}, {56}, {0}, {40}, {92}, {63}, {35}, {62}, {32},
	}
	list := Constructor()
	for i, m := range methods {
		switch m {
		case "addAtHead":
			list.AddAtHead(params[i][0])
		case "addAtTail":
			list.AddAtTail(params[i][0])
		case "deleteAtIndex":
			list.DeleteAtIndex(params[i][0])
		case "addAtIndex":
			list.AddAtIndex(params[i][0], params[i][1])
		case "get":
			list.Get(params[i][0])
		}
		// list.Print()
	}

}

/**
* Your MyLinkedList object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.Get(index);
* obj.AddAtHead(val);
* obj.AddAtTail(val);
* obj.AddAtIndex(index,val);
* obj.DeleteAtIndex(index);
 */
