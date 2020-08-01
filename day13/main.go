package main

import (
	"fmt"
	"io/ioutil"

	"github.com/yavosh/advent-of-code-2019/util"

	"github.com/yavosh/advent-of-code-2019/computer"
)

const (
	TileTypeEmpty  = 0
	TileTypeWall   = 1
	TileTypeBlock  = 2
	TileTypePaddle = 3
	TileTypeBall   = 4

	JoystickDirectionLeft  = -1
	JoystickDirectionRight = 1
	JoystickDirectionNone  = 0
)

var (
	valueMap = map[int16]string{
		TileTypeEmpty:  " ",
		TileTypeWall:   "*",
		TileTypeBlock:  "-",
		TileTypePaddle: "=",
		TileTypeBall:   "@",
	}
)

type GameState struct {
	gameLoops int
	score     int
	sizeX     int
	sizeY     int
	ball      *Ball
	joystick  *Joystick
	paddle    *Paddle
	screen    []int16
}

type Joystick struct {
	direction int64
}

type Ball struct {
	position  Vector
	direction Vector
}

type Vector struct {
	x, y int64
}

func (v Vector) String() string {
	return fmt.Sprintf("%d,%d", v.x+1, v.y+1)
}

func (game *GameState) PredictCollide() Vector {

	ballCopy := Ball{
		position: Vector{
			x: game.ball.position.x,
			y: game.ball.position.y,
		},
		direction: Vector{
			x: game.ball.direction.x,
			y: game.ball.direction.y,
		},
	}

	var steps = 0
	for {

		ballCopy.position.x = ballCopy.position.x - game.ball.direction.x
		ballCopy.position.y = ballCopy.position.y - game.ball.direction.y
		if ballCopy.position.y == game.paddle.position.y { // paddle
			return ballCopy.position
		}

		steps++
		if steps > 10 {
			break
		}
	}

	return Vector{ballCopy.position.x, ballCopy.position.y}
}

func (ball *Ball) VDirection() string {
	if ball.direction.y > 0 {
		return "U"
	}

	if ball.direction.y < 0 {
		return "D"
	}

	return "I"
}

func (ball *Ball) HDirection() string {
	if ball.direction.x > 0 {
		return "R"
	}

	if ball.direction.x < 0 {
		return "L"
	}

	return "I"
}

func (ball *Ball) Position(x int64, y int64) {
	//fmt.Printf("> Move ball %v[%v] -> %d,%d\n", ball.position, ball.direction, x, y)

	if ball.position.x != -1 {
		ball.direction.x = ball.position.x - x
	}

	if ball.position.y != -1 {
		ball.direction.y = ball.position.y - y
	}

	ball.position.x = x
	ball.position.y = y
}

type Paddle struct {
	position Vector
}

func (ball *Ball) Status() string {
	return fmt.Sprintf("[ball v=%s h=%s %d,%d]", ball.VDirection(), ball.HDirection(), ball.position.x+1, ball.position.y+1)
}

func (paddle *Paddle) Status() string {
	return fmt.Sprintf("[pad %d,%d]", paddle.position.x+1, paddle.position.y+1)
}

func (joystick *Joystick) Status() string {

	direction := "I"
	if joystick.direction > 0 {
		direction = "R"
	}

	if joystick.direction < 0 {
		direction = "L"
	}

	return fmt.Sprintf("[j %s]", direction)
}

func (paddle *Paddle) Position(x int64, y int64) {
	//fmt.Printf("> Move paddle %s -> %d,%d\n", paddle.position, x+1, y+1)
	paddle.position.x = x
	paddle.position.y = y
}

func (game *GameState) setInput() {
	//fmt.Printf(">>> ball:%v y=%d\n", game.ball, game.ball.direction.y)

	// if ball is moving up move towards middle
	if game.ball.direction.y > 0 {
		//fmt.Printf("Ball moving up \n")
		// follow ball future postion
		if game.ball.position.x-game.ball.direction.x > game.paddle.position.x {
			game.joystick.direction = JoystickDirectionRight
		}

		// follow ball future postion
		if game.ball.position.x-game.ball.direction.x < game.paddle.position.x {
			game.joystick.direction = JoystickDirectionLeft
		}

		if game.ball.position.x-game.ball.direction.x == game.paddle.position.x {
			game.joystick.direction = JoystickDirectionNone
		}
	}

	// if ball is moving downward move towards point of losing
	if game.ball.direction.y < 0 {
		//fmt.Printf("Ball moving down \n")
		collidePoint := game.PredictCollide()
		deltaX := util.AbsInt64(collidePoint.x - game.paddle.position.x)
		// fmt.Printf("PredictCollide point:[%s] deltaX=%d\n", collidePoint, deltaX)

		// Paddle has 1 position tolerance
		if deltaX <= 1 {
			game.joystick.direction = JoystickDirectionNone
		}

		if deltaX > 1 {
			if collidePoint.x > game.paddle.position.x {
				game.joystick.direction = JoystickDirectionRight
			}

			// one move predict -- + game.ball.direction.x
			if collidePoint.x < game.paddle.position.x {
				game.joystick.direction = JoystickDirectionLeft
			}

			if collidePoint.x == game.paddle.position.x {
				game.joystick.direction = JoystickDirectionNone
			}
		}

	}
}

