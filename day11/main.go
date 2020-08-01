package main

import (
	"fmt"
	"time"

	"github.com/yavosh/advent-of-code-2019/computer"
)

const (
	DirectionUp               = 0
	DirectionRight            = 90
	DirectionDown             = 180
	DirectionLeft             = 270
	ColourBlack    ColourType = 0
	ColourWhite    ColourType = 1
)

var (
	program = `3,8,1005,8,319,1106,0,11,0,0,0,104,1,104,0,3,8,1002,8,-1,10,101,1,10,10,4,10,108,1,8,10,4,10,1001,8,0,28,2,1008,7,10,2,4,17,10,3,8,102,-1,8,10,101,1,10,10,4,10,1008,8,0,10,4,10,1002,8,1,59,3,8,1002,8,-1,10,101,1,10,10,4,10,1008,8,0,10,4,10,1001,8,0,81,1006,0,24,3,8,1002,8,-1,10,101,1,10,10,4,10,108,0,8,10,4,10,102,1,8,105,2,6,13,10,1006,0,5,3,8,1002,8,-1,10,101,1,10,10,4,10,108,0,8,10,4,10,1002,8,1,134,2,1007,0,10,2,1102,20,10,2,1106,4,10,1,3,1,10,3,8,102,-1,8,10,101,1,10,10,4,10,108,1,8,10,4,10,1002,8,1,172,3,8,1002,8,-1,10,1001,10,1,10,4,10,108,1,8,10,4,10,101,0,8,194,1,103,7,10,1006,0,3,1,4,0,10,3,8,1002,8,-1,10,1001,10,1,10,4,10,1008,8,1,10,4,10,101,0,8,228,2,109,0,10,1,101,17,10,1006,0,79,3,8,1002,8,-1,10,1001,10,1,10,4,10,108,0,8,10,4,10,1002,8,1,260,2,1008,16,10,1,1105,20,10,1,3,17,10,3,8,1002,8,-1,10,1001,10,1,10,4,10,1008,8,1,10,4,10,1002,8,1,295,1,1002,16,10,101,1,9,9,1007,9,1081,10,1005,10,15,99,109,641,104,0,104,1,21101,387365733012,0,1,21102,1,336,0,1105,1,440,21102,937263735552,1,1,21101,0,347,0,1106,0,440,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,21102,3451034715,1,1,21101,0,394,0,1105,1,440,21102,3224595675,1,1,21101,0,405,0,1106,0,440,3,10,104,0,104,0,3,10,104,0,104,0,21101,0,838337454440,1,21102,428,1,0,1105,1,440,21101,0,825460798308,1,21101,439,0,0,1105,1,440,99,109,2,22101,0,-1,1,21102,1,40,2,21101,0,471,3,21101,461,0,0,1106,0,504,109,-2,2106,0,0,0,1,0,0,1,109,2,3,10,204,-1,1001,466,467,482,4,0,1001,466,1,466,108,4,466,10,1006,10,498,1102,1,0,466,109,-2,2105,1,0,0,109,4,2101,0,-1,503,1207,-3,0,10,1006,10,521,21101,0,0,-3,21202,-3,1,1,22102,1,-2,2,21101,1,0,3,21102,540,1,0,1105,1,545,109,-4,2105,1,0,109,5,1207,-3,1,10,1006,10,568,2207,-4,-2,10,1006,10,568,22102,1,-4,-4,1106,0,636,22102,1,-4,1,21201,-3,-1,2,21202,-2,2,3,21102,587,1,0,1105,1,545,21201,1,0,-4,21101,0,1,-1,2207,-4,-2,10,1006,10,606,21102,0,1,-1,22202,-2,-1,-2,2107,0,-3,10,1006,10,628,22102,1,-1,1,21102,1,628,0,105,1,503,21202,-2,-1,-2,22201,-4,-2,-4,109,-5,2106,0,0`
)

// ColourType .
type ColourType int64

// RobotPosition .
type RobotPosition struct {
	x int64
	y int64
}

// Robot .
type Robot struct {
	pos         RobotPosition
	direction   int64
	panel       map[string]ColourType
	grid        [][]int
	gridSize    int64
	leftCorner  RobotPosition
	rightCorner RobotPosition
}

// Position .
func (r *Robot) Position() string {
	return fmt.Sprintf("%d/%d", r.pos.x, r.pos.y)
}

// Paint .
func (r *Robot) Paint(colour ColourType) {
	fmt.Printf("Paint pos:%s colour:%d\n", r.Position(), colour)
	r.panel[r.Position()] = colour
	r.grid[r.pos.x+50][r.pos.y+50] = int(colour)

}

