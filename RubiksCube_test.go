package main

import "testing"

func TestRandomBlock(t *testing.T) {
	NewRubiksCube()
	piece := randomBlock()
	assertBlock(piece, t)
}

func TestGetSide(t *testing.T) {
	NewRubiksCube()
	side := getSide()
	if (len(side) != 3) {
		t.Error("Rubiks Cube Side should be 3 x 3")
	}
	var i, j int
	for i = 0; i < 3; i++  {
		for j = 0; j < 3; j++ {
			assertBlock(side[i][j], t)
		}
	}
}

func assertBlock(block string, t *testing.T) {
	if (block != WHITE && block != RED && block != BLUE && block != YELLOW && block != ORANGE && block != GREEN) {
		t.Error("Rubiks block should be one of red, blue, green, white, orange, yellow: " + block)
	}
}
