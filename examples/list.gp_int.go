///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/vipally/gogp]
// Last update at: [Wed Nov 23 2016 17:48:17]
// Generate from:
//   [github.com/vipally/gogp/examples/gp/list.gp]
//   [github.com/vipally/gogp/examples/example.gpg] [list_int]
//
// Tool [github.com/vipally/gogp] info:
// CopyRight 2016 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Blog    : http://blog.csdn.net/vipally
// Site    : https://github.com/vipally
// BuildAt : [Oct  8 2016 10:34:35]
// Version : 3.0.0.final
//
///////////////////////////////////////////////////////////////////

package examples

////////////////////////////////////////////////////////////////////////////////

var gIntListGbl struct {
	cmp CmpInt
}

func init() {
	gIntListGbl.cmp = gIntListGbl.cmp.CreateByName("")
}

//double-way cycle link list node
type IntListNode struct {
	int
	prev, next *IntListNode
}

func (this *IntListNode) Get() int {
	return this.int
}

func (this *IntListNode) Set(v int) (old int) {
	old, this.int = this.int, v
	return
}

func (this *IntListNode) Next() (r *IntListNode) {
	if this != nil {
		r = this.next
	}
	return
}

func (this *IntListNode) Prev() (r *IntListNode) {
	if this != nil {
		r = this.prev
	}
	return
}

type IntListNodeVisitor struct {
	node, head *IntListNode
}

func (this *IntListNodeVisitor) Reset() {
	this.node = nil
}

func (this *IntListNodeVisitor) Next() (ok bool) {
	if this.node == nil {
		if ok = this.head != nil; ok {
			this.node = this.head
		}
	} else {
		this.node = this.node.next
		ok = this.node != this.head
	}
	return
}

func (this *IntListNodeVisitor) Prev() (ok bool) {
	if this.node == nil {
		if ok = this.head != nil; ok {
			this.node = this.head.prev
		}
	} else {
		this.node = this.node.prev
		ok = this.node != this.head.prev
	}
	return
}

func (this *IntListNodeVisitor) Get() *IntListNode {
	return this.node
}

func (this *IntList) Visitor() *IntListNodeVisitor {
	n := &IntListNodeVisitor{node: nil, head: this.head}
	return n
}

//list object
type IntList struct {
	head *IntListNode
}

//new object
func NewIntList() *IntList {
	return &IntList{}
}

//func (this *IntList) Len() int {
//	return 0
//}

func (this *IntList) Front() *IntListNode {
	return this.head
}
func (this *IntList) Back() (r *IntListNode) {
	if this.head != nil {
		r = this.head.prev
	}
	return
}

func (this *IntList) Clear() {
	this.head = nil
}

func (this *IntList) RotateForward() {
	if this.head != nil {
		this.head = this.head.next
	}
}
func (this *IntList) RotateBackward() {
	if this.head != nil {
		this.head = this.head.prev
	}
}

func (this *IntList) PushFront(v int) *IntListNode {
	n := &IntListNode{int: v}
	return this.InsertFront(n)
}

func (this *IntList) PushBack(v int) *IntListNode {
	n := &IntListNode{int: v}
	return this.InsertBack(n)
}

func (this *IntList) InsertFront(node *IntListNode) (n *IntListNode) {
	if n = this.InsertBack(node); n != nil {
		this.RotateBackward()
	}
	return
}

func (this *IntList) InsertBack(node *IntListNode) (n *IntListNode) {
	if n = node; n != nil {
		if this.head == nil {
			this.head, n.next, n.prev = n, n, n
		} else {
			n.next = this.head
			n.prev = this.head.prev
			this.head.prev.next = n
			this.head.prev = n
		}
	}
	return
}

func (this *IntList) PopFront() (v int, ok bool) {
	if n := this.Remove(this.Front()); n != nil {
		v, ok = n.Get(), true
	}
	return
}

func (this *IntList) PopBack() (v int, ok bool) {
	if n := this.Remove(this.Back()); n != nil {
		v, ok = n.Get(), true
	}
	return
}

func (this *IntList) InsertFrontList(other *IntList) (ok bool) {
	rotate := !this.Empty()
	if ok = this.InsertBackList(other); ok && rotate {
		this.RotateBackward()
	}
	return
}

func (this *IntList) InsertBackList(other *IntList) (ok bool) {
	if ok = !other.Empty(); ok {
		if this.Empty() {
			this.head = other.head
		} else {
			myback, oback := this.Back(), other.Back()
			myback.next = other.head
			oback.next = this.head
			other.head.prev = myback
			this.head.prev = oback
		}
		other.Clear()
	}
	return
}

func (this *IntList) PushBefore(v int, mark *IntListNode) (n *IntListNode) {
	if mark != nil {
		n = &IntListNode{int: v}
		n = this.InsertBefore(n, mark)
	}
	return
}

