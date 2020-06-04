package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Point .
type Point struct {
	X, Y, Moves int
	Direction   string
}

// Distance .
func (p *Point) Distance(other *Point) int {
	distance := math.Sqrt(math.Pow((float64(other.X-p.X)), 2) + math.Pow(float64(other.Y-p.Y), 2))
	return int(distance)
}

func (p *Point) String() string {
	return fmt.Sprintf("p{x:%d,y:%d}", p.X, p.Y)
}

// Move according to definition [direction][positions]
// ex U5 up 5
func (p *Point) Move(movement string) *Point {

	direction := movement[0:1]
	positions := movement[1:]
	pos, _ := strconv.Atoi(positions)

	switch direction {
	case "U":
		return &Point{X: p.X, Y: p.Y + pos, Direction: direction, Moves: pos}
	case "D":
		return &Point{X: p.X, Y: p.Y - pos, Direction: direction, Moves: pos}
	case "R":
		return &Point{X: p.X + pos, Y: p.Y, Direction: direction, Moves: pos}
	case "L":
		return &Point{X: p.X - pos, Y: p.Y, Direction: direction, Moves: pos}
	}

	// Bad point
	return &Point{X: p.X, Y: p.Y}
}

// Line .
type Line struct {
	Parent     *Line
	Start      *Point
	End        *Point
	MoveAction string
	Moves      int
}

func (l *Line) String() string {
	return fmt.Sprintf("l{start:%s,end:%s}", l.Start, l.End)
}

// AbsInt .
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// ClosestIntersectPointDistance .
func ClosestIntersectPointDistance(linesA []Line, linesB []Line) int {

	smallestDistance := -1
	points := IntersectingPointsOnLines(linesA, linesB)

	for _, point := range points {
		distance := AbsInt(point.X) + AbsInt(point.Y)

		if smallestDistance == -1 || distance < smallestDistance {
			smallestDistance = distance
		}
	}

	return smallestDistance

}

// ClosestAbsIntersectPointDistance .
func ClosestAbsIntersectPointDistance(linesA []Line, linesB []Line) int {

	smallestDistance := -1
	distanceValues := IntersectingPointsAbsDistanceOnLines(linesA, linesB)

	for _, distanceValue := range distanceValues {

		if smallestDistance == -1 || distanceValue < smallestDistance {
			smallestDistance = distanceValue
		}
	}

	return smallestDistance

}

// LessIntersectSteps .
func LessIntersectSteps(linesA []Line, linesB []Line) int {

	smallestDistance := -1
	points := IntersectingPointsOnLines(linesA, linesB)

	for _, point := range points {
		distance := AbsInt(point.X) + AbsInt(point.Y)

		if smallestDistance == -1 || distance < smallestDistance {
			smallestDistance = distance
		}
	}

	return smallestDistance

}

// IntersectingPointsOnLines .
func IntersectingPointsOnLines(linesA []Line, linesB []Line) []Point {
	points := make([]Point, 0)
	for _, line := range linesA {
		for _, otherLine := range linesB {
			intersects, intersectionPoint := line.IntersectsSegment(&otherLine)
			if intersects {
				// fmt.Printf("* Intersection line1=%v line2=%v at %s\n", line, otherLine, intersectionPoint)
				points = append(points, *intersectionPoint)
			}
		}
	}

	return points
}

// IntersectingPointsAbsDistanceOnLines .
func IntersectingPointsAbsDistanceOnLines(linesA []Line, linesB []Line) []int {
	distanceValues := make([]int, 0)
	for _, line := range linesA {
		for _, otherLine := range linesB {
			intersects, intersectionPoint := line.IntersectsSegment(&otherLine)
			if intersects {
				// fmt.Printf("* Intersection line1=%v line2=%v at %s\n", line, otherLine, intersectionPoint)
				distanceValues = append(distanceValues,
					line.TotalDistance(intersectionPoint)+otherLine.TotalDistance(intersectionPoint))
			}
		}
	}

	return distanceValues
}

// TotalDistance measure distance to origin
func (l *Line) TotalDistance(point *Point) int {

	parentDistance := 0
	p := l.Parent
	for p != nil {
		parentDistance = parentDistance + p.Moves
		p = p.Parent
	}

	return parentDistance + l.Start.Distance(point)
}

