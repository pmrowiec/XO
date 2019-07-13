package main

import (
	"math/rand"
	"strconv"
	"time"
)

func CreateGame(settings GameStart) Game {
	var newGame Game
	newGame.Player1Id = settings.Player1Id
	newGame.Player2Id = settings.Player2Id
	newGame.BoardSize = settings.BoardSize
	newGame.Marks = make(map[int]string)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(2)
	if i == 0 {
		newGame.NextPlayer = newGame.Player1Id
		newGame.Marks[newGame.Player1Id] = "X"
		newGame.Marks[newGame.Player2Id] = "O"
	} else {
		newGame.NextPlayer = newGame.Player2Id
		newGame.Marks[newGame.Player2Id] = "X"
		newGame.Marks[newGame.Player1Id] = "O"
	}
	newGame.Id = strconv.Itoa(newGame.Player1Id) + "-" + strconv.Itoa(newGame.Player2Id)
	newGame.GameBoard = *MakeBoard(newGame.BoardSize)
	return newGame
}

func MakeBoard(size int) *Board {
	board := make(Board, size)
	for i := range board {
		board[i] = make(BoardRow, size)
	}
	return &board
}
