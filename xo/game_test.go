package main

import "testing"

func TestCreatingGame(t *testing.T) {
	gameSettings := GameStart{1, 2, 3}
	game := CreateGame(gameSettings)
	if game.BoardSize != 3 {
		t.Errorf("Creating game with board size = %d; want 3", game.BoardSize)
	}
}

func TestCreatingGameBoardSizeOutOfRange(t *testing.T) {
	gameSettings := GameStart{1, 2, 1}
	gameSettings.Validate()
	game := CreateGame(gameSettings)
	if game.BoardSize != 3 {
		t.Errorf("Creating game with board size = %d; want 3", game.BoardSize)
	}
}

func TestCreatingGameWrongPlayerId(t *testing.T) {
	gameSettings := GameStart{0, 2, 1}
	if gameSettings.Validate() {
		t.Errorf("Creating game with playerId = 0 shouldn't pass validation")
	}
}

func TestMakingFirstMove(t *testing.T) {
	gameSettings := GameStart{1, 2, 3}
	game := CreateGame(gameSettings)
	game.Move(0, 0, game.NextPlayer)
	if game.GameBoard[0][0] != "X" {
		t.Errorf("(0,0) should contain an X mark")
	}
}

func TestMakingSecondMove(t *testing.T) {
	gameSettings := GameStart{1, 2, 3}
	game := CreateGame(gameSettings)
	game.Move(0, 0, game.NextPlayer)
	game.Move(0, 1, game.NextPlayer)
	if game.GameBoard[0][1] != "O" {
		t.Errorf("(0,1) should contain an O mark")
	}
}

func TestMakingMoveByWrongPlayer(t *testing.T) {
	gameSettings := GameStart{1, 2, 3}
	game := CreateGame(gameSettings)
	firstPlayer := game.NextPlayer
	game.Move(0, 0, firstPlayer)
	game.Move(0, 1, firstPlayer) // same player twice in a row
	if game.GameBoard[0][1] != "" {
		t.Errorf("The move shouldn't have executed")
	}
}

func TestDetectingVictoryRow(t *testing.T) {
	gameSettings := GameStart{1, 2, 3}
	game := CreateGame(gameSettings)
	firstPlayer := game.NextPlayer
	game.Move(0, 0, game.NextPlayer) // p1
	game.Move(1, 1, game.NextPlayer) // p2
	game.Move(0, 1, game.NextPlayer) // p1
	game.Move(1, 2, game.NextPlayer) // p2
	game.Move(0, 2, game.NextPlayer) // p1 - victory
	if !game.Completed {
		t.Errorf("Game should have been completed")
	}
	if game.Winner != firstPlayer {
		t.Errorf("Game should have been won by the player making the first move")
	}
}

func TestDetectingVictoryColumn(t *testing.T) {
	gameSettings := GameStart{1, 2, 3}
	game := CreateGame(gameSettings)
	firstPlayer := game.NextPlayer
	game.Move(0, 0, game.NextPlayer) // p1
	game.Move(1, 1, game.NextPlayer) // p2
	game.Move(1, 0, game.NextPlayer) // p1
	game.Move(1, 2, game.NextPlayer) // p2
	game.Move(2, 0, game.NextPlayer) // p1 - victory
	if !game.Completed {
		t.Errorf("Game should have been completed")
	}
	if game.Winner != firstPlayer {
		t.Errorf("Game should have been won by the player making the first move")
	}
}

func TestDetectingVictoryDiagonalLR(t *testing.T) {
	gameSettings := GameStart{1, 2, 3}
	game := CreateGame(gameSettings)
	firstPlayer := game.NextPlayer
	game.Move(0, 0, game.NextPlayer) // p1
	game.Move(1, 0, game.NextPlayer) // p2
	game.Move(1, 1, game.NextPlayer) // p1
	game.Move(1, 2, game.NextPlayer) // p2
	game.Move(2, 2, game.NextPlayer) // p1 - victory
	if !game.Completed {
		t.Errorf("Game should have been completed")
	}
	if game.Winner != firstPlayer {
		t.Errorf("Game should have been won by the player making the first move")
	}
}

func TestDetectingVictoryDiagonalRL(t *testing.T) {
	gameSettings := GameStart{1, 2, 3}
	game := CreateGame(gameSettings)
	firstPlayer := game.NextPlayer
	game.Move(0, 2, game.NextPlayer) // p1
	game.Move(1, 0, game.NextPlayer) // p2
	game.Move(1, 1, game.NextPlayer) // p1
	game.Move(1, 2, game.NextPlayer) // p2
	game.Move(2, 0, game.NextPlayer) // p1 - victory
	if !game.Completed {
		t.Errorf("Game should have been completed")
	}
	if game.Winner != firstPlayer {
		t.Errorf("Game should have been won by the player making the first move")
	}
}

func TestDetectingDraw(t *testing.T) {
	gameSettings := GameStart{1, 2, 3}
	game := CreateGame(gameSettings)
	game.Move(0, 1, game.NextPlayer) // p1
	game.Move(0, 0, game.NextPlayer) // p2
	game.Move(1, 0, game.NextPlayer) // p1
	game.Move(0, 2, game.NextPlayer) // p2
	game.Move(1, 1, game.NextPlayer) // p1
	game.Move(1, 2, game.NextPlayer) // p2
	game.Move(2, 0, game.NextPlayer) // p1
	game.Move(2, 1, game.NextPlayer) // p2
	game.Move(2, 2, game.NextPlayer) // p1
	if !game.Completed {
		t.Errorf("Game should have been completed")
	}
	if game.Winner != 0 {
		t.Errorf("Game should have ended with a draw")
	}
}
