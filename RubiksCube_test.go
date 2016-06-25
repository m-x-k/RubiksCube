package main

import "testing"

func TestRandomPiece(t *testing.T) {
	piece := randomPiece()
	assertPiece(piece, t)
}

func TestGetSide(t *testing.T) {
	side := getSide()
	if (len(side) != 3) {
		t.Error("Rubiks Cube Side should be 3 x 3")
	}
	var i, j int
	for i = 0; i < 3; i++  {
		for j = 0; j < 3; j++ {
			assertPiece(side[i][j], t)
		}
	}
}

func assertPiece(piece string, t *testing.T) {
	if (piece != WHITE && piece != RED && piece != BLUE &&
	piece != YELLOW && piece != ORANGE && piece != GREEN) {
		t.Error("Rubiks piece should be one of red, blue, green, white, orange, yellow: " + piece)
	}
}

