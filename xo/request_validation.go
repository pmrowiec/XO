package main

func (gameSettings *GameStart) Validate() bool {
	if gameSettings.BoardSize < 3 || gameSettings.BoardSize > 9 {
		gameSettings.BoardSize = 3
	}
	if gameSettings.Player1Id < 1 || gameSettings.Player2Id < 1 || gameSettings.Player1Id == gameSettings.Player2Id {
		return false
	}
	return true
}

func (move *GameMove) Validate() bool {
	if move.X < 0 || move.Y < 0 || move.PlayerId < 1 {
		return false
	}
	return true
}