func (game *GameState) actions() {
	game.gameLoops++

	if game.paddle.position.x == game.ball.position.x && game.paddle.position.y == game.ball.position.y {
		//fmt.Printf("Collision %d\b", game.gameLoops)
	}
}

func (game *GameState) Input() int64 {
	//time.Sleep(5 * time.Millisecond)

	game.actions()
	//game.drawScreen()
	//game.drawScreenStatus()
	return game.joystick.direction
}

func mainloop(arcadeCode string, input func() int64, output chan int64, exit chan int64) {
	fmt.Printf("start main loop\n")
	memory := computer.LoadInstructionsWithMemoryAlloc(arcadeCode, 8192)
	memory[0] = 2
	computer.RunWithChannelsCallback(memory, "arcanoid", input, output, exit)
	close(exit)
}

func (game *GameState) drawScreenBuffer() {
	fmt.Printf("Screen buffer:\n")
	fmt.Println()
	fmt.Printf("%+v\n", game.screen)
	fmt.Println()

	fmt.Printf("Bitmap:\n")
	rows := len(game.screen) / game.sizeX
	for r := 0; r < rows; r++ {
		line := game.screen[r*game.sizeX : (r+1)*game.sizeX]
		fmt.Printf("%03d -- %v\n", r, line)
	}
}

// screen is square
func (game *GameState) drawScreen() {
	fmt.Println()
	fmt.Println("+-1      +-10      +-20      +-30      +40")
	fmt.Println("|        |         |         |         |")
	for y := 0; y < game.sizeY; y++ {
		for x := 0; x < game.sizeX; x++ {
			value := game.screen[y*game.sizeX+x]
			fmt.Print(valueMap[value])
		}
		fmt.Printf(" -- %03d ", y+1)
		fmt.Println()
	}
	fmt.Println("|        |         |         |         |")
	fmt.Println("+-1      +-10      +-20      +-30      +40")

}

func (game *GameState) drawScreenStatus() {
	fmt.Println()
	fmt.Printf(">> score [%06d] [%03d] %s %s %s\n",
		game.score, game.gameLoops, game.ball.Status(), game.joystick.Status(), game.paddle.Status())
	fmt.Println()
}

func main() {
	arcadeCode, err := ioutil.ReadFile("./arcade_code_d13.txt")
	if err != nil {
		panic(err)
	}

	var tilesByType = map[int16]int{
		TileTypeEmpty:  0,
		TileTypeWall:   0,
		TileTypeBlock:  0,
		TileTypePaddle: 0,
		TileTypeBall:   0,
	}

	// var score = 0
	//var screenSizeX = 40
	//var screenSizeY = 24

	var game = GameState{
		score: 0,
		sizeX: 40,
		sizeY: 24,
		// rows * pixels screen is square
		screen: make([]int16, 40*24),
		ball: &Ball{
			position: Vector{
				x: -1,
				y: -1,
			},
		},
		paddle: &Paddle{},
		joystick: &Joystick{
			direction: 0,
		},
	}

	var output = make(chan int64, 0)
	var exit = make(chan int64)

	go mainloop(string(arcadeCode), game.Input, output, exit)

	var done = false
	var x, y, value int64

	for !done {

		select {
		case x = <-output:
			y = <-output
			value = <-output
			tilesByType[int16(value)]++
			if x == -1 && y == 0 {
				game.score = int(value)
				//fmt.Printf("*** output: x=%d y=%d score=%d \n", x, y, value)
				continue
			}

			positionInBuffer := int(y)*game.sizeX + int(x)
			game.screen[positionInBuffer] = int16(value)
			//fmt.Printf("*** output: x=%d y=%d tile=%d pos=%d added=%d\n", x, y, value, positionInBuffer, game.screen[positionInBuffer])

			if value == TileTypeBall {
				game.ball.Position(x, y)
				//ball.Position(x, y)
				//fmt.Printf("*** ball position is now x=%d y=%d positionInBuffer=%d\n", x, y, positionInBuffer)
			}

			if value == TileTypePaddle {
				game.paddle.Position(x, y)
				//fmt.Printf("*** paddle position is now x=%d y=%d positionInBuffer=%d\n", x, y, positionInBuffer)
			}

			if value == TileTypeBall || value == TileTypePaddle {
				game.setInput()

			}

		case out := <-exit:
			fmt.Printf("*** Exit: %v\n", out)
			//fmt.Printf("*** tilesByType: %+v\n", tilesByType)
			fmt.Println()
			fmt.Println()
			// game.drawScreen()
			game.drawScreenStatus()
			done = true
			break
		default:
			//fmt.Printf("*** noop \n")
			// input <- 0
		}
	}
}
