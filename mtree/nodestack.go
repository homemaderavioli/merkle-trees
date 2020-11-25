package mtree

type nodeStack []*Node

func (s nodeStack) Empty() bool  { return len(s) == 0 }
func (s nodeStack) Peek() *Node  { return s[len(s)-1] }
func (s *nodeStack) Put(i *Node) { (*s) = append((*s), i) }
func (s *nodeStack) Pop() *Node {
	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}
