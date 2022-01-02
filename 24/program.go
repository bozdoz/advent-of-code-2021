package main

var blockDiffs = [14][3]int{
	{1, 13, 0},
	{1, 11, 3},
	{1, 14, 8},
	{26, -5, 5},
	{1, 14, 13},
	{1, 10, 9},
	{1, 12, 6},
	{26, -14, 1},
	{26, -8, 1},
	{1, 13, 2},
	{26, 0, 7},
	{26, -5, 5},
	{26, -9, 8},
	{26, -1, 15},
}

func block(w, z, zdiv, xdiff, ydiff int) int {
	var x, y int
	x = z
	x %= 26
	z /= zdiv
	x += xdiff
	if x == w {
		x = 0
	} else {
		x = 1
	}
	y = 25*x + 1 // 26 or 1
	z *= y
	y = w + ydiff
	y *= x
	z += y

	return z
}
