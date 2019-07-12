package main

import (
	"math/rand"
	"strconv"
	"time"
)

var games Games

func init() {
	games = make(map[string]*Game)
}

func CreateGame(settings GameStart) Game {
	var newGame Game
	newGame.Player1Id = settings.Player1Id
	newGame.Player2Id = settings.Player2Id
	newGame.BoardSize = settings.BoardSize
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(2)
	if i == 0 {
		newGame.NextPlayer = newGame.Player1Id
	} else {
		newGame.NextPlayer = newGame.Player2Id
	}
	newGame.Id = strconv.Itoa(newGame.Player1Id) + "-" + strconv.Itoa(newGame.Player2Id)
	newGame.Started = time.Now()
	newGame.GameBoard = *MakeBoard(newGame.BoardSize)
	games[newGame.Id] = &newGame
	return newGame
}

func MakeBoard(size int) *Board {
	board := make(Board, size)
	for i := range board {
		board[i] = make(BoardRow, size)
	}
	return &board
}

func FindGame(id string) *Game {
	for _, g := range games {
		if g.Id == id {
			return g
		}
	}
	return nil
}