func (this *IntList) PushAfter(v int, mark *IntListNode) (n *IntListNode) {
	if mark != nil {
		n = &IntListNode{int: v}
		n = this.InsertAfter(n, mark)
	}
	return
}

func (this *IntList) InsertBefore(node, mark *IntListNode) (n *IntListNode) {
	if n = node; node != nil && mark != nil {
		n.next = mark
		n.prev = mark.prev
		mark.prev = n
		if this.head == mark {
			this.RotateBackward()
		}
	}
	return
}

func (this *IntList) InsertAfter(node, mark *IntListNode) (n *IntListNode) {
	if n = node; node != nil && mark != nil {
		n.next = mark.next
		n.prev = mark
		mark.next = n
	}
	return
}

func (this *IntList) RemoveFront() (n *IntListNode) {
	return this.Remove(this.Front())
}

func (this *IntList) RemoveBack() (n *IntListNode) {
	return this.Remove(this.Back())
}

func (this *IntList) Remove(node *IntListNode) (n *IntListNode) {
	if node != nil && node.next != nil && node.prev != nil {
		n = node
		if node.next == node {
			this.head = nil
		} else if node == this.head {
			this.head = node.next
		}
		node.next.prev = node.prev
		node.prev.next = node.next
		node.next, node.prev = nil, nil
	}
	return
}

func (this *IntList) Reachable(node, dest *IntListNode) (ok bool) {
	if ok = (node == dest) && node != nil; !ok && node != nil && dest != nil {
		for p := node; p != nil && p != node; p = p.next {
			if ok = (p == dest); ok {
				break
			}
		}
	}
	return
}

func (this *IntList) IsValidNode(node *IntListNode) bool {
	return this.Reachable(this.head, node)
}

func (this *IntList) MoveFront(node *IntListNode) (r *IntListNode) {
	if r = this.MoveBack(node); r != nil {
		this.RotateBackward()
	}
	return
}

func (this *IntList) MoveBack(node *IntListNode) (r *IntListNode) {
	if node != nil && node.next != nil && node.prev != nil && node.next != node { //bug:node is back?
		node.next.prev = node.prev
		node.prev.next = node.next
		node.next = this.head
		node.prev = this.head.prev
		this.head.prev.next = node
		this.head.prev = node
		r = node
	}
	return
}

func (this *IntList) MoveBefore(node, mark *IntListNode) (r *IntListNode) {
	if node != nil && mark != nil && node != mark && node.next != mark {
		if r = this.Remove(node); r != nil {
			node.next = mark.prev.next
			node.prev = mark.prev
			mark.prev.next = node
			mark.prev = node
		}
	}
	return
}

func (this *IntList) MoveAfter(node, mark *IntListNode) (r *IntListNode) {
	if node != nil && mark != nil && node != mark && mark.next != node {
		if r = this.Remove(node); r != nil {
			node.next = mark.next
			node.prev = mark
			mark.next.prev = node
			mark.next = node
		}
	}
	return
}

func (this *IntList) Empty() bool {
	return this.head == nil
}

func (this *IntList) Reverse() {
	v := this.Visitor()
	this.Clear()
	for v.Next() {
		this.InsertFront(v.Get())
	}
}

func (this *IntList) Sort() {
	this.mergeSort()
}

//STL's merge sort algorithm for list
func (this *IntList) mergeSort() {
	if nil == this.head || this.head == this.head.next { //0 or 1 element, no need to sort
		return
	}

	var (
		hand, unsorted IntList
		binList        [64]IntList //save temp list that len=2^i
		nFilledBin     = 0
	)

	for unsorted = *this; !unsorted.Empty(); {
		hand.InsertFront(unsorted.RemoveFront())
		i := 0
		for ; i < nFilledBin && !binList[i].Empty(); i++ {
			binList[i].merge(&hand)
			hand, binList[i] = binList[i], hand
		}
		hand, binList[i] = binList[i], hand
		if i == nFilledBin {
			nFilledBin++
		}
	}

	for i := 1; i < nFilledBin; i++ {
		binList[i].merge(&binList[i-1])
	}

	*this = binList[nFilledBin-1]
}

//merge other sorted-list  to this sorted-list
func (this *IntList) merge(other *IntList) {
	if this.Empty() || other.Empty() {
		this.InsertBackList(other)
		return
	}

	p, po := this.Front(), other.Front()
	for p != nil && po != nil {
		if gIntListGbl.cmp.F(po.int, p.int) {
			n := other.RemoveFront()
			po = other.Front()
			p = this.InsertBefore(n, p)
		} else {
			if p = p.next; p == this.Front() {
				p = nil
			}
		}
	}
	if po != nil {
		this.InsertBackList(other)
	}
}
