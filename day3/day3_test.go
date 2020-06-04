package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	test1InputA = `R8,U5,L5,D3`
	test1InputB = `U7,R6,D4,L4`

	test2InputA = `R75,D30,R83,U83,L12,D49,R71,U7,L72`
	test2InputB = `U62,R66,U55,R34,D71,R55,D58,R83`

	test3InputA = `R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51`
	test3InputB = `U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`
)

func TestLineIntersection(t *testing.T) {
	testCases := []struct {
		l1           []int // line1 x1,y1, x2,y2
		l2           []int // line2 x1,y1, x2,y2
		intersects   bool
		intersection []int
	}{
		{[]int{3, 5, 3, 2}, []int{6, 3, 2, 3}, true, []int{3, 3}},
		{[]int{0, 2, 8, 2}, []int{0, 0, 5, 0}, false, []int{0, 0}},
	}

	for _, testCase := range testCases {
		line1 := &Line{
			Start: &Point{X: testCase.l1[0], Y: testCase.l1[1]},
			End:   &Point{X: testCase.l1[2], Y: testCase.l1[3]},
		}

		line2 := &Line{
			Start: &Point{X: testCase.l2[0], Y: testCase.l2[1]},
			End:   &Point{X: testCase.l2[2], Y: testCase.l2[3]},
		}

		intersects, point := line1.IntersectsSegment(line2)
		require.Equal(t, testCase.intersects, intersects,
			fmt.Sprintf("expecting l1=%s l2=%s to intersect=%t was %t",
				line1, line2, testCase.intersects, intersects))

		if testCase.intersects {
			assert.Equal(t, testCase.intersection[0], point.X)
			assert.Equal(t, testCase.intersection[1], point.Y)
		} else {
			assert.Nil(t, point)
		}
	}

}

func TestClosestIntersectPoint(t *testing.T) {

	testCases := []struct {
		inA                 string
		inB                 string
		expectedDistance    int
		expectedAbsDistance int
	}{
		{test1InputA, test1InputB, 6, 30},
		{test2InputA, test2InputB, 159, 610},
		{test3InputA, test3InputB, 135, 410},
	}

	for _, testCase := range testCases {
		linesA := linesFromInput(testCase.inA)
		linesB := linesFromInput(testCase.inB)
		distance := ClosestIntersectPointDistance(linesA, linesB)
		absDistance := ClosestAbsIntersectPointDistance(linesA, linesB)

		fmt.Printf("distance=%d absDistance=%d", distance, absDistance)
		assert.Equal(t, testCase.expectedDistance, distance,
			fmt.Sprintf("ina:%s inb:%s expected [%d!=%d]", testCase.inA, testCase.inB, testCase.expectedDistance, distance),
		)

		assert.Equal(t, testCase.expectedAbsDistance, absDistance,
			fmt.Sprintf("ina:%s inb:%s expected abs distance [%d!=%d]", testCase.inA, testCase.inB, testCase.expectedAbsDistance, absDistance),
		)
	}

	t.FailNow()
}

func TestRunCode(t *testing.T) {

	// rx0 NaN ry0 1.4 rx1 1 ry1 NaN
	// rx0 NaN ry0 0.6666666666666666 rx1 0.75 ry1 NaN
	// res { x: 3, y: 3 }

	linesA := linesFromInput(test1InputA)
	linesB := linesFromInput(test1InputB)

	for _, line := range linesA {
		fmt.Printf("LineA %v\n", line)
	}

	for _, line := range linesB {
		fmt.Printf("LineB %v\n", line)
	}

	for _, line := range linesA {
		for _, otherLine := range linesB {
			intersects, intersectionPoint := line.IntersectsSegment(&otherLine)
			if intersects {
				fmt.Printf("* Intersection line1=%v line2=%v at %s\n", line, otherLine, intersectionPoint)
			}
		}
	}

}
