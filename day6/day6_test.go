package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var orbitData = `
COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

var orbitMoveData = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`

func TestOrbitData(t *testing.T) {
	totalOrbits := countOrbits(orbitData)
	fmt.Printf("totalOrbits=%d\n", totalOrbits)
	assert.Equal(t, 42, totalOrbits)
}

func TestMoveOrbit(t *testing.T) {
	moveOrbits := countMoveOrbits(orbitMoveData, "YOU", "SAN")
	fmt.Printf("totalOrbits=%d\n", moveOrbits)
	assert.Equal(t, 4, moveOrbits)
	t.Fail()
}
