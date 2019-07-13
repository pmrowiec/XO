package main

func (game Game) GetGameState() GameState {
	var state GameState
	state.Id = game.Id
	state.Player1Id = game.Player1Id
	state.Player2Id = game.Player2Id
	state.Player1Mark = game.Marks[game.Player1Id]
	state.Player2Mark = game.Marks[game.Player2Id]
	state.GameBoard = game.GameBoard
	state.Completed = game.Completed
	state.NextPlayer = game.NextPlayer
	state.Winner = game.Winner
	return state
}
