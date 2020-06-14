package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

const (
	colorBlack       = 0
	colorWhite       = 1
	colorTransparent = 2
)

func stringToIntArray(input string) []int {
	out := make([]int, len(input))

	for pos, r := range input {
		out[pos], _ = strconv.Atoi(string(r))
	}
	return out
}

func imageToLayers(input []int, width int, height int) [][]int {
	pixelsPerLayer := width * height
	layersCount := len(input) / pixelsPerLayer
	layers := make([][]int, layersCount)
	for layer := 0; layer < layersCount; layer++ {
		layers[layer] = input[layer*pixelsPerLayer : (layer+1)*pixelsPerLayer]
	}
	return layers
}

func layerToRows(layer []int, width int) [][]int {
	rowsCount := len(layer) / width
	rows := make([][]int, rowsCount)
	for row := 0; row < rowsCount; row++ {
		rows[row] = layer[row*width : (row+1)*width]
	}
	return rows
}

func composeLayers(layers [][]int) []int {
	composedLayer := make([]int, len(layers[0]))
	for pixelIndex := range layers[0] {
		for layerIndex := range layers {
			currentPixel := layers[layerIndex][pixelIndex]
			if currentPixel != colorTransparent {
				composedLayer[pixelIndex] = currentPixel
				break
			}
		}
	}

	return composedLayer
}

func printLayer(layer []int, width int) {
	rows := layerToRows(layer, width)
	for _, row := range rows {
		for _, pixelValue := range row {
			if pixelValue == colorBlack {
				fmt.Printf("_")
			} else if pixelValue == colorWhite {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func countNumbersInLayer(layer []int, number int) int {
	acc := 0
	for _, value := range layer {
		if value == number {
			acc++
		}
	}
	return acc
}

func main() {
	imagePixels, err := ioutil.ReadFile("./image_d8.txt")
	if err != nil {
		panic(err)
	}

	input := stringToIntArray(string(imagePixels))
	layers := imageToLayers(input, 25, 6)

	composed := composeLayers(layers)
	printLayer(composed, 25)
	//HZCZU
}
