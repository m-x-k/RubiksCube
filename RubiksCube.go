package main

import "fmt"
import "math/rand"
import "strconv"
import "errors"

const (
	WHITE  = "w"
	BLUE   = "b"
	RED    = "r"
	GREEN  = "g"
	YELLOW = "y"
	ORANGE = "o"
)

var m = map[string]int{}

func main() {
	fmt.Println("Starting Rubiks Cube")
	NewRubiksCube()

	printSide(getSide(), 1)
	printSide(getSide(), 2)
	printSide(getSide(), 3)
	printSide(getSide(), 4)
	printSide(getSide(), 5)
	printSide(getSide(), 6)

	ensureNoBlocksRemaining(m)

	fmt.Println("Ending Rubiks Cube")
}

func NewRubiksCube() {
	m[WHITE]  = 9
	m[BLUE]   = 9
	m[RED]    = 9
	m[GREEN]  = 9
	m[YELLOW] = 9
	m[ORANGE] = 9
}

func ensureNoBlocksRemaining(m map[string]int) {
	for _, value := range m {
		if (value != 0) {
			panic(errors.New("Something went wrong all blocks are not used"))
		}
	}
}

func randomBlock() string {
	var rand = rand.Intn(6)
	var block = intToBlock(rand)
	if (m[block] >= 1) {
		m[block] -= 1
		return block
	}
	if (m[block] < 0) {
		panic(errors.New("Too many blocks used"))
	}

	return randomBlock()
}

func intToBlock(i int) string {
	if (i == 0) {
		return WHITE
	}
	if (i == 1) {
		return BLUE
	}
	if (i == 2) {
		return RED
	}
	if (i == 3) {
		return GREEN
	}
	if (i == 4) {
		return YELLOW
	}
	if (i == 5) {
		return ORANGE
	}
	panic(errors.New("Unknown block"))
}

func getSide() [3][3]string {
	var a = [3][3]string{
		{randomBlock(), randomBlock(), randomBlock()},
		{randomBlock(), randomBlock(), randomBlock()},
		{randomBlock(), randomBlock(), randomBlock()},
	}
	return a
}

func printSide(a [3][3]string, s int) {
	var i, j int
	fmt.Println("     SIDE" + strconv.Itoa(s) + "    ")
	fmt.Println(" -------------")
	for  i = 0; i < len(a); i++ {
		var line string
		for j = 0; j < len(a); j++ {
			line = line + " | " + a[i][j]
		}
		fmt.Println(line + " |")
	}
	fmt.Println(" -------------")
}