package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// AbsInt .
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// https://stackoverflow.com/questions/11907947/how-to-check-if-a-point-lies-on-a-line-between-2-other-points
func pointOnLine(p1 []int, p2 []int, targetPoint []int) bool {
	dxc := targetPoint[0] - p1[0]
	dyc := targetPoint[1] - p1[1]

	dxl := p2[0] - p1[0]
	dyl := p2[1] - p1[1]

	cross := dxc*dyl - dyc*dxl

	return cross == 0
}

func angleOfLine(p1 []int, p2 []int) float64 {
	// https://stackoverflow.com/questions/49425796/how-to-get-the-angle-of-two-cross-line
	angle := math.Atan2(float64(p2[1]-p1[1]), float64(p2[0]-p1[0]))
	return angle
}

func angleOfLineDeg(p1 []int, p2 []int) float64 {
	return angleOfLine(p1, p2) * 180 / math.Pi
}

func angleOfLineDeg360(p1 []int, p2 []int) float64 {
	// (x > 0 ? x : (2*PI + x)) * 360 / (2*PI)
	// https://stackoverflow.com/questions/1311049/how-to-map-atan2-to-degrees-0-360
	// https://stackoverflow.com/questions/9614109/how-to-calculate-an-angle-from-points
	angle := angleOfLine(p1, p2) * 180 / math.Pi

	if angle < 0 {
		return 360 + angle
	}

	return angle
}

func rotateAngle(angle float64, degrees float64) float64 {
	angle = angle + degrees
	if angle < 0 {
		return angle + 360
	}

	if angle >= 360 {
		return angle - 360
	}

	return angle
}

func distance(p1 []int, p2 []int) float64 {
	acc := math.Pow(float64(p2[0]-p1[0]), 2.0) + math.Pow(float64(p2[1]-p1[1]), 2)
	return math.Sqrt(acc)
}

func pointOnLineBetween(p1 []int, p2 []int, targetPoint []int) bool {
	dxc := targetPoint[0] - p1[0]
	dyc := targetPoint[1] - p1[1]

	dxl := p2[0] - p1[0]
	dyl := p2[1] - p1[1]

	cross := dxc*dyl - dyc*dxl

	if cross != 0 {
		return false
	}

	if AbsInt(dxl) >= AbsInt(dyl) {
		if dxl > 0 {
			return p1[0] <= targetPoint[0] && targetPoint[0] <= p2[0]
		} else {
			return p2[0] <= targetPoint[0] && targetPoint[0] <= p1[0]
		}
	} else {
		if dyl > 0 {
			return p1[1] <= targetPoint[1] && targetPoint[1] <= p2[1]
		} else {
			return p2[1] <= targetPoint[1] && targetPoint[1] <= p1[1]
		}
	}
}

func slope(p1 []int, p2 []int) float32 {
	if p2[0] == p1[0] {
		return 1000000
	}

	return float32(p2[1]-p1[1]) / float32(p2[0]-p1[0])
}

func topPoint(result map[string]int) (string, int) {
	max := 0
	maxPoint := ""
	for key, val := range result {
		if val > max {
			max = val
			maxPoint = key
		}
	}
	return maxPoint, max
}

func reducePoints(mapdata [][]int) map[string]int {
	results := make(map[string]int)
	for _, p1 := range mapdata {
		visibleAsteroids := 0

		// keep a record of asteroids which were
		// discarded so they do not get attempted again
		eliminated := make(map[string]bool)

		//fmt.Printf("Testing %v\n", p1)
		for _, p2 := range mapdata {
			if p1[0] == p2[0] && p1[1] == p2[1] {
				// same point skip
				continue
			}

			p2String := fmt.Sprintf("%d,%d", p2[0], p2[1])
			if _, ok := eliminated[p2String]; ok {
				// skip asteroids which we should not evaluate
				// continue
			}

			found := false
			for _, otherPoint := range mapdata {
				if p1[0] == otherPoint[0] && p1[1] == otherPoint[1] {
					// same point p1 skip
					continue
				}
				if p2[0] == otherPoint[0] && p2[1] == otherPoint[1] {
					// same point p2 skip
					continue
				}

				otherPointString := fmt.Sprintf("%d,%d", otherPoint[0], otherPoint[1])
				if _, ok := eliminated[otherPointString]; ok {
					// skip asteroids which we should not evaluate
					// continue
				}

				// fmt.Printf("Testing %v %v %v\n", p1, p2, otherPoint)
				found = pointOnLineBetween(p1, p2, otherPoint)
				if found {
					//fmt.Printf("Found %v %v %v\n", p1, p2, otherPoint)
					eliminated[p2String] = true
					break
				}
			}

			if !found {
				//fmt.Printf("Visibible %v -> %v \n", p1, p2)
				visibleAsteroids++
			}
		}
		coordString := fmt.Sprintf("%d,%d", p1[0], p1[1])
		results[coordString] = visibleAsteroids
	}

	return results
}

