package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func maxPoints(points []Point) int {
	if len(points) < 3 {
		return len(points)
	}
	var max int
	for i := range points {
		verticalNum := 0
		equalNum := 0
		m := map[float64]int{}
		for j := range points {
			if i == j {
				continue
			}
			if points[i].X == points[j].X && points[i].Y == points[j].Y {
				equalNum++
			} else if points[i].X == points[j].X {
				verticalNum++
			} else {
				ratio := float64(points[i].Y-points[j].Y) / float64(points[i].X-points[j].X)
				m[ratio]++
			}
		}
		if verticalNum+equalNum > max {
			max = verticalNum + equalNum
		}
		for k := range m {
			if m[k]+equalNum > max {
				max = m[k] + equalNum
			}
		}
	}
	return max + 1
}

func main() {
	res := maxPoints([]Point{
		{
			X: 0,
			Y: 0,
		}, {
			X: 94911151,
			Y: 94911150,
		}, {
			X: 94911152,
			Y: 94911151,
		},
	})

	fmt.Printf("%#v\n", res)
}
