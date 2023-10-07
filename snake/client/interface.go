package client

import "github.com/theboarderline/ebiten-utilities/snake/object/snake"

type Client interface {
	// Connect description of the Go function.
	//
	// Connect connects to the server.
	// Returns an error if the connection fails.
	Connect() error

	// IsConnected returns a boolean value indicating whether the connection is established.
	//
	// No parameters.
	// Returns a boolean value.
	IsConnected() bool

	// GetPlayers returns a list of players.
	//
	// No parameters.
	GetPlayers() map[string]*snake.Snake

	// Register is a function that registers a name.
	//
	// It takes a string parameter called 'name' and returns an error.
	Register(name string) error

	// Deregister is a function that deregisters a given name.
	//
	// It takes a parameter:
	// - name (string): The name to be deregistered.
	//
	// It returns an error if the deregistration fails.
	Deregister(name string) error

	// SendMessage sends the given message.
	//
	// message: The message to send.
	// error: An error if the message failed to send.
	SendMessage(message []byte) error

	// GetMessage returns the message and an error.
	//
	// It takes no parameters.
	// It returns a string and an error if there was an issue during retrieval.
	GetMessage() (string, error)

	// Cleanup is a function that performs cleanup operations.
	//
	// It takes no parameters.
	// It returns an error if there was an issue during cleanup.
	Cleanup() error
}
