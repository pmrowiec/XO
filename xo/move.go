package main

import "fmt"

func MakeMove(x int, y int, game *Game, playerId int) bool {
	if !CanMove(x, y, game, playerId) {
		return false
	}
	game.GameBoard[x][y] = playerId
	if CheckRow(x, game, playerId) {
		return true
	}
	if CheckColumn(y, game, playerId) {
		return true
	}
	if x == y && CheckDiagonal(1, game, playerId) {
		return true
	}
	if x+y == (game.BoardSize-1) && CheckDiagonal(-1, game, playerId) {
		return true
	}
	game.Moves++
	SetNextPlayer(game)
	return false
}

func CanMove(x int, y int, game *Game, playerId int) bool {
	if game.Completed == true {
		return false
	}
	if playerId != game.NextPlayer {
		return false
	}
	if x >= game.BoardSize || x < 0 || y >= game.BoardSize || y < 0 {
		return false
	}
	if game.GameBoard[x][y] != 0 {
		return false
	}
	if game.Moves >= (game.BoardSize * game.BoardSize) {
		return false
	}
	return true
}

func SetNextPlayer(game *Game) {
	if game.Player1Id == game.NextPlayer {
		game.NextPlayer = game.Player2Id
	} else {
		game.NextPlayer = game.Player1Id
	}
	fmt.Println(game.NextPlayer)
}

func CheckRow(x int, game *Game, playerId int) bool {
	for i := 0; i < game.BoardSize; i++ {
		if game.GameBoard[x][i] != playerId {
			return false
		}
	}
	return true
}

func CheckColumn(y int, game *Game, playerId int) bool {
	for i := 0; i < game.BoardSize; i++ {
		if game.GameBoard[i][y] != playerId {
			return false
		}
	}
	return true
}

func CheckDiagonal(direction int, game *Game, playerId int) bool {
	j := 0
	if direction == -1 {
		j = game.BoardSize - 1
	}
	for i := 0; i < game.BoardSize; i++ {
		if game.GameBoard[i][j] != playerId {
			return false
		}
		j = j + direction
	}
	return true
}
