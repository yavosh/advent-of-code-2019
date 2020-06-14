package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToIntArrayut(t *testing.T) {
	out := stringToIntArray("123456789012")
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, out)
}

func TestImageToLayers(t *testing.T) {
	input := stringToIntArray("123456789012")
	layers := imageToLayers(input, 3, 2)
	fmt.Printf("layers: %d\n", layers)

	assert.Equal(t, 2, len(layers))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, layers[0])
	assert.Equal(t, []int{7, 8, 9, 0, 1, 2}, layers[1])
	t.Fail()
}

func TestLayerToRows(t *testing.T) {
	input := stringToIntArray("123456")
	rows := layerToRows(input, 3)
	fmt.Printf("rows: %d\n", rows)
	assert.Equal(t, 2, len(rows))
	assert.Equal(t, []int{1, 2, 3}, rows[0])
	assert.Equal(t, []int{4, 5, 6}, rows[1])
	t.Fail()
}

func TestComposePixels(t *testing.T) {
	input := stringToIntArray("0222112222120000")
	layers := imageToLayers(input, 2, 2)

	finalLayer := composeLayers(layers)
	fmt.Printf("finalLayer: %v\n", finalLayer)
	printLayer(finalLayer, 2)
	t.Fail()
}

func TestDay8P1(t *testing.T) {
	imagePixels, err := ioutil.ReadFile("./image_d8.txt")
	if err != nil {
		panic(err)
	}

	input := stringToIntArray(string(imagePixels))
	layers := imageToLayers(input, 25, 6)

	leastZeros := 25*6 + 1
	leastZerosIndex := -1

	for index, layer := range layers {
		count := countNumbersInLayer(layer, 0)
		fmt.Printf("count=%d \n", count)

		if count < leastZeros {
			leastZeros = count
			leastZerosIndex = index
		}
	}

	fmt.Printf("leastZerosIndex=%d leastZeros=%d\n", leastZerosIndex, leastZeros)

	countOnes := countNumbersInLayer(layers[leastZerosIndex], 1)
	countTwos := countNumbersInLayer(layers[leastZerosIndex], 2)

	fmt.Printf("countOnes=%d countTwos=%d result=%d\n", countOnes, countTwos, countOnes*countTwos)
	t.Fail()
}

