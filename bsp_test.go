package bsp

import "testing"

func BenchmarkA(b *testing.B) {
	b.Log(GetCoef(Line{
		A: Point{0, 0},
		B: Point{0, 1},
	}))

	b.Log(GetCoef(Line{
		A: Point{2, 5},
		B: Point{0, 1},
	}))
	b.Log(GetCoef(Line{
		A: Point{0, 1},
		B: Point{2, 5},
	}))
}

func BenchmarkB(b *testing.B) {
	// x=0
	l1 := Line{A: Point{0, 0}, B: Point{0, 1}}
	// x=1
	l2 := Line{A: Point{2, 1}, B: Point{2, 5}}
	// y=1
	l3 := Line{A: Point{3, 1}, B: Point{5, 1}}
	// y=2
	l4 := Line{A: Point{2, 2}, B: Point{8, 2}}
	// y=2x+1
	l5 := Line{A: Point{0, 1}, B: Point{2, 5}}
	// y=-3x+5
	l6 := Line{A: Point{0, 5}, B: Point{1, 2}}

	b.Log(l1, l2, "expected 0 {0, 0}")
	b.Log(GetIntersect(l1, l2))

	b.Log(l1, l3, "expected 2 {0, 1}")
	b.Log(GetIntersect(l1, l3))

	b.Log(l3, l4, "expected 0 {0, 0}")
	b.Log(GetIntersect(l3, l4))

	b.Log(l5, l1, "expected 2 {0, 1}")
	b.Log(GetIntersect(l5, l1))

	b.Log(l6, l2, "expected 2 {2, -1}")
	b.Log(GetIntersect(l6, l2))

	b.Log(l5, l3, "expected 2 {0, 1}")
	b.Log(GetIntersect(l5, l3))

	b.Log(l6, l4, "expected 2 {1, 2}")
	b.Log(GetIntersect(l6, l4))

	b.Log(l5, l6, "expected 2 {4/5, 13/5}")
	b.Log(GetIntersect(l5, l6))

	b.Log(l1, l1, "expected 1 {0, 0}")
	b.Log(GetIntersect(l1, l1))
	b.Log(l3, l3, "expected 1 {0, 0}")
	b.Log(GetIntersect(l3, l3))
	b.Log(l6, l6, "expected 1 {0, 0}")
	b.Log(GetIntersect(l6, l6))
}

func TestC(b *testing.T) {
	var lines []Line = []Line{
		// {Point{0, 0}, Point{0, 0}},
		{Point{0, 0}, Point{4, 0}},
		{Point{0, 1}, Point{0, -1}},
		{Point{7, 2}, Point{5, -2}},
		{Point{-2, 2}, Point{1, -3}},
	}
	bsp := GenerateBSP(lines, IsInPlan)

	b.Log(bsp.String())
	b.Fatal()
}
