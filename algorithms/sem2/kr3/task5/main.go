package main

import "fmt"

type Color bool

const (
	Red   Color = true
	Black Color = false
)

type Node struct {
	key    int
	color  Color
	left   *Node
	right  *Node
	parent *Node
}

type RBTree struct {
	root *Node
	nil_ *Node
}

func NewRBTree() *RBTree {
	sentinel := &Node{color: Black}
	return &RBTree{root: sentinel, nil_: sentinel}
}

func (t *RBTree) rotateLeft(x *Node) {
	y := x.right
	x.right = y.left
	if y.left != t.nil_ {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == t.nil_ {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *RBTree) rotateRight(x *Node) {
	y := x.left
	x.left = y.right
	if y.right != t.nil_ {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == t.nil_ {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

func (t *RBTree) insertFixup(z *Node) {
	for z.parent.color == Red {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y.color == Red {
				z.parent.color = Black
				y.color = Black
				z.parent.parent.color = Red
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.rotateLeft(z)
				}
				z.parent.color = Black
				z.parent.parent.color = Red
				t.rotateRight(z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if y.color == Red {
				z.parent.color = Black
				y.color = Black
				z.parent.parent.color = Red
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rotateRight(z)
				}
				z.parent.color = Black
				z.parent.parent.color = Red
				t.rotateLeft(z.parent.parent)
			}
		}
	}
	t.root.color = Black
}

func (t *RBTree) Insert(key int) {
	z := &Node{key: key, color: Red, left: t.nil_, right: t.nil_, parent: t.nil_}
	y := t.nil_
	x := t.root
	for x != t.nil_ {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == t.nil_ {
		t.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
	t.insertFixup(z)
}

func (t *RBTree) transplant(u, v *Node) {
	if u.parent == t.nil_ {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

func (t *RBTree) minimum(x *Node) *Node {
	for x.left != t.nil_ {
		x = x.left
	}
	return x
}

func (t *RBTree) deleteFixup(x *Node) {
	for x != t.root && x.color == Black {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == Red {
				w.color = Black
				x.parent.color = Red
				t.rotateLeft(x.parent)
				w = x.parent.right
			}
			if w.left.color == Black && w.right.color == Black {
				w.color = Red
				x = x.parent
			} else {
				if w.right.color == Black {
					w.left.color = Black
					w.color = Red
					t.rotateRight(w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = Black
				w.right.color = Black
				t.rotateLeft(x.parent)
				x = t.root
			}
		} else {
			w := x.parent.left
			if w.color == Red {
				w.color = Black
				x.parent.color = Red
				t.rotateRight(x.parent)
				w = x.parent.left
			}
			if w.right.color == Black && w.left.color == Black {
				w.color = Red
				x = x.parent
			} else {
				if w.left.color == Black {
					w.right.color = Black
					w.color = Red
					t.rotateLeft(w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = Black
				w.left.color = Black
				t.rotateRight(x.parent)
				x = t.root
			}
		}
	}
	x.color = Black
}

func (t *RBTree) Delete(key int) {
	z := t.root
	for z != t.nil_ {
		if key == z.key {
			break
		} else if key < z.key {
			z = z.left
		} else {
			z = z.right
		}
	}
	if z == t.nil_ {
		return
	}
	y := z
	yOrigColor := y.color
	var x *Node
	if z.left == t.nil_ {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == t.nil_ {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y = t.minimum(z.right)
		yOrigColor = y.color
		x = y.right
		if y.parent == z {
			x.parent = y
		} else {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}
	if yOrigColor == Black {
		t.deleteFixup(x)
	}
}

func colorStr(c Color) string {
	if c == Red {
		return "R"
	}
	return "B"
}

func (t *RBTree) inorder(n *Node) {
	if n == t.nil_ {
		return
	}
	t.inorder(n.left)
	fmt.Printf("key=%d(%s) ", n.key, colorStr(n.color))
	t.inorder(n.right)
}

func (t *RBTree) Print() {
	t.inorder(t.root)
	fmt.Println()
}

func main() {
	tree := NewRBTree()
	for _, v := range []int{10, 20, 30, 15, 5, 25, 1} {
		tree.Insert(v)
	}
	fmt.Print("После вставок: ")
	tree.Print()

	tree.Delete(20)
	fmt.Print("После удаления 20: ")
	tree.Print()

	tree.Delete(5)
	fmt.Print("После удаления 5: ")
	tree.Print()
}