// ReadCamera .
func (r *Robot) ReadCamera() ColourType {
	p := r.Position()
	if _, ok := r.panel[p]; !ok {
		// default
		r.panel[p] = ColourBlack
		fmt.Printf("ReadCamera default pos:%s colour:%d\n", p, r.panel[p])
		return ColourBlack
	}

	fmt.Printf("ReadCamera pos:%s colour:%d\n", p, r.panel[p])
	return r.panel[p]
}

// Turn .
func (r *Robot) Turn(direction int) *Robot {

	before := fmt.Sprintf("direction:%d", r.direction)
	if direction == 1 {
		r.direction += 90
	}

	if direction == 0 {
		r.direction -= 90
	}

	if r.direction < 0 {
		r.direction = r.direction + 360
	}

	if r.direction >= 360 {
		r.direction = r.direction - 360
	}

	after := fmt.Sprintf("direction:%d", r.direction)
	fmt.Printf("Turn: direction=%d before=%s after=%s\n", direction, before, after)

	return r
}

// Move .
func (r *Robot) Move() *Robot {
	before := r.Position()
	if r.direction == DirectionUp {
		r.pos.y--
	}

	if r.direction == DirectionDown {
		r.pos.y++
	}

	if r.direction == DirectionLeft {
		r.pos.x--
	}

	if r.direction == DirectionRight {
		r.pos.x++
	}

	if r.leftCorner.x > r.gridSize+r.pos.x {
		r.leftCorner.x = r.gridSize + r.pos.x
	}

	if r.leftCorner.y > r.gridSize+r.pos.y {
		r.leftCorner.y = r.gridSize + r.pos.y
	}

	if r.rightCorner.x < r.gridSize+r.pos.x {
		r.rightCorner.x = r.gridSize + r.pos.x
	}

	if r.rightCorner.y < r.gridSize+r.pos.y {
		r.rightCorner.y = r.gridSize + r.pos.y
	}

	after := r.Position()
	fmt.Printf("Move direction=%d before:%s after:%s\n", r.direction, before, after)
	return r
}

func runRobot() {
	var input = make(chan int64, 10)
	var output = make(chan int64, 10)
	var exit = make(chan int64)

	grid := make([][]int, 100)
	for x := 0; x < 100; x++ {
		grid[x] = make([]int, 100)
	}

	robot := Robot{
		pos:       RobotPosition{x: 0, y: 0},
		direction: DirectionUp,
		panel: map[string]ColourType{
			"0/0": ColourWhite,
		},
		grid:     grid,
		gridSize: int64(len(grid) / 2),
	}

	go func() {
		// memory := computer.LoadInstructions(program)
		memory := computer.LoadInstructionsWithMemoryAlloc(string(program), 8192)
		_, lastOut := computer.
			RunWithChannels(memory, "ms-paint", input, output, exit)

		fmt.Printf("*** Done ... %v \n", lastOut)
	}()

	steps := 0
	done := false
	for !done {

		select {
		case out := <-exit:
			fmt.Printf("*** Exit1: %v\n", out)
			done = true
			break
		default:
			fmt.Printf("*** Step steps1=%d\n", steps)
		}

		steps++
		colour := int64(robot.ReadCamera())

		input <- colour
		time.Sleep(10)

		newColour := <-output
		moveDirection := <-output

		fmt.Printf("Input: newColour=%d moveDirection=%d\n", newColour, moveDirection)

		robot.Paint(ColourType(newColour))
		robot.Turn(int(moveDirection))
		robot.Move()

		select {
		case out := <-exit:
			fmt.Printf("*** Exit1: %v\n", out)
			done = true
			break
		default:
			fmt.Printf("*** Step steps1=%d\n", steps)
		}

		if steps > 50000 {
			fmt.Printf("*** Infinite loop\n")
			break
		}
	}

	//fmt.Printf("Robot work: %v\n", robot.panel)
	fmt.Printf("Robot panels: %d\n", len(robot.panel))
	fmt.Printf("Robot steps: %d\n", steps)

	for x := 0; x < len(robot.grid); x++ {
		for y := 0; y < len(robot.grid); y++ {
			if robot.grid[x][y] == int(ColourBlack) {
				fmt.Printf(".")
				continue
			}

			if robot.grid[x][y] == int(ColourWhite) {
				fmt.Printf("#")
				continue
			}

			fmt.Printf("x")

		}

		fmt.Println()
	}

	fmt.Printf("Left top bound: %v\n", robot.leftCorner)
	fmt.Printf("Right bottom bound: %v\n", robot.rightCorner)

}

func main() {
	runRobot()
}
