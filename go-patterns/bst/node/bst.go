package bst

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	height int
}

func (n *Node) Heigth() int {
	if n == nil {
		return 0
	}
	h := max(n.Right.Heigth(), n.Left.Heigth()) + 1
	return h
}

func rotateLeft(c *Node) *Node {
	root := c.Right
	if c.Right.Left != nil {
		root = c.Right.Left
		c.Right.Left = nil
		root.Right = c.Right
	}
	c.Right = nil
	tmp := c
	root.Left = tmp
	return root
}

func (n *Node) Balance() int {
	return n.Right.Heigth() - n.Left.Heigth()
}

func (n *Node) IsBalanced() bool {
	if n.Balance() > 1 || n.Balance() < -1 {
		return false
	}
	return true
}

func (t *Tree) Insert(node *Node) {
	t.Node = insert(t.Node, node) //insert
}

func insert(root *Node, node *Node) *Node {
	if root.Value > node.Value {
		root.Left = insertIfNull(root.Left, node)
	} else if root.Value < node.Value {
		root.Right = insertIfNull(root.Right, node)
	}
	if root.Balance() > 1 {
		root = rotateLeft(root)
	}
	return root
	//if equal do nothing
}

func insertIfNull(root *Node, node *Node) *Node {
	if root == nil {
		return node
	} else {
		return insert(root, node)
	}
}

type Tree struct {
	Node *Node
}

func (t *Tree) Add(value int) {
	n := NewNode(value)
	if t.Node == nil {
		t.Node = NewNode(value)
		return
	}
	t.Insert(n)
}

func NewTree() *Tree {
	return &Tree{}
}
