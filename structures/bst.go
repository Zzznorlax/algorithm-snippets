package structures

import "fmt"

type BSTNode struct {
	Left   *BSTNode
	Right  *BSTNode
	Parent *BSTNode
	Key    int
}

type BST struct {
	Root *BSTNode
}

func (node *BSTNode) insert(key int) {

	if key > node.Key {
		if node.Right != nil {
			node.Right.insert(key)
		} else {
			node.Right = &BSTNode{Key: key, Parent: node}
		}
	} else {
		if node.Left != nil {
			node.Left.insert(key)
		} else {
			node.Left = &BSTNode{Key: key, Parent: node}
		}
	}
}

func (node *BSTNode) search(key int) *BSTNode {

	if key > node.Key {
		if node.Right == nil {
			return nil
		} else if node.Right.Key == key {
			return node.Right
		} else {
			return node.Right.search(key)
		}
	} else {
		if node.Left == nil {
			return nil
		} else if node.Left.Key == key {
			return node.Left
		} else {
			return node.Left.search(key)
		}
	}
}

func (node *BSTNode) minimum() *BSTNode {

	min := node

	for min.Left != nil {
		min = min.Left
	}

	return min
}

func (node *BSTNode) successor() *BSTNode {
	successor := node

	if successor.Right != nil {
		successor = successor.Right.minimum()
	}

	return successor
}

func (node *BSTNode) inorderTraversal() {
	if node == nil {
		return
	}

	if node.Left != nil {
		node.Left.inorderTraversal()
	}

	fmt.Printf("%v\n", node.Key)

	if node.Right != nil {
		node.Right.inorderTraversal()
	}
}

func (t *BST) Delete(key int) *BSTNode {
	node := t.Search(key)

	if node == nil {
		return node
	}

	if node.Left == nil && node.Right == nil {

		if node.Parent.Right == node {
			node.Parent.Right = nil
		} else {
			node.Parent.Left = nil
		}

	} else if node.Right != nil && node.Left == nil {
		node.Right.Parent = node.Parent

		if node.Parent.Right == node {
			node.Parent.Right = node.Right
		} else {
			node.Parent.Left = node.Right
		}

	} else if node.Left != nil && node.Right == nil {
		node.Left.Parent = node.Parent

		if node.Parent.Right == node {
			node.Parent.Right = node.Left
		} else {
			node.Parent.Left = node.Left
		}

	} else {

		successor := node.successor()

		successor.Parent.Left = successor.Right
		if successor.Right != nil {
			successor.Right.Parent = successor.Parent
		}

		successor.Left = node.Left
		if node.Left != nil {
			node.Left.Parent = successor
		}

		successor.Parent = node.Parent
		if node.Parent.Right == node {
			node.Parent.Right = successor
		} else {
			node.Parent.Left = successor
		}

		successor.Right = node.Right
		if node.Right != nil {
			node.Right.Parent = successor
		}

	}

	node.Parent = nil
	node.Left = nil
	node.Right = nil

	return node
}

func (t *BST) Insert(key int) {

	if t.Root == nil {
		t.Root = &BSTNode{Key: key}

	} else {
		t.Root.insert(key)
	}

}

func (t *BST) Search(key int) *BSTNode {
	return t.Root.search(key)
}

func (t *BST) InorderTraversal() {
	t.Root.inorderTraversal()
}
