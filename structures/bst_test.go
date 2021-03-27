package structures_test

import (
	"algorithms/structures"
	"testing"
)

func initBST() *structures.BST {

	bst := structures.BST{}

	bst.Insert(5)
	bst.Insert(6)
	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(7)
	bst.Insert(9)
	bst.Insert(1)
	bst.Insert(8)
	bst.Insert(0)

	// bst.InorderTraversal()

	return &bst
}

func TestInsert(t *testing.T) {
	bst := initBST()

	key := 17

	bst.Insert(key)

	if bst.Search(key) == nil {
		t.Errorf("Key %v not inserted", key)
	}
}

func TestDelete(t *testing.T) {
	bst := initBST()

	key := 17

	bst.Insert(key)

	bst.Delete(key)

	if bst.Search(key) != nil {
		t.Errorf("Key %v not deleted", key)
	}
}
