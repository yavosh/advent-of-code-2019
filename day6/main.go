package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func orbitsToMap(orbitData string) map[string][]string {
	nodes := strings.Split(orbitData, "\n")
	orbits := make(map[string][]string, 0)
	for _, node := range nodes {
		if strings.TrimSpace(node) == "" {
			continue
		}
		orbitSides := strings.Split(node, ")")
		gravitron := orbitSides[0]
		satellite := orbitSides[1]
		if _, ok := orbits[gravitron]; !ok {
			orbits[gravitron] = make([]string, 0, 1)
		}
		orbits[gravitron] = append(orbits[gravitron], satellite)
	}

	return orbits
}

func reverseOrbits(orbits map[string][]string) map[string]string {
	orbitsReverse := make(map[string]string, 0)
	for planet, val := range orbits {
		for _, sattelite := range val {
			orbitsReverse[sattelite] = planet
		}
	}

	return orbitsReverse
}

func countMoveOrbits(orbitData string, from string, to string) int {

	orbitsReverse := reverseOrbits(orbitsToMap(orbitData))

	// move backwards from both nodes until we find a common
	// parent, when we do sum the total moves from both sides
	fromPlanet, fromPlanetOK := orbitsReverse[from]
	toPlanet, toPlanetOK := orbitsReverse[to]

	// Steps from origin
	fromPlanetPath := make(map[string]int)
	toPlanetPath := make(map[string]int)

	stepsFrom := 0
	stepsTo := 0

	for true {

		if fromPlanetOK {

			// check if other path has a common planet
			if _, ok := toPlanetPath[fromPlanet]; ok {
				fmt.Printf("*** Found common from juncture, fromPlanet=%s\n", fromPlanet)
				return toPlanetPath[fromPlanet] + stepsFrom
			}

			fromPlanetPath[fromPlanet] = stepsFrom
			fromPlanet, fromPlanetOK = orbitsReverse[fromPlanet]
			stepsFrom = stepsFrom + 1
		}

		if toPlanetOK {

			// check if other path has a common planet
			if _, ok := fromPlanetPath[toPlanet]; ok {
				fmt.Printf("*** Found common to juncture, fromPlanet=%s\n", fromPlanet)
				return fromPlanetPath[toPlanet] + stepsTo
			}

			toPlanetPath[toPlanet] = stepsTo
			toPlanet, toPlanetOK = orbitsReverse[toPlanet]
			stepsTo = stepsTo + 1
		}

		//fmt.Printf("fromPlanetPath: %v\n", fromPlanetPath)
		//fmt.Printf("toPlanetPath: %v\n", toPlanetPath)

		if !fromPlanetOK && !toPlanetOK {
			// Both have reached end
			break
		}

	}

	return 0
}

func countOrbits(orbitData string) int {

	orbitsReverse := reverseOrbits(orbitsToMap(orbitData))
	//fmt.Printf("orbitsReverse: %v\n", orbitsReverse)

	totalOrbits := 0
	for _, planet := range orbitsReverse {
		totalPlanetOrbits := 0
		totalPlanetOrbits++
		orbitedBy, ok := orbitsReverse[planet]
		for ok {
			orbitedBy, ok = orbitsReverse[orbitedBy]
			totalPlanetOrbits++
		}

		//fmt.Printf("satellite=%s planet=%s totalOrbits=%d\n", satellite, planet, totalPlanetOrbits)
		totalOrbits += totalPlanetOrbits
	}

	return totalOrbits
}

func main() {
	orbitData, err := ioutil.ReadFile("./orbit_data_d6_p1.txt")
	if err != nil {
		panic(err)
	}

	totalOrbits := countOrbits(string(orbitData))
	orbitsToSanta := countMoveOrbits(string(orbitData), "YOU", "SAN")
	fmt.Printf("total orbits: %d\n", totalOrbits)
	fmt.Printf("orbits to santa orbits: %d\n", orbitsToSanta)
}
