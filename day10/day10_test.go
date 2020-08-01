package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testMap1 = `.#..#
.....
#####
....#
...##`

	testMap2 = `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`

	testMap3 = `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`

	testMap4 = `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`

	testMap5 = `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`

	testMap6 = `.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....X...###..
..#.#.....#....##`
)

func TestAngle(t *testing.T) {
	angle := angleOfLine([]int{1, 1}, []int{2, 2})
	angleDeg := angleOfLineDeg([]int{1, 1}, []int{2, 2})
	fmt.Printf("Ang: %f %f\n", angle, angleDeg)

	distance := distance([]int{0, 0}, []int{1, 0})
	fmt.Printf("distance: %f\n", distance)

	t.Fail()
}

func TestMapLoad(t *testing.T) {
	mapData := loadMap(testMap1)
	assert.Equal(t, 10, len(mapData))
	assert.Equal(t, 2, len(mapData[0]))

}

func TestSlopesLoad2(t *testing.T) {
	results := reducePoints(loadMap(testMap2))
	fmt.Printf("%+v\n", results)
	result, count := topPoint(results)
	fmt.Printf("%+v %+v\n", result, count)
	assert.Equal(t, "5,8", result)
	assert.Equal(t, 33, count)
	t.Fail()
}

func TestAsteroidMaps(t *testing.T) {
	testCases := []struct {
		mapData    string
		bestResult string
		bestCount  int
	}{
		{testMap1, "3,4", 8},
		{testMap2, "5,8", 33},
		{testMap3, "1,2", 35},
		{testMap4, "6,3", 41},
		{testMap5, "11,13", 210},
	}

	for _, testCase := range testCases {
		results := reducePoints(loadMap(testCase.mapData))
		fmt.Printf("%+v\n", results)
		result, count := topPoint(results)
		fmt.Printf("%+v %+v\n", result, count)
		assert.Equal(t, testCase.bestResult, result)
		assert.Equal(t, testCase.bestCount, count)
	}

	t.Fail()

}

func TestVaporise(t *testing.T) {
	mapData := loadMap(testMap6)
	results := vaporize(mapData, []int{8, 3})
	assert.Equal(t, 1, results["8,1"])
	assert.Equal(t, 2, results["9,0"])
}

func TestVaporiseBig(t *testing.T) {
	mapData := loadMap(testMap5)
	results := vaporize(mapData, []int{11, 13})
	assert.Equal(t, 1, results["11,12"])
	assert.Equal(t, 2, results["12,1"])
	assert.Equal(t, 10, results["12,8"])
	assert.Equal(t, 20, results["16,0"])
	assert.Equal(t, 50, results["16,9"])
	assert.Equal(t, 100, results["10,16"])
	assert.Equal(t, 199, results["9,6"])
	assert.Equal(t, 200, results["8,2"])
	assert.Equal(t, 201, results["10,9"])
	assert.Equal(t, 299, results["11,1"])
}
