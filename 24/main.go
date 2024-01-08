package main

import (
	"fmt"
	"math"
)

// Структура Point представляет точку на плоскости.
type Point struct {
	x, y float64
}

// Конструктор для структуры Point.
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Метод DistanceTo вычисляет расстояние до другой точки.
func (p *Point) DistanceTo(other *Point) float64 {
	dx := other.x - p.x
	dy := other.y - p.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(1.0, 2.0)
	p2 := NewPoint(4.0, 6.0)
	fmt.Println(p1.DistanceTo(p2))
}
