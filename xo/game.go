package main

import "time"

type Board []BoardRow
type BoardRow []int

type Game struct {
	Id         string
	Player1Id  int
	Player2Id  int
	BoardSize  int
	GameBoard  Board
	Completed  bool
	NextPlayer int
	Moves      int
	Winner     int
	Started    time.Time
}

type Games map[string]*Game

type GameMove struct {
	GameId   string `json:"gameid"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	PlayerId int    `json:"playerid"`
}

type GameState struct {
	Id         string `json:"id"`
	Player1Id  int    `json:"player1id"`
	Player2Id  int    `json:"player2id"`
	GameBoard  Board  `json:"board"`
	Completed  bool   `json:"completed"`
	NextPlayer int    `json:"nextplayer"`
	Winner     int    `json:"winner"`
}

type GameStart struct {
	Player1Id int `json:"player1id"`
	Player2Id int `json:"player2id"`
	BoardSize int `json:"boardsize"`
}
