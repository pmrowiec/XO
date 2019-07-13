<h1>XO</h1>

<h3>Requirements:</h3>
<p>Build a simple service to allow two users to remotely play Noughts & Crosses.</p>
<p>A user, either X or O, should be able to make a move, see the current state of the game and find out the winner via an API. The design of this simple API is up to you. No authentication is necessary.</p>

<h3>What has been done:</h3>
<p>If this was a production code, I would have started with clarifying what the exact requirements are. Since this is a hypothetical task, I decided to turn uncertainties into assumptions as follow:</p>
<ul>
<li>Players can play the game remotely through a website - my service consists of a few web services, the front side of the game is left for implementation.</li>
<li>"authentication isn't needed", but players have unique ids - if the wider system allows user registration, the ids are ids of registered players; otherwise, players connecting to the server are assigned guest ids for the duration of their session (this part is outside the scope of my project, I just assume players have unique ids).</li>
<li>I assumed that prior to starting a game, players pair up - one player invites the other to the game. The first player is the owner of the game "room", the second one is a guest. This also indicates the id of the room and, in result, the id of any future games played by the pair (the room/game id is a concatenation of ids of both players, in a specific order: room owner's id + '-' + guest's id).
<ul>
<li>Another approach I considered was allowing players to match up against any other random player or, if there are no other players in the queue, an AI opponent. I believe this would be a good addition to the project - that way users would be able to play this quick game without having to involve a friend. I assumed that the original goal was to allow two specific players to remotely play this game against each other, hence my approach.</li></ul></li>
<li>The request to start a new game comes with ids of both players and the size of the board (the original version of Noughts and Crosses uses a 3x3 board, but players can choose any size between 3 and 9)</li>
<li>A new game is created with the generated unique id (a concatenation of both players' ids, as mentioned before). A board with the requested size is initialized and one of the two players is randomly chosen to start the game. The player making the first move is "X", the other - "O".</li>
<li>I had a dilemma what to do in case of multiple requests to start a new game for the same pair of players - there should only be at most one active game for the same pair of players, but what if there's a request to create a new game before the previous one ends? Similarly, what if both players press 'new game' at the same time? 
<ul><li>I decided to allow players to restart the game (or, in other words, create a new one) when the active one is in progress - if the players are friends, we can assume they want to actually play the game, not annoy each other by restarting the game when it doesn't go well for them. Also, since games can be created with larger board sizes (e.g. 9x9), players should be able to decide to finish it prematurely and start a new one, potentially on a different size board. </li>
<li>Allowing players to restart the game at any time could cause problems when both players press the 'new game' button at the same time (creating one game and instantly overwriting it with another instance), which could result in obsolete state of the game presented to one of the players until the next update cycle (both requests would return game state info about two different games, possibly with different board sizes and different players chosen to make the first move). This could be solved using locks and other principles of concurrency or, in a simplified version, we can assume that only the owner of the pair of players is allowed to create a new game. This still isn't a perfect solution, but with an appropriate debounce added in the front-end (blocking the owner of the pair from creating a new game within a few seconds from the previous one), it should do the job.</li></ul></li>
<li>After a game is created, its current state is returned to the player (the board of a chosen size, the id of the player whose turn it is, the state of the game - completed or not (in this case not, as no moves have been made), assignment of marks ("X"/"O") etc (see API reference below).</li>
<li>There's a service that returns the state of a chosen game (same as above, but doesn't create a new game, only returns the state of the game if it exists).</li>
<li>Players send requests with their moves (the request contains the id of the player, the id of the game and the coordinates - x and y - of where they want to place their mark).</li>
<li>To check whether the game has been completed (won by the player making the move), we check the row/column containing the cell with the current move. If the move is made at diagonals, we also check them.</li>
<li>We keep track of the total number of moves made within each game - a game can end with no winner if all fields have been marked. </li>
<li>Each player periodically sends a request to ask about the current state of the game (every few seconds), to find out about the move made by the other player or to find out that a new game has been made by the pair's owner. </li>
</ul>
<h3>What's left:</h3>
<ul>
<li>In order to avoid conflicts (two players sending their moves at the same time - or the same player sending multiple requests in a row - before the logic behind the previous one has fully been executed), we should consider using a semaphore or other tools for handling concurrency. (The game state available to game's front-end knows whose turn it is, so it can also block sending move requests if it's the other player's turn).</li>
<li>We don't delete games straight away after they have been completed - the other player needs to have a chance to find out what the result was. We should also keep in mind that players may abandon games in progress. To avoid cluttering server's memory with unnecessary, inactive games, we could store games in an ordered data structure based on the time of the last update of the game and periodically delete games older than some agreed age.</li>
<li>Authentication/assignment of guest ids to players connecting to the server.</li>
<li>Making pairs of players - either through one player inviting the other or by random assignment of an opponent.</li>
<li>AI opponent.</li>
</ul>
<h3>Running the project:</h3>
<p>Installing dependencies needed for the project:</p>
<ul><li>go get (in terminal)</li></ul>
<p>Building the project:</p>
<ul><li>go build (in terminal)</li></ul>
<p>Building the project:</p>
<ul><li>xo (in terminal)</li></ul>
<p>This starts the server at http://localhost:8080</p>

<h3>Using the project / playing a sample game:</h3>
<ul><li>An app or a browser add-on for sending POST HTTP requests will be needed</li>
<li>We assume the players have already been paired and have their unique ids. Any positive integer ids will be fine, but they have to be unique. We'll assume 1 (owner of the game "room") and 2 (invited guest).</li>
<li>Send a request to create a new game on a 3x3 board:
<pre>http://localhost:8080/game
method: POST
request body - JSON object - GameStart:
{
	"player1id" : 1,
	"player2id" : 2,
	"boardsize" : 3,
}
</pre>
The response will tell us the id of the player expected to make the first move, along with other info. Note that the generated game id is 1-2 (player1id + '-' + player2id). The board will be empty, so it will look like:
<pre>_|_|_
_|_|_
 | |</pre>
</li>
<li>Send a request to make the first move:
<pre>http://localhost:8080/game/1-2/move
method: POST
request body - JSON object - GameMove:
{
	"playerid" : 1, <----- 1 or 2, depending on whose turn it is, the information was returned from previous service call
	"x" : 0,
	"y" : 2
}
</pre>
The board will look like:
<pre>_|_|X
_|_|_
 | |</pre>
</li>
<li>Continue to make moves, alternating the playerid. If the next moves are O - (1,1), X - (0,1), O - (1,2), X - (0,0), the board will look like this:
<pre>X|X|X
_|O|O
 | |</pre>
The response will now say that the game has been completed and the winner is player1:
<pre>
{
	"id" : "1-2",
	"player1id" : 1,
	"player2id" : 2,
	"player1mark" : "X",
	"player2mark" : "O",
	"board" : [["X","X","X"],["","O","O"],["","",""]],
	"completed" : true,
	"nextplayer" : 1,
	"winnerid" : 1
}
</pre>
</li>
<li>At any point, you can ask the server about the state of the game:</li>
<pre>http://localhost:8080/game/1-2
method: POST
response body - JSON object - GameState:
{
	"id" : "1-2",
	"player1id" : 1,
	"player2id" : 2,
	"player1mark" : "X",
	"player2mark" : "O",
	"board" : [["","X","X"],["","O",""],["","",""]],
	"completed" : false,
	"nextplayer" : 2,
	"winnerid" : 0
}</pre></ul>

<h3>Tests:</h3>
<p>A few tests have been implemented to validate the core logic (creating a new game, making a move, checking the winner/end of the game with a draw).</p>
<p>Running the tests: </p>
<ul><li>go test -v(in terminal)</li></ul>

<h3>API reference:</h3>
<p><b><u>creating a new game:</u></b></p>
<pre>http://localhost:8080/game
method: POST
request body - JSON object - GameStart:
{
	"player1id" : 1,
	"player2id" : 2,
	"boardsize" : 3,
}

player1id - unique id of the owner of the game room 
player2id - unique id of the guest player
boardSize - any number in the range [3, 9], for creating a 3x3 to 9x9 board. If a number outside the range is provided, a 3x3 board is created

response body - JSON object - GameState:
{
	"id" : "1-2",
	"player1id" : 1,
	"player2id" : 2,
	"player1mark" : "X",
	"player2mark" : "O",
	"board" : [["","X","X"],["","O",""],["","",""]],
	"completed" : false,
	"nextplayer" : 2,
	"winnerid" : 0
}

id - id of the game
player1id - unique id of the owner of the game room 
player2id - unique id of the guest player
player1mark - "X" or "O", depending on which player was chosen to start the game
player2mark - "X" or "O", depending on which player was chosen to start the game
board - a 2D array representing the board
completed - true is the game is finished, false if it is in progress
nextplayer - id of the player whose turn it is
winnerid - id of the player who has won the game (if the game is completed), 0 if the game is in progress or ended with a draw
</pre>

<p><b><u>checking current state of the game:</u></b></p>
<pre>http://localhost:8080/game/{gameId}
method: POST
response body - JSON object - GameState, same as above</pre>

<p><b><u>making a move</u></b></u></b></p>
<pre>http://localhost:8080/game/{gameId}/move
method: POST
request body - JSON object - GameMove:
{
	"playerid" : 2,
	"x" : 1,
	"y" : 0
}

gameid - unique id of the game, concatenation of both players' id separated with a dash; the first id is the owner of the game, the second is the guest
playerid - player making the move
x - x coordinate of the move
y - y coordinate of the move

response body - JSON object - GameState, same as above</pre>
