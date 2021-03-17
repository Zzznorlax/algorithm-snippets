### HW 1 Problem 9: Skyline

Given n rectangular buildings in a 2-dimensional city,
computes the skyline of these buildings, eliminating hidden lines.
All buildings share common bottom (absolutely flat surface at height 0),
and every building Bi (1<=i<=n) is represented by triplet (Li, Hi, Ri),
where Li and Hi denote the x coordinates of the left and right of the ith building, and Hi denotes its height.
A skyline of a set of n buildings is a list of x coordinates and the heights connecting them arranged in order from left to right.

Example: The skyline of the buildings
{(3, 13, 9), (1, 11, 5), (12, 7, 16), (14, 3, 25), (19, 18, 22), (2, 6, 7), (23, 13, 29), (23, 4, 28)} is
{(1, 11), (3, 13), (9, 0), (12, 7), (16, 3), (19, 18), (22, 3), (23, 13), (29, 0)}.

Note that there must be no consecutive horizontal lines of equal height in the output skyline.

**(a)** Let the size of a skyline be the total number of elements (coordinates and heights) in its list.
Design an algorithm for combining a skyline A of size n1 and a skyline B of size n2 into one skyline S of size O(n1+n2).
Your algorithm should run in O(n1+n2) worst-case time. Justify the running time of your algorithm.

In the following code, time complexities are specified for each operation,
see comments. So the overall worst-case time complexity for `MergeSkylines` is O(n1 + n2)

**Algorithm**:
```go
// Point represents a point in skyline
type Point struct {
	Pos    int
	Height int
}

// MergeSkylines implements the merging two skylines into one algorithm.
// Time complexity O(n1 + n2).
func MergeSkylines(a []Point, b []Point) []Point {

	indexA := 0
	indexB := 0

	var mergedSkyline []Point

	// Loops through points in the lists of points.
	// O(n1 + n2)
	for indexA < len(a) && indexB < len(b) {

		var currentPoint Point
		var refPointStart Point
		var refPointEnd Point

		// Finds the leftmost point and points at its both sides on the other skyline.
		if a[indexA].Pos < b[indexB].Pos {
			currentPoint = a[indexA]

			refPointEnd = b[indexB]
			refPointStart = b[indexB]
			if indexB > 0 {
				refPointStart = b[indexB-1]
			}

			indexA += 1

		} else {
			currentPoint = b[indexB]

			refPointEnd = a[indexA]
			refPointStart = a[indexA]
			if indexA > 0 {
				refPointStart = a[indexA-1]
			}

			indexB += 1
		}

		// Checks and updates the height of the point.
		// O(1)
		currentPoint = overwritePointHeight(currentPoint, refPointStart, refPointEnd)

		mergedSkyline = append(mergedSkyline, currentPoint)

	}

	// Appends the rest of the points to the merged skyline.
	// In either condition time complexity = O(n1 + n2)
	if indexA == len(a) {
		// O(n1 + n2)
		for i := indexB; i < len(b); i++ {
			mergedSkyline = append(mergedSkyline, b[i])
		}
	} else {
		// O(n1 + n2)
		for i := indexA; i < len(a); i++ {
			mergedSkyline = append(mergedSkyline, a[i])
		}
	}

	// O(n1 + n2)
	mergedSkyline = trimSkyline(mergedSkyline)

	return mergedSkyline
}

// trimSkyline removes redundant points from a merged skyline.
// Time complexity = O(n)
func trimSkyline(skyline []Point) []Point {
	var processedSkyline []Point

	for i := 0; i < len(skyline); i++ {
		// Skips if two consecutive points' height are the same.
		if i >= 1 && skyline[i].Height == skyline[i-1].Height {
			continue
		}
		processedSkyline = append(processedSkyline, skyline[i])
	}

	return processedSkyline
}

// overwritePointHeight compares the current point with ref points, overwrites the height of currentPoint if needed.
// Time complexity O(1)
func overwritePointHeight(currentPoint Point, refPointStart Point, refPointEnd Point) Point {
	if refPointStart.Pos <= currentPoint.Pos && currentPoint.Pos <= refPointEnd.Pos {

		if refPointStart.Height > currentPoint.Height {
			currentPoint.Height = refPointStart.Height
		}
	}

	return currentPoint
}
```

**(b)** Design an O(n log n) algorithm for finding the skyline of n buildings.
Any algorithm that requires O(n^2) running time will not receive any credit. Justify the running time of your algorithm.

Time complexities are specified in comments for each operation, so the overall time complexity for `ToSkyline` is O(n log n).

**Algorithm**:
```go
// Building represents a building with attributes left pos, right pos and height
type Building struct {
	Left   int
	Right  int
	Height int
}


// ToSkyline implements the algorithm which convert a list of buildings to a skyline using divide and conquer
// in order to achieve time complexity = O(n log n). In the conquer part, we use the MergeSkylines from (a).
// Time complexity = O(n log n)
func ToSkyline(buildings []Building, start int, end int) []Point {

	// O(1)
	if start == end {
		building := buildings[start]
		return []Point{
			{
				Pos:    building.Left,
				Height: building.Height,
			},
			{
				Pos:    building.Right,
				Height: 0,
			},
		}
	}

	// Divides the list of building.
	mid := (start + end) / 2

	// O(n log n), while n is half of the original size.
	leftSkyline := ToSkyline(buildings, start, mid)

	// O(n log n), while n is half of the original size.
	rightSkyLine := ToSkyline(buildings, mid+1, end)

	// O(n)
	return MergeSkylines(leftSkyline, rightSkyLine)
}
```

**(c)** Briefly justify the correctness of part (a) and part (b). [Hint: Loop invariant; mathematical induction.]

**(a):**
**Loop invariant**:
pos of points in the merged skyline are always sorted from left to right.

**Initialization**:
Merged skyline has no point

**Maintenance**:
Because we build skyline start with left point of the building,
and also merge them start with the leftmost point. So points in the merged skyline are sorted in the same order as initialization.

**Termination**:
When points in either skyline are in the merged skyline, we append the rest of the points in the other skyline to the merged skyline from the leftmost to the rightmost, as what we've done in maintenance. So points in the merged skyline remain the same order, from the leftmost to rightmost.

**(b):**
**Loop invariant**:
pos of points in the merged skyline are always sorted from left to right.

**Initialization**:
Merged skyline has no point

**Maintenance**:
Because we build skyline start with left point of the building then the right point, so when len(building) = 1,
skyline is sorted.
When len(building) > 1, merging start with the leftmost point. So points in the merged skyline are sorted in the same order as initialization.

**Termination**:
At termination, the skyline is merged from all buildings' skylines. From maintenance, all sub-skyline are sorted, and all merged skylines are sorted too, so the final skyline is sorted.
