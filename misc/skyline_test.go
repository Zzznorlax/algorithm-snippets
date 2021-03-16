package misc_test

import (
	"algorithms/misc"
	"testing"
)

func TestMergeSkylines(t *testing.T) {

	var a = []misc.Point{
		{
			Pos:    1,
			Height: 2,
		},
		{
			Pos:    2,
			Height: 0,
		},
		{
			Pos:    4,
			Height: 7,
		},
		{
			Pos:    6,
			Height: 2,
		},
		{
			Pos:    7,
			Height: 0,
		},
		{
			Pos:    9,
			Height: 5,
		},
		{
			Pos:    10,
			Height: 0,
		},
	}

	var b = []misc.Point{
		{
			Pos:    2,
			Height: 7,
		},
		{
			Pos:    3,
			Height: 0,
		},
		{
			Pos:    4,
			Height: 2,
		},
		{
			Pos:    5,
			Height: 0,
		},
		{
			Pos:    6,
			Height: 9,
		},
		{
			Pos:    8,
			Height: 0,
		},
		{
			Pos:    10,
			Height: 0,
		},
	}

	merged := misc.MergeSkylines(a, b)
	t.Logf("Result: %v", merged)
}

func TestToSkyline(t *testing.T) {
	var buildings = []misc.Building{
		{
			Left:   3,
			Right:  9,
			Height: 13,
		},
		{
			Left:   1,
			Right:  5,
			Height: 11,
		},
		{
			Left:   12,
			Right:  16,
			Height: 7,
		},
		{
			Left:   14,
			Right:  25,
			Height: 3,
		},
		{
			Left:   19,
			Right:  22,
			Height: 18,
		},
		{
			Left:   2,
			Right:  7,
			Height: 6,
		},
		{
			Left:   23,
			Right:  29,
			Height: 13,
		},
		{
			Left:   23,
			Right:  28,
			Height: 4,
		},
		{
			Left:   14,
			Right:  25,
			Height: 3,
		},
		{
			Left:   1,
			Right:  2,
			Height: 3,
		},
	}

	skyline := misc.ToSkyline(buildings, 0, len(buildings)-1)

	t.Logf("Result: %v", skyline)
}
