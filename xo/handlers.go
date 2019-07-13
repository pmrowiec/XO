package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteResponse(w http.ResponseWriter, responseObj interface{}, httpStatus int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpStatus)
	if err := json.NewEncoder(w).Encode(responseObj); err != nil {
		panic(err)
	}
}

func ReadRequest(r *http.Request) []byte {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 500))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	return body
}

func GetGameContext(r *http.Request) *Game {
	vars := mux.Vars(r)
	gameId := vars["gameId"]
	return FindGame(gameId)
}

func GameGetState(w http.ResponseWriter, r *http.Request) {
	game := GetGameContext(r)
	if game == nil {
		WriteResponse(w, nil, http.StatusNotFound)
	} else {
		WriteResponse(w, game.GetGameState(), http.StatusFound)
	}
}

func GameCreate(w http.ResponseWriter, r *http.Request) {
	var gameSettings GameStart
	body := ReadRequest(r)
	if err := json.Unmarshal(body, &gameSettings); err != nil || !gameSettings.Validate() {
		WriteResponse(w, "Invalid game settings", http.StatusUnprocessableEntity)
		return
	}

	newGame := AddGame(gameSettings)
	WriteResponse(w, newGame.GetGameState(), http.StatusCreated)
}

func GameMakeMove(w http.ResponseWriter, r *http.Request) {
	var move GameMove
	body := ReadRequest(r)
	if err := json.Unmarshal(body, &move); err != nil || !move.Validate() {
		WriteResponse(w, "Invalid move", http.StatusUnprocessableEntity)
		return
	}
	game := GetGameContext(r)
	if game == nil {
		WriteResponse(w, nil, http.StatusNotFound)
		return
	}

	game.Move(move.X, move.Y, move.PlayerId)
	WriteResponse(w, game.GetGameState(), http.StatusOK)
}
