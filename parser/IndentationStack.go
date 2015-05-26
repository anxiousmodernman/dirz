package parser

type indentationStack struct {
	top  *StackItem
	size int
}

type StackItem struct {
	value interface{} // All types satisfy empty interface
	next  *StackItem
}

func (s *indentationStack) Len() int {
	return s.size
}

func (s *indentationStack) Push(value interface{}) {
	s.top = &StackItem{value, s.top}
	s.size++
}

func (s *indentationStack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}
