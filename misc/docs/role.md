### DIY Problem

**Question**:
In computer systems security, role-based access control (RBAC) or role-based security is an approach to restricting system access to authorized users. It is used by the majority of enterprises with more than 500 employees, and can implement mandatory access control (MAC) or discretionary access control (DAC).

Within an organization, roles are created for various job functions. The permissions to perform certain operations are assigned to specific roles. Members or staff (or other system users) are assigned particular roles, and through those role assignments acquire the permissions needed to perform particular system functions. Since users are not assigned permissions directly, but only acquire them through their role (or roles), management of individual user rights becomes a matter of simply assigning appropriate roles to the user's account; this simplifies common operations, such as adding a user, or changing a user's department.

All roles have their own unique ID, and the relation of roles are defined using a pair of parent-child ID.
e.g. `{parent: 1, child: 2,}`


Design an algorithm which can check if a role is a child of one specific role recursively, with time complexity lower than O(n^2).


**Answer**:
```go
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

```

First we build a adjacency list, while it loops through all relations, we have time complexity = O(n)

Then we use the adjacency list in `IsChild` like following example:
```go
aList := misc.BuildAdjacencyList(relations)
isChild := misc.IsChild(root, child, aList)
```

In `IsChild`, we check if `child` is in `root`'s children, if not, we remove `root`'s children mappings from the adjacency list, just so if the role relations are not `DAG`, we can avoid looping through `root`'s children again.

And because we only loop through each mapping once, the worst case running time is O(n), in which we loop through every key and every relation in each mapping.
