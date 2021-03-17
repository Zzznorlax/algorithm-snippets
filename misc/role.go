package misc

type RoleRelation struct {
	Parent int
	Child  int
}

// O(n)
func BuildAdjacencyList(relations []RoleRelation) map[int][]int {

	aList := make(map[int][]int)

	for _, r := range relations {
		aList[r.Parent] = append(aList[r.Parent], r.Child)
	}

	return aList
}

// O(n)
func IsChild(root int, child int, aList map[int][]int) bool {

	// O(1)
	nextChildren, ok := aList[root]

	if ok {
		for _, c := range nextChildren {

			if child == c {
				return true

			}
		}

		// Deletes the root's mapping so it'll be skipped next time.
		delete(aList, root)

		// Recursively checks if target role is child of root's children
		for _, c := range nextChildren {
			if IsChild(c, child, aList) {
				return true
			}
		}

	}

	return false
}