func vaporize(mapdata [][]int, origin []int) map[string]int {

	distinctAngles := make([]float64, 0)
	angles := make(map[float64][]string, 0)
	distances := make(map[string]float64, 0)
	nukedAsteroids := make(map[string]int, 0)

	//fmt.Printf("Testing %v\n", p1)
	for _, p2 := range mapdata {
		if origin[0] == p2[0] && origin[1] == p2[1] {
			// same point skip
			continue
		}

		angle := angleOfLineDeg360(origin, p2)
		// normalise for orientation of cartesian & asteroid
		angle = rotateAngle(angle, -270)

		distance := distance(origin, p2)
		p2id := fmt.Sprintf("%d,%d", p2[0], p2[1])

		if _, ok := angles[angle]; !ok {
			angles[angle] = []string{p2id}
			distinctAngles = append(distinctAngles, angle)
		} else {
			angles[angle] = append(angles[angle], p2id)
		}

		distances[p2id] = distance
	}

	//fmt.Printf("angles: %v\n\n", angles)
	//fmt.Printf("distinctAngles: %v\n\n", distinctAngles)
	//fmt.Printf("distances: %v\n\n", distances)

	sort.Slice(distinctAngles, func(i int, j int) bool {
		return distinctAngles[i] < distinctAngles[j]
	})

	//fmt.Printf("distinctAngles sorted: %v\n\n", distinctAngles)

	// remove one from total asteroids
	otherAsteroids := len(mapdata) - 1
	pass := 1
	nukeorder := 1

	for len(nukedAsteroids) < otherAsteroids {
		fmt.Printf("Pass %d\n", pass)
		for _, angle := range distinctAngles {
			asteroids := angles[angle]
			//fmt.Printf("Angle %f -> %v \n", angle, asteroids)

			smallestDistance := 999999999.0
			smallestTarget := ""
			// find smallest distance
			for _, target := range asteroids {

				if _, ok := nukedAsteroids[target]; ok {
					// already nuked skip
					continue
				}

				//fmt.Printf("distance target=%s distance=%f \n", target, distances[target])

				if smallestDistance > distances[target] {
					smallestDistance = distances[target]
					smallestTarget = target
				}
			}

			if smallestTarget != "" {
				// nuke smallest smallestTarget
				fmt.Printf("NUKE order=%d asteroid=%s\n", nukeorder, smallestTarget)
				nukedAsteroids[smallestTarget] = nukeorder
				nukeorder++
			}
		}
		pass++
	}

	return nukedAsteroids
}

func loadMap(mapData string) [][]int {

	res := make([][]int, 0)
	lines := strings.Split(mapData, "\n")
	for yIndex, line := range lines {
		for xIndex, char := range line {
			if char == '#' {
				//fmt.Printf("pos: x=%d,y=%d\n", xIndex, yIndex)
				res = append(res, []int{xIndex, yIndex})
			}
		}
	}

	return res
}

var (
	map1 = `..............#.#...............#....#....
#.##.......#....#.#..##........#...#......
..#.....#....#..#.#....#.....#.#.##..#..#.
...........##...#...##....#.#.#....#.##..#
....##....#...........#..#....#......#.###
.#...#......#.#.#.#...#....#.##.##......##
#.##....#.....#.....#...####........###...
.####....#.......#...##..#..#......#...#..
...............#...........#..#.#.#.......
........#.........##...#..........#..##...
...#..................#....#....##..#.....
.............#..#.#.........#........#.##.
...#.#....................##..##..........
.....#.#...##..............#...........#..
......#..###.#........#.....#.##.#......#.
#......#.#.....#...........##.#.....#..#.#
.#.............#..#.....##.....###..#..#..
.#...#.....#.....##.#......##....##....#..
.........#.#..##............#..#...#......
..#..##...#.#..#....#..#.#.......#.##.....
#.......#.#....#.#..##.#...#.......#..###.
.#..........#...##.#....#...#.#.........#.
..#.#.......##..#.##..#.......#.###.......
...#....###...#......#..#.....####........
.............#.#..........#....#......#...
#................#..................#.###.
..###.........##...##..##.................
.#.........#.#####..#...##....#...##......
........#.#...#......#.................##.
.##.....#..##.##.#....#....#......#.#....#
.....#...........#.............#.....#....
........#.##.#...#.###.###....#.#......#..
..#...#.......###..#...#.##.....###.....#.
....#.....#..#.....#...#......###...###...
#..##.###...##.....#.....#....#...###..#..
........######.#...............#...#.#...#
...#.....####.##.....##...##..............
###..#......#...............#......#...#..
#..#...#.#........#.#.#...#..#....#.#.####
#..#...#..........##.#.....##........#.#..
........#....#..###..##....#.#.......##..#
.................##............#.......#..`
)

func main() {
	mapData := loadMap(map1)
	results := reducePoints(mapData)
	//fmt.Printf("%+v\n", results)
	result, count := topPoint(results)
	fmt.Printf("%+v %+v\n", result, count)
	fmt.Println()
	vaporiseResults := vaporize(mapData, []int{26, 36})

	for id, order := range vaporiseResults {
		if order == 200 {
			fmt.Printf("vaporise order 200 %s\n", id)
		}
	}
	fmt.Printf("vaporiseResults %+v\n", vaporiseResults)
}
