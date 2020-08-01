package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testInput1 = `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`

	testInput2 = `<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`
)

func TestVelocityAndGravity(t *testing.T) {
	moons := LoadMoons(testInput1)
	pairs := MakePairs(moons)

	for step := 0; step < 10; step++ {
		for _, pair := range pairs {
			pair.A.ApplyGravity(pair.B)
		}

		for _, moon := range moons {
			moon.ApplyVelocity()
		}
	}

	systemEnergy := 0
	for _, moon := range moons {
		systemEnergy += moon.TotalEnergy()
	}

	fmt.Printf("System energy %d\n", systemEnergy)
	assert.Equal(t, 179, systemEnergy)
}

func TestVelocityAndGravity2(t *testing.T) {
	moons := LoadMoons(testInput2)
	pairs := MakePairs(moons)
	for step := 0; step < 100; step++ {

		for _, pair := range pairs {
			pair.A.ApplyGravity(pair.B)
		}

		for _, moon := range moons {
			moon.ApplyVelocity()
		}
	}

	systemEnergy := 0
	for _, moon := range moons {
		systemEnergy += moon.TotalEnergy()
	}
	assert.Equal(t, 1940, systemEnergy)
}

func TestVelocityAndGravity1a(t *testing.T) {

	moons := LoadMoons(testInput1)
	totalSystemEnergy := RunSimulation(moons, 2772)
	assert.Equal(t, 0, totalSystemEnergy)
}

func TestVelocityAndGravity2a(t *testing.T) {

	moons := LoadMoons(testInput2)
	totalSystemEnergy := RunSimulation(moons, 2775)
	assert.Equal(t, 2775, totalSystemEnergy)
}

func TestGetOrbits1(t *testing.T) {
	moons := LoadMoons(testInput1)
	orbits := CalculateOrbits(moons)

	for index, orbit := range orbits {
		fmt.Printf("orbits: %d %d\n", index, len(orbit))
	}

	timeInSpace := 1

	//fmt.Printf("orbits: %v\n", orbits)

	for true {
		isBigBang := true
		for moonIndex := range moons {
			moonPosition := orbits[moonIndex][timeInSpace%len(orbits[moonIndex])]
			//fmt.Printf("timeInSpace: %d %d %+v \n", timeInSpace, moonIndex, moonPosition)

			if moonPosition.IsSamePosition(orbits[moonIndex][0]) {
				fmt.Printf("*** *** space: %d %d %+v %+v\n",
					timeInSpace, moonIndex, moonPosition, orbits[moonIndex][0])
			}

			if moonPosition.IsSame(orbits[moonIndex][0]) {
				fmt.Printf("*** timeInSpace: %d %d %+v %+v\n",
					timeInSpace, moonIndex, moonPosition, orbits[moonIndex][0])
			} else {
				isBigBang = false
			}

		}

		if isBigBang {
			fmt.Printf("Detected big bang event at light year %d\n", timeInSpace)
			break
		}

		timeInSpace++

		if timeInSpace > 4000 {
			break
		}
	}

	t.Fail()
}

func TestCalculateOrbitalRepetitions1(t *testing.T) {

	fmt.Printf("Calculate orbits\n")
	moons := LoadMoons(testInput1)
	repeatx, repeaty, repeatz := CalculateOrbitalRepetitions(moons)
	fmt.Printf("Calculate coordinates done xrep=%d yrep=%d zrep=%d\n", repeatx, repeaty, repeatz)

	bigBang := CalculateBigBang(repeatx, repeaty, repeatz)

	assert.Equal(t, 17, repeatx)
	assert.Equal(t, 27, repeaty)
	assert.Equal(t, 43, repeatz)
	assert.Equal(t, 2772, bigBang)
}

func TestCalculateOrbitalRepetitions2(t *testing.T) {

	fmt.Printf("Calculate orbits\n")
	moons := LoadMoons(testInput2)
	repeatx, repeaty, repeatz := CalculateOrbitalRepetitions(moons)
	fmt.Printf("Calculate coordinates done xrep=%d yrep=%d zrep=%d\n", repeatx, repeaty, repeatz)

	bigBang := CalculateBigBang(repeatx, repeaty, repeatz)

	assert.Equal(t, 2027, repeatx)
	assert.Equal(t, 5897, repeaty)
	assert.Equal(t, 4701, repeatz)
	assert.Equal(t, 4686774924, bigBang)

}

func TestRunSimulation(t *testing.T) {

	moons := LoadMoons(testInput1)
	totalSystemEnergy := RunSimulation(moons, 30)
	assert.Equal(t, 0, totalSystemEnergy)
}
