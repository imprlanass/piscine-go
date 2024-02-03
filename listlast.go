package piscine

func ListLast(l *List) interface{} {
	if l == nil || l.Head == nil {
		return nil
	}

	return l.Tail.Data
}
