///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/vipally/gogp]
// Last update at: [Sat Apr 01 2017 22:48:09]
// Generate from:
//   [github.com/vipally/gogp/examples/gp/list.gp]
//   [github.com/vipally/gogp/examples/example2/example2.gpg] [person]
//
// Tool [github.com/vipally/gogp] info:
// CopyRight 2016 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Blog    : http://blog.csdn.net/vipally
// Site    : https://github.com/vipally
// BuildAt : 
// Version : 3.0.0.final
// 
///////////////////////////////////////////////////////////////////

package example2

////////////////////////////////////////////////////////////////////////////////

var gPersonListGbl struct {
	cmp CmpPerson
}

func init() {
	gPersonListGbl.cmp = gPersonListGbl.cmp.CreateByName("")
}

//double-way cycle link list node
type PersonListNode struct {
	*Person
	prev, next *PersonListNode
}

func (this *PersonListNode) Get() *Person {
	return this.Person
}

func (this *PersonListNode) Set(v *Person) (old *Person) {
	old, this.Person = this.Person, v
	return
}

func (this *PersonListNode) Next() (r *PersonListNode) {
	if this != nil {
		r = this.next
	}
	return
}

func (this *PersonListNode) Prev() (r *PersonListNode) {
	if this != nil {
		r = this.prev
	}
	return
}

type PersonListNodeVisitor struct {
	node, head *PersonListNode
}

func (this *PersonListNodeVisitor) Reset() {
	this.node = nil
}

func (this *PersonListNodeVisitor) Next() (ok bool) {
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

func (this *PersonListNodeVisitor) Prev() (ok bool) {
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

func (this *PersonListNodeVisitor) Get() *PersonListNode {
	return this.node
}

func (this *PersonList) Visitor() *PersonListNodeVisitor {
	n := &PersonListNodeVisitor{node: nil, head: this.head}
	return n
}

//list object
type PersonList struct {
	head *PersonListNode
}

//new object
func NewPersonList() *PersonList {
	return &PersonList{}
}

//func (this *PersonList) Len() int {
//	return 0
//}

func (this *PersonList) Front() *PersonListNode {
	return this.head
}
func (this *PersonList) Back() (r *PersonListNode) {
	if this.head != nil {
		r = this.head.prev
	}
	return
}

func (this *PersonList) Clear() {
	this.head = nil
}

func (this *PersonList) RotateForward() {
	if this.head != nil {
		this.head = this.head.next
	}
}
func (this *PersonList) RotateBackward() {
	if this.head != nil {
		this.head = this.head.prev
	}
}

func (this *PersonList) PushFront(v *Person) *PersonListNode {
	n := &PersonListNode{Person: v}
	return this.InsertFront(n)
}

func (this *PersonList) PushBack(v *Person) *PersonListNode {
	n := &PersonListNode{Person: v}
	return this.InsertBack(n)
}

func (this *PersonList) InsertFront(node *PersonListNode) (n *PersonListNode) {
	if n = this.InsertBack(node); n != nil {
		this.RotateBackward()
	}
	return
}

func (this *PersonList) InsertBack(node *PersonListNode) (n *PersonListNode) {
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

func (this *PersonList) PopFront() (v *Person, ok bool) {
	if n := this.Remove(this.Front()); n != nil {
		v, ok = n.Get(), true
	}
	return
}

func (this *PersonList) PopBack() (v *Person, ok bool) {
	if n := this.Remove(this.Back()); n != nil {
		v, ok = n.Get(), true
	}
	return
}

func (this *PersonList) InsertFrontList(other *PersonList) (ok bool) {
	rotate := !this.Empty()
	if ok = this.InsertBackList(other); ok && rotate {
		this.RotateBackward()
	}
	return
}

func (this *PersonList) InsertBackList(other *PersonList) (ok bool) {
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

func (this *PersonList) PushBefore(v *Person, mark *PersonListNode) (n *PersonListNode) {
	if mark != nil {
		n = &PersonListNode{Person: v}
		n = this.InsertBefore(n, mark)
	}
	return
}

func (this *PersonList) PushAfter(v *Person, mark *PersonListNode) (n *PersonListNode) {
	if mark != nil {
		n = &PersonListNode{Person: v}
		n = this.InsertAfter(n, mark)
	}
	return
}

func (this *PersonList) InsertBefore(node, mark *PersonListNode) (n *PersonListNode) {
	if n = node; node != nil && mark != nil && node != mark {
		n.next = mark
		n.prev = mark.prev
		mark.prev.next = n
		mark.prev = n
		if this.head == mark {
			this.RotateBackward()
		}
	}
	return
}

func (this *PersonList) InsertAfter(node, mark *PersonListNode) (n *PersonListNode) {
	if n = node; node != nil && mark != nil {
		n.next = mark.next
		n.prev = mark
		mark.next = n
	}
	return
}

func (this *PersonList) RemoveFront() (n *PersonListNode) {
	return this.Remove(this.Front())
}

func (this *PersonList) RemoveBack() (n *PersonListNode) {
	return this.Remove(this.Back())
}

func (this *PersonList) Remove(node *PersonListNode) (n *PersonListNode) {
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

func (this *PersonList) Reachable(node, dest *PersonListNode) (ok bool) {
	if ok = (node == dest) && node != nil; !ok && node != nil && dest != nil {
		for p := node; p != nil && p != node; p = p.next {
			if ok = (p == dest); ok {
				break
			}
		}
	}
	return
}

func (this *PersonList) IsValidNode(node *PersonListNode) bool {
	return this.Reachable(this.head, node)
}

func (this *PersonList) MoveFront(node *PersonListNode) (r *PersonListNode) {
	if r = this.MoveBack(node); r != nil {
		this.RotateBackward()
	}
	return
}

func (this *PersonList) MoveBack(node *PersonListNode) (r *PersonListNode) {
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

func (this *PersonList) MoveBefore(node, mark *PersonListNode) (r *PersonListNode) {
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

func (this *PersonList) MoveAfter(node, mark *PersonListNode) (r *PersonListNode) {
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

func (this *PersonList) Empty() bool {
	return this.head == nil
}

func (this *PersonList) Reverse() {
	v := this.Visitor()
	this.Clear()
	for v.Next() {
		this.InsertFront(v.Get())
	}
}

func (this *PersonList) Sort() {
	this.mergeSort()
}

//STL's merge sort algorithm for list
func (this *PersonList) mergeSort() {
	if nil == this.head || this.head == this.head.next { //0 or 1 element, no need to sort
		return
	}

	var (
		hand       PersonList
		binList    [64]PersonList //save temp list that len=2^i
		nFilledBin = 0
	)

	for !this.Empty() {
		hand.InsertFront(this.RemoveFront())
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
func (this *PersonList) merge(other *PersonList) {
	if this.Empty() || other.Empty() {
		this.InsertBackList(other)
		return
	}

	p, po := this.Front(), other.Front()
	for p != nil && po != nil {
		if gPersonListGbl.cmp.F(po.Person, p.Person) {
			n := other.RemoveFront()
			po = other.Front()
			this.InsertBefore(n, p)
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
