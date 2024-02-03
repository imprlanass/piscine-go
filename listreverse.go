package piscine

func ListReverse(l *List) {
	curr := l.Head
	var previous, next *NodeL

	for curr != nil {
		next = curr.Next
		curr.Next = previous
		previous = curr
		curr = next
	}

	l.Tail = l.Head
	l.Head = previous
}