// IntersectsSegment check if lines intersect within each line segment
func (l *Line) IntersectsSegment(b *Line) (bool, *Point) {

	if l.Start.X == 0 && l.Start.Y == 0 {
		// ignore 0,0
		return false, nil
	}

	if b.Start.X == 0 && b.Start.Y == 0 {
		// ignore 0,0
		return false, nil
	}

	// fmt.Printf("line:%s intersects line:%s\n", l, b)

	// Line l represented as a1x + b1y = c1
	a1 := l.End.Y - l.Start.Y
	b1 := l.Start.X - l.End.X
	c1 := a1*l.Start.X + b1*l.Start.Y

	a2 := b.End.Y - b.Start.Y
	b2 := b.Start.X - b.End.X
	c2 := a2*b.Start.X + b2*b.Start.Y

	denominator := a1*b2 - a2*b1

	if denominator == 0 {
		return false, nil
	}

	x := (b2*c1 - b1*c2) / denominator
	y := (a1*c2 - a2*c1) / denominator

	var rx0 = 1.1
	var ry0 = 1.1
	var rx1 = 1.1
	var ry1 = 1.1

	if l.End.X-l.Start.X != 0 {
		rx0 = float64(x-l.Start.X) / float64(l.End.X-l.Start.X)
	}

	if (l.End.Y - l.Start.Y) != 0 {
		ry0 = float64(y-l.Start.Y) / float64(l.End.Y-l.Start.Y)
	}

	if b.End.X-b.Start.X != 0 {
		rx1 = float64(x-b.Start.X) / float64(b.End.X-b.Start.X)
	}

	if b.End.Y-b.Start.Y != 0 {
		ry1 = float64(y-b.Start.Y) / float64(b.End.Y-b.Start.Y)
	}

	// fmt.Printf("rx0=%f ry0=%f rx1=%f ry1=%f line=%s otherLine=%s (x=%d y=%d)\n", rx0, ry0, rx1, ry1, l, b, x, y)

	if ((rx0 >= 0.0 && rx0 <= 1.0) || (ry0 >= 0.0 && ry0 <= 1.0)) && ((rx1 >= 0.0 && rx1 <= 1.0) || (ry1 >= 0.0 && ry1 <= 1.0)) {
		return true, &Point{X: x, Y: y}
	}

	return false, nil

}

// Intersects .
// https://www.geeksforgeeks.org/program-for-point-of-intersection-of-two-lines/
func (l *Line) Intersects(b *Line) (bool, *Point) {

	if l.Start.X == 0 && l.Start.Y == 0 {
		// ignore 0,0
		return false, nil
	}

	// fmt.Printf("line:%s intersects line:%s\n", l, b)
	// Line l represented as a1x + b1y = c1
	a1 := l.End.Y - l.Start.Y
	b1 := l.Start.X - l.End.X
	c1 := a1*l.Start.X + b1*l.Start.Y

	a2 := b.End.Y - b.Start.Y
	b2 := b.Start.X - b.End.X
	c2 := a2*b.Start.X + b2*b.Start.Y

	denominator := a1*b2 - a2*b1

	if denominator == 0 {
		return false, nil
	}

	x := (b2*c1 - b1*c2) / denominator
	y := (a1*c2 - a2*c1) / denominator

	return true, &Point{X: x, Y: y}
}

func linesFromInput(input string) []Line {
	lineSegments := strings.Split(input, ",")
	lines := make([]Line, 0)
	for _, segment := range lineSegments {
		if len(lines) == 0 {
			start := &Point{X: 0, Y: 0}
			end := start.Move(segment)
			lines = append(lines, Line{Start: start, End: end, Parent: nil, MoveAction: segment, Moves: end.Moves})
		} else {
			prevLine := lines[len(lines)-1]
			start := prevLine.End
			end := start.Move(segment)
			lines = append(lines, Line{Start: start, End: end, Parent: &prevLine, MoveAction: segment, Moves: end.Moves})
		}
	}

	return lines
}

