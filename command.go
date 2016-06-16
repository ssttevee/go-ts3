package ts3

import (
	"fmt"
	"github.com/ssttevee/go-ts3/util"
	"log"
	"strings"
)

// A Command is a wrapper for Server Query commands
type Command interface {
	Command() string
}

// Command sends a string command to the Server Query
func (c *Client) Command(command string) (ServerResponse, error) {
	// TODO add verbose option
	log.Println("send cmd:", command)
	if _, err := fmt.Fprintln(c.t, command); err != nil {
		return "", err
	}

	var payload string
	for {
		out, err := c.readLine()
		if err != nil {
			return "", err
		}

		if strings.HasPrefix(out, "error") {
			res := parseResponse(out)

			id, ok := res["id"]
			if !ok {
				return "", ErrBadResponseFormat
			}

			msg, ok := res["msg"]
			if !ok {
				return "", ErrBadResponseFormat
			}

			ret := ServerResponse(strings.TrimSpace(payload))
			if res["id"] != "0" {
				return ret, &Error{id, ts3util.Unescape(msg)}
			}

			return ret, nil
		}

		payload += out + "\n"
	}
}

// Commandf formats a string according to a
// format specifier and sends it to the Server Query.
func (c *Client) Commandf(format string, a ...interface{}) (ServerResponse, error) {
	return c.Command(fmt.Sprintf(format, a...))
}

// DoCommand sends a command to the Server Query
func (c *Client) DoCommand(cmd Command) (ServerResponse, error) {
	return c.Command(cmd.Command())
}
