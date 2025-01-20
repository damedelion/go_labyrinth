package utils

type Point struct {
	Row, Col int
}

type Node struct {
	Point
	Distance int
	Index    int
}
