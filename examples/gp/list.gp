//#GOGP_IGNORE_BEGIN
///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/vipally/gogp]
// Last update at: [Tue Nov 22 2016 15:09:31]
// Generate from:
//   [github.com/vipally/gogp/examples/gp/list.gp.go]
//   [github.com/vipally/gogp/examples/gp/gp.gpg] [GOGP_REVERSE_list]
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
//#GOGP_IGNORE_END

<PACKAGE>

//#GOGP_REQUIRE(github.com/vipally/gogp/lib/fakedef,_)

//#GOGP_REQUIRE(github.com/vipally/gogp/examples/gp/functorcmp)

////////////////////////////////////////////////////////////////////////////////

var g<GLOBAL_NAME_PREFIX>ListGbl struct {
	cmp Cmp<GLOBAL_NAME_PREFIX>
}

func init() {
	g<GLOBAL_NAME_PREFIX>ListGbl.cmp = g<GLOBAL_NAME_PREFIX>ListGbl.cmp.CreateByName("#GOGP_GPGCFG(GOGP_DefaultCmpType)")
}

//double-way cycle link list node
type <GLOBAL_NAME_PREFIX>ListNode struct {
	<VALUE_TYPE>
	prev, next *<GLOBAL_NAME_PREFIX>ListNode
}

func (this *<GLOBAL_NAME_PREFIX>ListNode) Get() <VALUE_TYPE> {
	return this.<VALUE_TYPE>
}

func (this *<GLOBAL_NAME_PREFIX>ListNode) Set(v <VALUE_TYPE>) (old <VALUE_TYPE>) {
	old, this.<VALUE_TYPE> = this.<VALUE_TYPE>, v
	return
}

func (this *<GLOBAL_NAME_PREFIX>ListNode) Next() (r *<GLOBAL_NAME_PREFIX>ListNode) {
	if this != nil {
		r = this.next
	}
	return
}

func (this *<GLOBAL_NAME_PREFIX>ListNode) Prev() (r *<GLOBAL_NAME_PREFIX>ListNode) {
	if this != nil {
		r = this.prev
	}
	return
}

type GOGGlobalNamePrefixListNodeVisitor struct {
	node, head *<GLOBAL_NAME_PREFIX>ListNode
}

func (this *GOGGlobalNamePrefixListNodeVisitor) Reset() {
	this.node = nil
}

