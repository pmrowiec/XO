package main

type Board []BoardRow
type BoardRow []string

type Game struct {
	Id         string
	Player1Id  int
	Player2Id  int
	Marks      map[int]string
	BoardSize  int
	GameBoard  Board
	Completed  bool
	NextPlayer int
	Moves      int
	Winner     int
}

type Games map[string]*Game

type GameMove struct {
	X        int `json:"x"`
	Y        int `json:"y"`
	PlayerId int `json:"playerid"`
}

type GameState struct {
	Id          string `json:"id"`
	Player1Id   int    `json:"player1id"`
	Player2Id   int    `json:"player2id"`
	Player1Mark string `json:"player1mark"`
	Player2Mark string `json:"player2mark"`
	GameBoard   Board  `json:"board"`
	Completed   bool   `json:"completed"`
	NextPlayer  int    `json:"nextplayer"`
	Winner      int    `json:"winnerid"`
}

type GameStart struct {
	Player1Id int `json:"player1id"`
	Player2Id int `json:"player2id"`
	BoardSize int `json:"boardsize"`
}
