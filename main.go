package main

import (
	"algorithms/misc"
	"fmt"
)

func main() {
	r := []misc.RoleRelation{
		{
			Parent: 1,
			Child:  2,
		},
		{
			Parent: 1,
			Child:  3,
		},
		{
			Parent: 2,
			Child:  3,
		},
	}

	aList := misc.BuildAdjacencyList(r)

	isChild := misc.IsChild(3, 2, aList)

	fmt.Printf("Is child %v", isChild)
}
