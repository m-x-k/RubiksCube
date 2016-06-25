package main

import "fmt"
import "math/rand"
import "strconv"

const (
	WHITE  = "w"
	BLUE   = "b"
	RED    = "r"
	GREEN  = "g"
	YELLOW = "y"
	ORANGE = "o"
)

func main() {
	fmt.Println("Starting Rubiks Cube")
	var a = getSide()
	printSide(a, 1)
	printSide(a, 2)
	printSide(a, 3)
	printSide(a, 4)
	printSide(a, 5)
	printSide(a, 6)
	fmt.Println("Ending Rubiks Cube")
}

func randomPiece() string {
	var r = rand.Intn(5)
	if (r == 0) {
		return WHITE
	}
	if (r == 1) {
		return BLUE
	}
	if (r == 2) {
		return RED
	}
	if (r == 3) {
		return GREEN
	}
	if (r == 4) {
		return YELLOW
	}
	return ORANGE
}

func getSide() [3][3]string {
	var a = [3][3]string{
		{randomPiece(), randomPiece(), randomPiece()},
		{randomPiece(), randomPiece(), randomPiece()},
		{randomPiece(), randomPiece(), randomPiece()},
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