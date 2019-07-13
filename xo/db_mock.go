package main

var games Games

func init() {
	games = make(map[string]*Game)
}

func AddGame(settings GameStart) Game {
	newGame := CreateGame(settings)
	// Overwrite previous game of the same pair of players if one exists
	games[newGame.Id] = &newGame
	return newGame
}

func FindGame(id string) *Game {
	for _, g := range games {
		if g.Id == id {
			return g
		}
	}
	return nil
}
