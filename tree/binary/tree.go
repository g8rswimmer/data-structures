package binary

import (
	"errors"
)

type Tree struct {
	root *node
}

func New(root Comparor) *Tree {
	return &Tree{
		root: &node{
			data: root,
		},
	}
}

func (t *Tree) Insert(obj Comparor) error {
	if t.root == nil {
		return errors.New("tree: must have a root node")
	}
	n := t.root
	for {
		switch n.data.Compare(obj) {
		case Left, Equal:
			if n.left == nil {
				n.left = &node{
					data: obj,
				}
				return nil
			}
			n = n.left
		case Right:
			if n.right == nil {
				n.right = &node{
					data: obj,
				}
				return nil
			}
			n = n.right
		default:
			return errors.New("tree: compare must return left or right")
		}
	}
}

func (t *Tree) Delete(obj Comparor) error {
	parent, child := t.get(obj)
	if child == nil {
		return errors.New("tree: unable to find node to delete")
	}
	switch {
	case child.left == nil && child.right == nil:
		t.deleteLeaf(parent, child)
		return nil
	case child.left == nil:
		t.replace(parent, child, child.right)
		return nil
	case child.right == nil:
		t.replace(parent, child, child.left)
		return nil
	}
	return errors.New("something")
}

func (t *Tree) replace(parent, child, replace *node) {
	switch {
	case parent.left == child.left:
		parent.left = replace
	case parent.right == child.right:
		parent.right = replace
	}
}

func (t *Tree) deleteLeaf(parent, child *node) {
	switch {
	case parent.left == child.left:
		parent.left = nil
	case parent.right == child.right:
		parent.right = nil
	}
}

func (t *Tree) Has(obj Comparor) bool {

	if _, n := t.get(obj); n != nil {
		return true
	}
	return false
}

func (t *Tree) get(obj Comparor) (parent *node, child *node) {
	parent = t.root
	child = t.root
	for child != nil {
		switch child.data.Compare(obj) {
		case Left:
			parent = child
			child = child.left
		case Right:
			parent = child
			child = child.right
		case Equal:
			return
		default:
			child = nil
			return
		}
	}
	return
}

func (t *Tree) Depth(path Path) int {
	switch path {
	case Left:
		return t.depth(t.root.left)
	case Right:
		return t.depth(t.root.right)
	default:
		return t.depth(t.root)
	}
}

func (t *Tree) depth(n *node) int {
	if n == nil {
		return 0
	}
	l := t.depth(n.left)
	r := t.depth(n.right)

	if l > r {
		return l + 1
	}
	return r + 1
}
func (t *Tree) Inorder() []interface{} {
	objs := []interface{}{}
	return t.inorder(t.root, objs)
}

func (t *Tree) inorder(n *node, objs []interface{}) []interface{} {
	if n == nil {
		return objs
	}
	objs = t.inorder(n.left, objs)
	objs = append(objs, n.data)
	objs = t.inorder(n.right, objs)
	return objs
}

func (t *Tree) Preorder() []interface{} {
	objs := []interface{}{}
	return t.preorder(t.root, objs)
}

func (t *Tree) preorder(n *node, objs []interface{}) []interface{} {
	if n == nil {
		return objs
	}
	objs = append(objs, n.data)
	objs = t.preorder(n.left, objs)
	objs = t.preorder(n.right, objs)
	return objs
}

func (t *Tree) Postorder() []interface{} {
	objs := []interface{}{}
	return t.postorder(t.root, objs)
}

func (t *Tree) postorder(n *node, objs []interface{}) []interface{} {
	if n == nil {
		return objs
	}
	objs = t.postorder(n.left, objs)
	objs = t.postorder(n.right, objs)
	objs = append(objs, n.data)
	return objs
}
