package misc

type Point struct {
	Pos    int
	Height int
}

type Building struct {
	Left   int
	Right  int
	Height int
}

func MergeSkylines(a []Point, b []Point) []Point {

	indexA := 0
	indexB := 0

	var mergedSkyline []Point

	for indexA < len(a) && indexB < len(b) {

		var currentPoint Point
		var refPointStart Point
		var refPointEnd Point

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

		currentPoint = overwritePointHeight(currentPoint, refPointStart, refPointEnd)

		mergedSkyline = append(mergedSkyline, currentPoint)

	}

	if indexA == len(a) {
		for i := indexB; i < len(b); i++ {
			mergedSkyline = append(mergedSkyline, b[i])
		}
	} else {
		for i := indexA; i < len(a); i++ {
			mergedSkyline = append(mergedSkyline, a[i])
		}
	}

	mergedSkyline = trimSkyline(mergedSkyline)

	return mergedSkyline
}

func ToSkyline(buildings []Building, start int, end int) []Point {

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

	mid := (start + end) / 2

	leftSkyline := ToSkyline(buildings, start, mid)
	rightSkyLine := ToSkyline(buildings, mid+1, end)

	return MergeSkylines(leftSkyline, rightSkyLine)
}

func trimSkyline(skyline []Point) []Point {
	var processedSkyline []Point

	for i := 0; i < len(skyline); i++ {
		if i >= 1 && skyline[i].Height == skyline[i-1].Height {
			continue
		}
		processedSkyline = append(processedSkyline, skyline[i])
	}

	return processedSkyline
}

func overwritePointHeight(currentPoint Point, refPointStart Point, refPointEnd Point) Point {
	if refPointStart.Pos <= currentPoint.Pos && currentPoint.Pos <= refPointEnd.Pos {

		if refPointStart.Height > currentPoint.Height {
			currentPoint.Height = refPointStart.Height
		}
	}

	return currentPoint
}
