package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

const (
	trimValues = " \n\t"
)

// Coordinate3D .
type Coordinate3D struct {
	x int
	y int
	z int
}

// MoonPair .
type MoonPair struct {
	A *Moon
	B *Moon
}

// Moon .
type Moon struct {
	position Coordinate3D
	velocity Coordinate3D
}

// MoonOrbitalState .
type MoonOrbitalState struct {
	coordinatesX []int
	coordinatesY []int
	coordinatesZ []int

	doneX bool
	doneY bool
	doneZ bool
}

func (moon *Moon) String() string {
	return fmt.Sprintf("p<x=%d y=%d z=%d> v<x=%d y=%d z=%d>",
		moon.position.x,
		moon.position.y,
		moon.position.z,
		moon.velocity.x,
		moon.velocity.y,
		moon.velocity.z,
	)
}

// IsSame check if two moons have the same position and velocity
func (moon *Moon) IsSame(other Moon) bool {
	return moon.position.x == other.position.x &&
		moon.position.y == other.position.y &&
		moon.position.z == other.position.z &&
		moon.velocity.x == other.velocity.x &&
		moon.velocity.y == other.velocity.y &&
		moon.velocity.z == other.velocity.z
}

// IsSamePosition check if two moons are at the same position
func (moon *Moon) IsSamePosition(other Moon) bool {
	return moon.position.x == other.position.x &&
		moon.position.y == other.position.y &&
		moon.position.z == other.position.z
}

// PotentialEnergy .
func (moon *Moon) PotentialEnergy() int {
	return AbsInt(moon.position.x) +
		AbsInt(moon.position.y) +
		AbsInt(moon.position.z)
}

// KineticEnergy .
func (moon *Moon) KineticEnergy() int {
	return AbsInt(moon.velocity.x) +
		AbsInt(moon.velocity.y) +
		AbsInt(moon.velocity.z)
}

// TotalEnergy .
func (moon *Moon) TotalEnergy() int {
	return moon.PotentialEnergy() * moon.KineticEnergy()
}

// ApplyVelocity .
func (moon *Moon) ApplyVelocity() {
	moon.position.x = moon.position.x + moon.velocity.x
	moon.position.y = moon.position.y + moon.velocity.y
	moon.position.z = moon.position.z + moon.velocity.z
}

// ApplyGravity .
func (moon *Moon) ApplyGravity(other *Moon) {
	if other.position.x > moon.position.x {
		moon.velocity.x++
		other.velocity.x--
	} else if other.position.x < moon.position.x {
		other.velocity.x++
		moon.velocity.x--
	}

	if other.position.y > moon.position.y {
		moon.velocity.y++
		other.velocity.y--
	} else if other.position.y < moon.position.y {
		moon.velocity.y--
		other.velocity.y++
	}

	if other.position.z > moon.position.z {
		moon.velocity.z++
		other.velocity.z--
	} else if other.position.z < moon.position.z {
		moon.velocity.z--
		other.velocity.z++
	}
}

// MakePairs .
func MakePairs(items []*Moon) []*MoonPair {
	var rest = items
	results := make([]*MoonPair, 0)
	for len(rest) > 0 {
		first := rest[0]
		rest = rest[1:]
		for _, other := range rest {
			pair := &MoonPair{A: first, B: other}
			results = append(results, pair)
		}
	}
	return results
}

// LoadMoons .
func LoadMoons(inputCoordinates string) []*Moon {
	fmt.Printf("Load orbits\n---\n%s\n---\n", inputCoordinates)
	inputValues := strings.Split(inputCoordinates, "\n")
	moons := make([]*Moon, len(inputValues))
	for moonIndex, input := range inputValues {
		keyValues := make(map[string]string, 0)
		for _, keyValuePair := range strings.Split(input[1:len(input)-1], ",") {
			pair := strings.Split(keyValuePair, "=")
			keyValues[strings.Trim(pair[0], trimValues)] = strings.Trim(pair[1], trimValues)
		}
		z, _ := strconv.Atoi(keyValues["z"])
		x, _ := strconv.Atoi(keyValues["x"])
		y, _ := strconv.Atoi(keyValues["y"])
		moon := &Moon{
			position: Coordinate3D{
				z: z,
				x: x,
				y: y,
			},
			velocity: Coordinate3D{
				z: 0, x: 0, y: 0,
			},
		}

		moons[moonIndex] = moon
	}

	return moons
}

// RunSimulation .
func RunSimulation(moons []*Moon, iterations int) int {
	pairs := MakePairs(moons)

	accumulatedSystemEnergy := 0
	lastSystemEnergy := 0
	for _, moon := range moons {
		lastSystemEnergy += moon.TotalEnergy()
	}

	// for _, moon := range moons {
	// 	fmt.Printf("%s\n", moon)
	// }

	for step := 0; step < iterations; step++ {
		for _, pair := range pairs {
			pair.A.ApplyGravity(pair.B)
		}

		for _, moon := range moons {
			moon.ApplyVelocity()
		}

		systemEnergy := 0
		for _, moon := range moons {
			systemEnergy += moon.TotalEnergy()
			//fmt.Printf("%d - %s\n", step+1, moon)
		}

		accumulatedSystemEnergy = accumulatedSystemEnergy + systemEnergy
		lastSystemEnergy = systemEnergy
	}

	systemEnergy := 0
	for _, moon := range moons {
		systemEnergy += moon.TotalEnergy()
	}

	return systemEnergy
}

