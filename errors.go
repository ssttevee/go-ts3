package ts3

import (
	"errors"
	"fmt"
)

var (
	// ErrNoResponse is returned when the server doesn't response
	ErrNoResponse = errors.New("No response from the server query. Make sure your IP address is whitelisted.")
	// ErrBadResponseFormat is returned when the server returns a malformed response
	ErrBadResponseFormat = errors.New("The response is in the wrong format. Make sure your server software is up to date.")
	// ErrBadWelcome is returned when the server sends an unexpected welcome message
	ErrBadWelcome = errors.New("Got an unexpected welcome message.  Make sure you're connecting to a TeamSpeak 3 server.")
)

// A Error holds an error received from the Server Query
type Error struct {
	Code    string
	Message string
}

// Error formats and returns the error
func (err *Error) Error() string {
	return fmt.Sprintf("ServerQuery error %s: %s", err.Code, err.Message)
}
