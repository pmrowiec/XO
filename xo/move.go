package main

func (game *Game) Move(x int, y int, playerId int) {
	if !game.CanMove(x, y, playerId) {
		return
	}
	game.GameBoard[x][y] = game.Marks[playerId]
	if game.CheckRow(x, playerId) {
		game.SetWinner(playerId)
		return
	}
	if game.CheckColumn(y, playerId) {
		game.SetWinner(playerId)
		return
	}
	if x == y && game.CheckDiagonal(1, playerId) {
		game.SetWinner(playerId)
		return
	}
	if x+y == (game.BoardSize-1) && game.CheckDiagonal(-1, playerId) {
		game.SetWinner(playerId)
		return
	}
	game.Moves++
	game.CheckDraw()
	game.SetNextPlayer()
	return
}

func (game *Game) CanMove(x int, y int, playerId int) bool {
	if game.Completed == true {
		return false
	}
	if playerId != game.NextPlayer {
		return false
	}
	if x >= game.BoardSize || x < 0 || y >= game.BoardSize || y < 0 {
		return false
	}
	if game.GameBoard[x][y] != "" {
		return false
	}
	if game.Moves >= (game.BoardSize * game.BoardSize) {
		return false
	}
	return true
}

func (game *Game) SetNextPlayer() {
	if game.Player1Id == game.NextPlayer {
		game.NextPlayer = game.Player2Id
	} else {
		game.NextPlayer = game.Player1Id
	}
}

func (game *Game) CheckRow(x int, playerId int) bool {
	for i := 0; i < game.BoardSize; i++ {
		if game.GameBoard[x][i] != game.Marks[playerId] {
			return false
		}
	}
	return true
}

func (game *Game) CheckColumn(y int, playerId int) bool {
	for i := 0; i < game.BoardSize; i++ {
		if game.GameBoard[i][y] != game.Marks[playerId] {
			return false
		}
	}
	return true
}

func (game *Game) CheckDiagonal(direction int, playerId int) bool {
	j := 0
	if direction == -1 {
		j = game.BoardSize - 1
	}
	for i := 0; i < game.BoardSize; i++ {
		if game.GameBoard[i][j] != game.Marks[playerId] {
			return false
		}
		j = j + direction
	}
	return true
}

func (game *Game) SetWinner(playerId int) {
	game.Winner = playerId
	game.Completed = true
}

func (game *Game) CheckDraw() {
	if game.Moves >= game.BoardSize*game.BoardSize {
		game.Completed = true
	}
}