// CalculateOrbits .
func CalculateOrbits(moons []*Moon) [][]Moon {

	initialState := make([]Moon, len(moons))
	orbits := make([][]Moon, len(moons))
	orbitComplete := make([]bool, len(moons))

	for index, moon := range moons {

		initialState[index] = Moon{
			position: moon.position,
			velocity: moon.velocity,
		}

		orbits[index] = make([]Moon, 0)
		// Add initial state as first orbit
		orbits[index] = append(orbits[index], initialState[index])
		orbitComplete[index] = false
	}

	pairs := MakePairs(moons)
	allComplete := false
	for step := 0; step < 90000000; step++ {
		for _, pair := range pairs {
			pair.A.ApplyGravity(pair.B)
		}

		for moonIndex, moon := range moons {
			moon.ApplyVelocity()
			if !orbitComplete[moonIndex] {
				if moon.IsSame(initialState[moonIndex]) {
					//fmt.Printf("moon: %+v", moon)
					orbitComplete[moonIndex] = true
					continue
				}

				position := Moon{
					position: moon.position,
					velocity: moon.velocity,
				}
				orbits[moonIndex] = append(orbits[moonIndex], position)
			}
		}

		for _, orbitComplete := range orbitComplete {
			allComplete = orbitComplete && allComplete
		}
	}

	return orbits
}

// CalculateOrbitalRepetitions create unique positions on each axis
func CalculateOrbitalRepetitions(moons []*Moon) (int, int, int) {
	pairs := MakePairs(moons)

	seenx := make([]string, 0)
	seeny := make([]string, 0)
	seenz := make([]string, 0)

	statex := make([]string, len(moons))
	statey := make([]string, len(moons))
	statez := make([]string, len(moons))

	done := make([]bool, 3) // x,y,z done flags

	for moonIndex, moon := range moons {
		// initial state
		statex[moonIndex] = strconv.Itoa(moon.position.x)
		statey[moonIndex] = strconv.Itoa(moon.position.y)
		statez[moonIndex] = strconv.Itoa(moon.position.z)
	}

	//fmt.Printf("step: %d x %v\n", 0, statex)
	//fmt.Printf("step: %d y %v\n", 0, statey)
	//fmt.Printf("step: %d z %v\n", 0, statez)

	x := strings.Join(statex, ",")
	y := strings.Join(statey, ",")
	z := strings.Join(statez, ",")

	initialx := x
	initialy := y
	initialz := z

	seenx = append(seenx, x)
	seeny = append(seeny, y)
	seenz = append(seenz, z)

	// has a limit to avoid infinite loops
	// for step := 0; step < 1000000; step++ {
	step := 0
	for {
		for _, pair := range pairs {
			pair.A.ApplyGravity(pair.B)
		}

		for moonIndex, moon := range moons {
			moon.ApplyVelocity()

			statex[moonIndex] = strconv.Itoa(moon.position.x)
			statey[moonIndex] = strconv.Itoa(moon.position.y)
			statez[moonIndex] = strconv.Itoa(moon.position.z)
		}

		//fmt.Printf("step: %d x %v\n", step+1, statex)
		//fmt.Printf("step: %d y %v\n", step+1, statey)
		//fmt.Printf("step: %d z %v\n", step+1, statez)

		x := strings.Join(statex, ",")
		y := strings.Join(statey, ",")
		z := strings.Join(statez, ",")

		if done[0] && done[1] && done[2] {
			break
		}

		if initialx == x {
			done[0] = true
		}

		if initialy == y {
			done[1] = true
		}

		if initialz == z {
			done[2] = true
		}

		if !done[0] {
			seenx = append(seenx, x)
		}

		if !done[1] {
			seeny = append(seeny, y)
		}

		if !done[2] {
			seenz = append(seenz, z)
		}

		// fmt.Printf("statex:%d %v\n", step+1, x)
		// fmt.Printf("statey:%d %v\n", step+1, y)
		// fmt.Printf("statez:%d %v\n", step+1, z)

		step++
	}

	//fmt.Printf("seenx:%d %v\n", len(seenx), seenx)
	//fmt.Printf("seeny:%d %v\n", len(seeny), seeny)
	//fmt.Printf("seenz:%d %v\n", len(seenz), seenz)

	return len(seenx), len(seeny), len(seenz)
}

// CalculateBigBang play god ?
func CalculateBigBang(xRepeats int, yRepeats int, zRepeatsint int) int {

	x := big.NewInt(int64(xRepeats + 1))
	y := big.NewInt(int64(yRepeats + 1))
	//z := big.NewInt(int64(zRepeatsint + 1))

	div := x.GCD(nil, nil, x, y)
	acc := x.Mul(x, y)

	gcd := x.GCD(nil, nil, x, y)
	//result := x.Mul(acc, z)

	fmt.Printf("div: %+v acc: %+v result: %+v", div, acc, gcd)
	return -1
}

func main() {

	input := `<x=-13, y=-13, z=-13>
<x=5, y=-8, z=3>
<x=-6, y=-10, z=-3>
<x=0, y=5, z=-5>`

	// moons := LoadMoons(input)
	// totalSystemEnergy := RunSimulation(moons, 1000)
	// fmt.Printf("totalSystemEnergy: %d\n", totalSystemEnergy)

	moons2 := LoadMoons(input)
	repeatx, repeaty, repeatz := CalculateOrbitalRepetitions(moons2)
	fmt.Printf("Calculate coordinates done xrep=%d yrep=%d zrep=%d\n", repeatx, repeaty, repeatz)

	/*
		import math
		>>> 268296 * 231614 // math.gcd(268296,231614)
		31070554872
		>>> 23326 * 31070554872 // math.gcd(23326,31070554872)
		362375881472136
	*/
}

// AbsInt .
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
