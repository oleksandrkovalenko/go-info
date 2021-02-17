package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p Point) Move(dX int, dY int) {
	p.X += dX
	p.Y += dY
}

type Square struct {
	Center Point
	Length int
}

func (t *Square) Move(dx int, dy int) {
	t.Center.Move(dx, dy)
}

func (t *Square) Area() int {
	return t.Length * t.Length
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length < 0 {
		return nil, fmt.Errorf("length must be >0")
	}
	point := Point{x, y}

	square := &Square{
		Center: point,
		Length: length,
	}
	return square, nil
}

func main() {
	point := Point{1,2}
	point.Move(2,1)
	fmt.Printf("%d, %d\n", point.X, point.Y)

	square, _ := NewSquare(1, 2, 3)
	fmt.Printf("%d, %d\n", square.Center.X, square.Center.Y)
	square.Move(2, 2)
	fmt.Printf("%d, %d\n", square.Center.X, square.Center.Y)
}
