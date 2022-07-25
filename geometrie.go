package bsp

import "fmt"

type Point struct {
	X, Y float64
}

type Line struct {
	A, B Point
}

// Form: V=false => Ax+B or V=true => x=A
type LineCoef struct {
	V    bool
	A, B float64
}

type IntersectResult int8

const (
	IntersectResult_NONE  IntersectResult = 0
	IntersectResult_EQ    IntersectResult = 1
	IntersectResult_POINT IntersectResult = 2
)

func (this Point) String() string {
	return fmt.Sprintf("{%.2f, %.2f}", this.X, this.Y)
}

func (this Line) String() string {
	return fmt.Sprintf("[%s, %s]", this.A.String(), this.B.String())
}

func (this Line) Valid() bool {
	return this.A != this.B
}

// Is p1 between p2 & p3
func betweenP(p1, p2, p3 Point) bool {
	return betweenN(p1.X, p2.X, p3.X) && betweenN(p1.Y, p2.Y, p3.Y)
}

// Is n1 between n2 & n3
func betweenN(n1, n2, n3 float64) bool {
	return (n2 <= n1 && n1 <= n3) || (n3 <= n1 && n1 <= n2)
}

// Is p is Ip or Left of the l
func isUpP(l Line, p Point) bool {
	coef := GetCoef(l)
	return (coef.V && coef.A < p.X) || (!coef.V && coef.A*p.X+coef.B < p.Y)
}

// Is l2.A is Up or Left of the l1
func isUpL(l1, l2 Line) bool {
	return isUpP(l1, l2.A)
}

// l1 is consider infinite because it is the plan that is used to cut all the items (lines)
func IsInPlan(l1, l2 Line) (BSPGenerationResult, Line, Line) {
	inter, p := GetIntersect(l1, l2)
	switch inter {
	case IntersectResult_EQ:
		return BSPGenerationResult_IN, Line{}, Line{}
	case IntersectResult_POINT:
		if betweenP(p, l2.A, l2.B) {
			if isUpP(l1, l2.A) {
				return BSPGenerationResult_CUT, Line{A: l2.A, B: p}, Line{A: p, B: l2.B}
			} else {
				return BSPGenerationResult_CUT, Line{A: p, B: l2.B}, Line{A: l2.A, B: p}
			}
		}
		// Goto check up or bottom
	case IntersectResult_NONE:
		// Goto check up or bottom
	}

	// Check up or bottom
	if isUpL(l1, l2) {
		return BSPGenerationResult_UP, Line{}, Line{} // Left or Up
	} else {
		return BSPGenerationResult_BOTTOM, Line{}, Line{} // Right or Bottom
	}
}

// TODO Memoize the result
func GetCoef(l Line) LineCoef {
	if l.A.X == l.B.X {
		return LineCoef{V: true, A: l.A.X}
	}
	a := (l.B.Y - l.A.Y) / (l.B.X - l.A.X)
	b := l.A.Y - (l.A.X * a)
	return LineCoef{A: a, B: b}
}

// TODO Memoize the result
func GetIntersect(l1, l2 Line) (IntersectResult, Point) {
	l1Coef := GetCoef(l1)
	l2Coef := GetCoef(l2)

	// if double vertical or double horizontal
	if l1Coef.V && l2Coef.V || (!l1Coef.V && !l2Coef.V && l1Coef.A == l2Coef.A) {
		if l1Coef == l2Coef {
			return IntersectResult_EQ, Point{}
		} else {
			return IntersectResult_NONE, Point{}
		}
	}
	// l1 is vertical and l2 is not vertical
	if l1Coef.V {
		return IntersectResult_POINT, Point{X: l1Coef.A, Y: l2Coef.A*l1Coef.A + l2Coef.B}
	}
	// l1 is not vertical and l2 is vertical
	if l2Coef.V {
		return IntersectResult_POINT, Point{X: l2Coef.A, Y: l1Coef.A*l2Coef.A + l1Coef.B}
	}

	x := (l2Coef.B - l1Coef.B) / (l1Coef.A - l2Coef.A)
	y := l1Coef.A*x + l1Coef.B
	return IntersectResult_POINT, Point{X: x, Y: y}
}
