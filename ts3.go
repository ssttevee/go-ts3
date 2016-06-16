// Package ts3 makes it easier to use with the TeamSpeak 3's Server Query interface
package ts3

import (
	"github.com/ssttevee/go-ts3/enum"
	"github.com/ssttevee/go-ts3/util"
	"github.com/ziutek/telnet"
	"io"
	"log"
	"strings"
	"time"
)

const (
	// WelcomeMessage is used to match the welcome message received from the server
	WelcomeMessage = "TS3"
)

var (
	// WelcomeTimeout is the longest time to wait for
	// before assuming that the server will not respond
	WelcomeTimeout = 5 * time.Second
	// ResponseBufferSize is the largest number
	// of lines to hold in the receiving buffer
	ResponseBufferSize = 10
)

type line struct {
	res string
	err error
}

// A Client interfaces with TeamSpeak 3's Server Query interface
type Client struct {
	t *telnet.Conn

	// NotifyFunc is called whenever a notify-prefixed
	// message is received from the server
	NotifyFunc func(response ServerResponse) error
	// DisconnectFunc is called when the client is disconnected
	// from the server and not does not try to reconnect
	DisconnectFunc func()

	lines chan line
	exit  bool
}

func dial(conn *telnet.Conn, err error) (*Client, error) {
	if err != nil {
		return nil, err
	}

	// ready the return object
	c := &Client{
		t:     conn,
		lines: make(chan line, ResponseBufferSize)}

	// BEGIN SETUP

	// start a go routine to read lines as they come
	go func() {
		for !c.exit {
			s, err := conn.ReadString(telnet.LF)
			if err == io.EOF {
				log.Println("connection lost")
				if c.DisconnectFunc != nil {
					c.DisconnectFunc()
				}

				return
			} else if err != nil {
				log.Fatal("read line loop: ", err)
			}

			if s := strings.TrimSpace(s); strings.HasPrefix(s, "notify") {
				go func() { // must be in a new goroutine or else it'll block
					if f := c.NotifyFunc; f != nil {
						if err := f(ServerResponse(s)); err != nil {
							log.Println("notify func error:", err)
							c.Commandf(`logadd loglevel=%s logmsg=NotifyCallback\sError:\s%s`, ts3enum.LogLevelError, ts3util.Escape(err.Error()))
						}
					}
				}()
			} else {
				c.lines <- line{s, err}
			}
		}
	}()

	// channel to capture welcome error
	done := make(chan error)

	// go routine to get welcome message
	go func() {
		welcome, err := c.readLine()
		if err != nil {
			done <- err
		} else if welcome != WelcomeMessage {
			done <- ErrBadWelcome
		}

		_, err = c.readLine()
		done <- err
	}()

	// Wait for welcome message or timeout
	select {
	case err := <-done:
		if err != nil {
			return nil, err
		}
	case <-time.NewTimer(WelcomeTimeout).C:
		return nil, ErrNoResponse
	}

	return c, nil
}

// Dial opens a new connection to a TeamSpeak 3's Server Query
func Dial(address string) (*Client, error) {
	return dial(telnet.Dial("tcp", address))
}

// DialTimeout opens a new connection to a TeamSpeak 3's Server Query with a specific timeout.
func DialTimeout(address string, timeout time.Duration) (*Client, error) {
	return dial(telnet.DialTimeout("tcp", address, timeout))
}

// Close cuts the connection to the server
func (c *Client) Close() error {
	if err := c.t.Close(); err != nil {
		return err
	}

	// stop go routines
	c.exit = true

	return nil
}

func (c *Client) readLine() (string, error) {
	line := <-c.lines
	return line.res, line.err
}
