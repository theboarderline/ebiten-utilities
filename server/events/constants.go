package events

const (
	// PLAYER_CONNECTED checks if the player is connected.
	PLAYER_CONNECTED = "PLAYER_CONNECTED"

	// PLAYER_CONNECT connects the player to the server
	PLAYER_CONNECT = "PLAYER_CONNECT"

	// PLAYER_DISCONNECT disconnects the player from the server
	PLAYER_DISCONNECT = "PLAYER_DISCONNECT"

	// ACK is a basic acknowledgement
	ACK = "ACK"

	// GET_PLAYERS retrieves the list of players in the game
	GET_PLAYERS = "GET_PLAYERS"

	// PLAYER_INPUT sends the player's keyboard input
	PLAYER_INPUT = "PLAYER_INPUT"

	// PLAYER_COUNT counts the number of players in the game
	PLAYER_COUNT = "PLAYER_COUNT"

	// PLAYER_CAPACITY counts the number of players in the game
	PLAYER_CAPACITY = "PLAYER_CAPACITY"

	// PLAYER_SCORE add score to the player
	PLAYER_SCORE = "PLAYER_SCORE"

	// DEFAULT_ENEMY_NAME
	ENEMY = "Enemy"
)