func main() {

	var inputA = "R998,U502,R895,D288,R416,U107,R492,U303,R719,D601,R783,D154,L236,U913,R833,D329,R28,D759,L270,D549,L245,U653,L851,U676,L211,D949,R980,U314,L897,U764,R149,D214,L195,D907,R534,D446,R362,D6,L246,D851,L25,U925,L334,U673,L998,U581,R783,U912,R53,D694,L441,U411,L908,D756,R946,D522,L77,U468,R816,D555,L194,D707,R97,D622,R99,D265,L590,U573,R132,D183,L969,D207,L90,D331,R88,D606,L315,U343,R546,U460,L826,D427,L232,U117,R125,U309,R433,D53,R148,U116,L437,U339,L288,D879,L52,D630,R201,D517,L341,U178,R94,U636,L759,D598,L278,U332,R192,U463,L325,U850,L200,U810,L686,U249,L226,D297,R915,D117,R56,D59,R760,U445,R184,U918,R173,D903,R212,D868,L88,D798,L829,U835,L563,U19,R480,D989,R529,D834,R515,U964,L876,D294,R778,D551,L457,D458,R150,D698,R956,D781,L310,D948,R50,U56,R98,U348,L254,U614,L654,D359,R632,D994,L701,D615,R64,D507,R668,D583,L687,D902,L564,D214,R930,D331,L212,U943,R559,U886,R590,D805,R426,U669,L141,D233,L573,D682,L931,U267,R117,D900,L944,U667,L838,D374,L406,U856,R987,D870,R716,D593,R596,D654,R653,U120,L666,U145,R490,D629,R172,D881,L808,D324,R956,D532,L475,U165,L503,U361,R208,U323,R568,D876,R663,D11,L839,D67,R499,U75,L643,U954,R94,D418,R761,D842,L213,D616,L785,D42,R707,D343,L513,D480,L531,D890,L899,D2,L30,D188,R32,U588,R480,U33,R849,U443,L666,U117,L13,D974,L453,U93,R960,D369,R332,D61,L17,U557,R818,D744,L124,U916,L454,D572,R451,D29,R711,D134,R481,U366,L327,U132,L819,U839,R485,U941,R224,U531,R688,U561,R958,D899,L315,U824,L408,D941,R517,D163,L878,U28,R767,D798,R227"
	var inputB = "L1009,U399,R373,U980,L48,U638,R725,U775,R714,D530,L887,D576,L682,D940,L371,D621,L342,D482,R676,D445,R752,U119,L361,D444,L769,D854,L874,U259,R332,U218,R866,U28,L342,D233,L958,U649,R998,U262,L8,D863,L283,D449,L73,D438,L516,D54,R964,D981,R338,U332,L761,U704,L705,D468,L115,U834,R367,D156,R480,U27,R846,U73,R846,D720,R811,D466,L407,U928,R816,U50,R90,D893,L930,D833,L159,D972,L823,U868,R689,D558,L777,D13,R844,D8,L168,U956,L111,D462,L667,U559,L839,U503,R906,D838,R83,D323,L782,U588,R599,D233,L700,U679,L51,U779,L110,D260,L201,U992,L43,D557,L628,D875,L201,U535,L246,D976,L546,D22,R600,D301,L542,D41,R532,U316,L765,D310,L666,D369,R853,U684,L457,U816,L667,U758,R798,U959,R893,D185,L842,U168,R68,D348,R394,D296,R966,D511,L319,U717,L57,U129,R843,U439,L744,D870,L162,D991,R77,D565,R494,U601,L851,U748,L96,U124,L379,D446,L882,U371,R133,U820,L935,D704,L670,D911,L182,U138,R844,U926,L552,D716,L849,U624,R723,U117,R252,D737,L216,U796,R156,U322,R812,D390,L50,D493,L665,U314,L584,U798,L11,U524,R171,U837,R981,U32,L277,U650,L865,U28,R399,U908,R652,D543,L779,D406,L839,D198,L190,D319,L776,U752,R383,D884,R385,D682,R729,D163,R252,U533,L690,D767,R533,D147,R366,U716,R548,U171,R932,U720,L9,D39,R895,U850,L276,D988,L528,U551,L262,D480,L275,D567,R70,D599,L814,U876,R120,U93,L565,U795,L278,D41,R695,D693,R208,U272,L923,U498,R238,U268,L244,U278,R965,U395,R990,U329,L478,D245,R980,D473,L702,U396,R358,U636,R400,D919,R240,U780,L251,D633,L55,D723,L529,U319,L299,D89,L251,D557,L705,D705,L391,D58,R241"

	distance := ClosestIntersectPointDistance(linesFromInput(inputA), linesFromInput(inputB))
	absDistance := ClosestAbsIntersectPointDistance(linesFromInput(inputA), linesFromInput(inputB))
	fmt.Printf("Result: distance=%d absDistance=%d\n", distance, absDistance)
}
