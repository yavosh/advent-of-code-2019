package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"github.com/yavosh/advent-of-code-2019/computer"
)

func thrusters(code string, input []int) int {
	codeInput := []int{0, 0}
	for thrusterNumber, thrusterInput := range input {
		codeInput[0] = thrusterInput
		fmt.Printf("int: thruster:%d %v\n", thrusterNumber+1, codeInput)
		_, output := computer.Run(computer.LoadInstructions(code), codeInput)

		fmt.Printf("out: %v\n", output)
		codeInput[1] = output[0]
	}
	// last output
	return codeInput[1]
}

func thrustersParallel(code string, input []int) int {

	type Thruster struct {
		mu        *sync.Mutex
		name      string
		codeValue string
		result    int
		input     chan int
		output    chan int
		exit      chan int
		done      bool
	}

	runThruster := func(t *Thruster, wg *sync.WaitGroup) int {
		defer wg.Done()
		memory := computer.LoadInstructions(t.codeValue)
		_, lastOutput := computer.RunWithChannels(memory, t.name, t.input, t.output, t.exit)
		close(t.input)
		return lastOutput
	}

	thrusterE := &Thruster{
		name:      "E",
		codeValue: code,
		input:     make(chan int),
		output:    make(chan int),
		exit:      make(chan int),
	}

	thrusterD := &Thruster{
		name:      "D",
		codeValue: code,
		input:     make(chan int),
		output:    make(chan int),
		exit:      make(chan int)}

	thrusterC := &Thruster{
		name:      "C",
		codeValue: code,
		input:     make(chan int),
		output:    make(chan int),
		exit:      make(chan int),
	}

	thrusterB := &Thruster{
		name:      "B",
		codeValue: code,
		input:     make(chan int),
		output:    make(chan int),
		exit:      make(chan int),
	}

	thrusterA := &Thruster{
		mu:        &sync.Mutex{},
		name:      "A",
		codeValue: code,
		input:     make(chan int, 5),
		output:    make(chan int),
		exit:      make(chan int),
	}

	var wg sync.WaitGroup

	wg.Add(5)
	go runThruster(thrusterA, &wg)
	go runThruster(thrusterB, &wg)
	go runThruster(thrusterC, &wg)
	go runThruster(thrusterD, &wg)
	go runThruster(thrusterE, &wg)

	fmt.Printf("# Starting thrusters input:%v\n", input)

	thrusterA.input <- input[0]
	thrusterA.input <- 0
	thrusterB.input <- input[1]
	thrusterC.input <- input[2]
	thrusterD.input <- input[3]
	thrusterE.input <- input[4]

	done := false
	for !done {

		select {
		case msgA := <-thrusterA.output:
			thrusterA.mu.Lock()
			if !thrusterA.done {
				//fmt.Println("received msgA sending to B ", msgA)
				thrusterB.input <- msgA
			} else {
				fmt.Println("received msgA after a was done", msgA)
			}
			thrusterA.mu.Unlock()
		case msgB := <-thrusterB.output:
			//fmt.Println("received msgB sending to B", msgB)
			thrusterC.input <- msgB
		case msgC := <-thrusterC.output:
			//fmt.Println("received msgC sending to D", msgC)
			thrusterD.input <- msgC
		case msgD := <-thrusterD.output:
			//fmt.Println("received msgD sending to E", msgD)
			thrusterE.input <- msgD
		case msgE := <-thrusterE.output:
			// sync a, sometime thruster e is faster to read thrusterA.done
			thrusterA.mu.Lock()
			if !thrusterA.done {
				//fmt.Println("received msgE sending to A", msgE)
				thrusterA.input <- msgE
			} else {
				fmt.Printf("Thruster E last result %d\n", msgE)
			}
			thrusterA.mu.Unlock()
		case exitA := <-thrusterA.exit:
			// sync a, sometime thruster e is faster to read thrusterA.done
			thrusterA.mu.Lock()
			//fmt.Println("$ Thruster A done", exitA)
			thrusterA.done = true
			thrusterA.result = exitA
			thrusterA.mu.Unlock()
		case exitB := <-thrusterB.exit:
			//fmt.Println("$ Thruster B done", exitB)
			thrusterB.done = true
			thrusterB.result = exitB
		case exitC := <-thrusterC.exit:
			//fmt.Println("$ Thruster C done", exitC)
			thrusterC.done = true
			thrusterC.result = exitC
		case exitD := <-thrusterD.exit:
			//fmt.Println("$ Thruster D done", exitD)
			thrusterD.done = true
			thrusterD.result = exitD
		case exitE := <-thrusterE.exit:
			//fmt.Println("$ Thruster E done", exitE)
			thrusterE.done = true
			thrusterE.result = exitE
			done = true
		case <-time.After(5 * time.Second):
			fmt.Printf("timeout i/o")
			panic("timeout i/o")
		}
	}

	wg.Wait()
	fmt.Printf("Results a=%d b=%d c=%d d=%d e=%d\n",
		thrusterA.result, thrusterB.result, thrusterC.result, thrusterD.result, thrusterE.result,
	)
	return thrusterE.result
}

func singlePass() {
	thrusterCode, err := ioutil.ReadFile("./thruster_code_d7.txt")
	if err != nil {
		panic(err)
	}

	maxValue := 0
	var maxInput []int

	for _, input := range permutations {
		out := thrusters(string(thrusterCode), input)

		if out > maxValue {
			maxValue = out
			maxInput = input
		}
	}

	fmt.Printf("max output: %d input: %v", maxValue, maxInput)
}

func parallel() {
	thrusterCode, err := ioutil.ReadFile("./thruster_code_d7.txt")
	if err != nil {
		panic(err)
	}

	maxValue := 0
	var maxInput []int

	for _, input := range permutations5to9 {
		fmt.Printf("input=%v\n", input)
		out := thrustersParallel(string(thrusterCode), input)

		if out > maxValue {
			maxValue = out
			maxInput = input
		}
	}

	fmt.Printf("max output: %d input: %v", maxValue, maxInput)
}

func main() {
	parallel()
}