func (this *GOGGlobalNamePrefixListNodeVisitor) Next() (ok bool) {
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

func (this *GOGGlobalNamePrefixListNodeVisitor) Prev() (ok bool) {
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

func (this *GOGGlobalNamePrefixListNodeVisitor) Get() *<GLOBAL_NAME_PREFIX>ListNode {
	return this.node
}

func (this *<GLOBAL_NAME_PREFIX>List) Visitor() *GOGGlobalNamePrefixListNodeVisitor {
	n := &GOGGlobalNamePrefixListNodeVisitor{node: nil, head: this.head}
	return n
}

//list object
type <GLOBAL_NAME_PREFIX>List struct {
	head *<GLOBAL_NAME_PREFIX>ListNode
}

//new object
func New<GLOBAL_NAME_PREFIX>List() *<GLOBAL_NAME_PREFIX>List {
	return &<GLOBAL_NAME_PREFIX>List{}
}

//func (this *<GLOBAL_NAME_PREFIX>List) Len() int {
//	return 0
//}

func (this *<GLOBAL_NAME_PREFIX>List) Front() *<GLOBAL_NAME_PREFIX>ListNode {
	return this.head
}
func (this *<GLOBAL_NAME_PREFIX>List) Back() (r *<GLOBAL_NAME_PREFIX>ListNode) {
	if this.head != nil {
		r = this.head.prev
	}
	return
}

func (this *<GLOBAL_NAME_PREFIX>List) Clear() {
	this.head = nil
}

func (this *<GLOBAL_NAME_PREFIX>List) RotateForward() {
	if this.head != nil {
		this.head = this.head.next
	}
}
func (this *<GLOBAL_NAME_PREFIX>List) RotateBackward() {
	if this.head != nil {
		this.head = this.head.prev
	}
}

func (this *<GLOBAL_NAME_PREFIX>List) PushFront(v <VALUE_TYPE>) *<GLOBAL_NAME_PREFIX>ListNode {
	n := &<GLOBAL_NAME_PREFIX>ListNode{<VALUE_TYPE>: v}
	return this.InsertFront(n)
}

func (this *<GLOBAL_NAME_PREFIX>List) PushBack(v <VALUE_TYPE>) *<GLOBAL_NAME_PREFIX>ListNode {
	n := &<GLOBAL_NAME_PREFIX>ListNode{<VALUE_TYPE>: v}
	return this.InsertBack(n)
}

func (this *<GLOBAL_NAME_PREFIX>List) InsertFront(node *<GLOBAL_NAME_PREFIX>ListNode) (n *<GLOBAL_NAME_PREFIX>ListNode) {
	if n = this.InsertBack(node); n != nil {
		this.RotateBackward()
	}
	return
}

func (this *<GLOBAL_NAME_PREFIX>List) InsertBack(node *<GLOBAL_NAME_PREFIX>ListNode) (n *<GLOBAL_NAME_PREFIX>ListNode) {
	if n = node; n != nil {
		if this.head != nil {
			this.head, n.next, n.prev = n, n, n
		} else {
			n.next = this.head
			n.prev = this.head.prev
		}
	}
	return
}

func (this *<GLOBAL_NAME_PREFIX>List) PopFront() (v <VALUE_TYPE>, ok bool) {
	if n := this.Remove(this.Front()); n != nil {
		v, ok = n.Get(), true
	}
	return
}

func (this *<GLOBAL_NAME_PREFIX>List) PopBack() (v <VALUE_TYPE>, ok bool) {
	if n := this.Remove(this.Back()); n != nil {
		v, ok = n.Get(), true
	}
	return
}

func (this *<GLOBAL_NAME_PREFIX>List) InsertFrontList(other *<GLOBAL_NAME_PREFIX>List) (ok bool) {
	rotate := !this.Empty()
	if ok = this.InsertBackList(other); ok && rotate {
		this.RotateBackward()
	}
	return
}

func (this *<GLOBAL_NAME_PREFIX>List) InsertBackList(other *<GLOBAL_NAME_PREFIX>List) (ok bool) {
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

func (this *<GLOBAL_NAME_PREFIX>List) PushBefore(v <VALUE_TYPE>, mark *<GLOBAL_NAME_PREFIX>ListNode) (n *<GLOBAL_NAME_PREFIX>ListNode) {
	if mark != nil {
		n = &<GLOBAL_NAME_PREFIX>ListNode{<VALUE_TYPE>: v}
		n = this.InsertBefore(n, mark)
	}
	return
}

func (this *<GLOBAL_NAME_PREFIX>List) PushAfter(v <VALUE_TYPE>, mark *<GLOBAL_NAME_PREFIX>ListNode) (n *<GLOBAL_NAME_PREFIX>ListNode) {
	if mark != nil {
		n = &<GLOBAL_NAME_PREFIX>ListNode{<VALUE_TYPE>: v}
		n = this.InsertAfter(n, mark)
	}
	return
}

func (this *<GLOBAL_NAME_PREFIX>List) InsertBefore(node, mark *<GLOBAL_NAME_PREFIX>ListNode) (n *<GLOBAL_NAME_PREFIX>ListNode) {
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

func (this *<GLOBAL_NAME_PREFIX>List) InsertAfter(node, mark *<GLOBAL_NAME_PREFIX>ListNode) (n *<GLOBAL_NAME_PREFIX>ListNode) {
	if n = node; node != nil && mark != nil {
		n.next = mark.next
		n.prev = mark
		mark.next = n
	}
	return
}

func (this *<GLOBAL_NAME_PREFIX>List) RemoveFront() (n *<GLOBAL_NAME_PREFIX>ListNode) {
	return this.Remove(this.Front())
}

func (this *<GLOBAL_NAME_PREFIX>List) RemoveBack() (n *<GLOBAL_NAME_PREFIX>ListNode) {
	return this.Remove(this.Back())
}

func (this *<GLOBAL_NAME_PREFIX>List) Remove(node *<GLOBAL_NAME_PREFIX>ListNode) (n *<GLOBAL_NAME_PREFIX>ListNode) {
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

func (this *<GLOBAL_NAME_PREFIX>List) Reachable(node, dest *<GLOBAL_NAME_PREFIX>ListNode) (ok bool) {
	if ok = (node == dest) && node != nil; !ok && node != nil && dest != nil {
		for p := node; p != nil && p != node; p = p.next {
			if ok = (p == dest); ok {
				break
			}
		}
	}
	return
}

func (this *<GLOBAL_NAME_PREFIX>List) IsValidNode(node *<GLOBAL_NAME_PREFIX>ListNode) bool {
	return this.Reachable(this.head, node)
}

func (this *<GLOBAL_NAME_PREFIX>List) MoveFront(node *<GLOBAL_NAME_PREFIX>ListNode) (r *<GLOBAL_NAME_PREFIX>ListNode) {
	if r = this.MoveBack(node); r != nil {
		this.RotateBackward()
	}
	return
}

func (this *<GLOBAL_NAME_PREFIX>List) MoveBack(node *<GLOBAL_NAME_PREFIX>ListNode) (r *<GLOBAL_NAME_PREFIX>ListNode) {
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

func (this *<GLOBAL_NAME_PREFIX>List) MoveBefore(node, mark *<GLOBAL_NAME_PREFIX>ListNode) (r *<GLOBAL_NAME_PREFIX>ListNode) {
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

func (this *<GLOBAL_NAME_PREFIX>List) MoveAfter(node, mark *<GLOBAL_NAME_PREFIX>ListNode) (r *<GLOBAL_NAME_PREFIX>ListNode) {
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

func (this *<GLOBAL_NAME_PREFIX>List) Empty() bool {
	return this.head == nil
}

func (this *<GLOBAL_NAME_PREFIX>List) Reverse() {
	v := this.Visitor()
	this.Clear()
	for v.Next() {
		this.InsertFront(v.Get())
	}
}

func (this *<GLOBAL_NAME_PREFIX>List) Sort() {
	this.mergeSort()
}

//STL's merge sort algorithm for list
func (this *<GLOBAL_NAME_PREFIX>List) mergeSort() {
	if nil == this.head || this.head == this.head.next { //0 or 1 element, no need to sort
		return
	}

	var (
		hand, unsorted <GLOBAL_NAME_PREFIX>List
		binList        [64]<GLOBAL_NAME_PREFIX>List //save temp list that len=2^i
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
func (this *<GLOBAL_NAME_PREFIX>List) merge(other *<GLOBAL_NAME_PREFIX>List) {
	if this.Empty() || other.Empty() {
		this.InsertBackList(other)
		return
	}

	p, po := this.Front(), other.Front()
	for p != nil && po != nil {
		if g<GLOBAL_NAME_PREFIX>ListGbl.cmp.F(po.<VALUE_TYPE>, p.<VALUE_TYPE>) {
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

