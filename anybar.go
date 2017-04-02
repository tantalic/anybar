// Package anybar provides a client for interacting with the
// AnyBar macOS menubar status indicator app.
package anybar // import "tantalic.com/anybar"

import "net"
import "fmt"

// Client is used to send commands to the AnyBar application.
type Client struct {
	// Port is the UDP port used to control the AnyBar application.
	Port int

	// Hostname is the host where AnyBar is running.
	Hostname string
}

// The following default values will be used on anybar.Client
// if not set.
const (
	DefaultPort     = 1738
	DefaultHostname = "localhost"
)

// Set changes the style of the AnyBar menu bar icon.
func (c *Client) Set(s Style) error {
	return c.send([]byte(s))
}

// Quit tells the AnyBar app to exit.
func (c *Client) Quit() error {
	return c.send([]byte(quit))
}

func (c *Client) send(b []byte) error {
	laddr, err := net.ResolveUDPAddr("udp", ":0")
	if err != nil {
		return err
	}

	anybarAddr, err := net.ResolveUDPAddr("udp4", c.addr())
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp4", laddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.WriteToUDP(b, anybarAddr)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) addr() string {
	if c.Hostname == "" {
		c.Hostname = DefaultHostname
	}

	if c.Port == 0 {
		c.Port = DefaultPort
	}

	return fmt.Sprintf("%s:%d", c.Hostname, c.Port)
}

// Style is used to represent the various appearances that can be applied
// to AnyBar icon in the menu bar. You should not create your own Style
// instances but instead use one of the provided constants in the package.
type Style string

// The following styles are available for controlling the appearance of the
// AnyBar menu bar icon.
const (
	White  = Style("white")
	Red    = Style("red")
	Orange = Style("orange")
	Yellow = Style("yellow")
	Green  = Style("green")
	Cyan   = Style("cyan")
	Blue   = Style("blue")
	Purple = Style("purple")
	Black  = Style("black")

	Question    = Style("question")
	Exclamation = Style("exclamation")
)

const (
	quit = "quit"
)
