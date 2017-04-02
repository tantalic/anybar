package anybar // import "tantalic.com/anybar"

import (
	"fmt"
	"os"
	"os/exec"
)

// Start opens AnyBar.app. If an instance of AnyBar is already running on
// .Port that instance will be quit first. This function only works on
// macOS and requires AnyBar.app to be installed.
func (c *Client) Start() error {
	if c.Hostname != "" && c.Hostname != "localhost" && c.Hostname != "127.0.0.1" {
		return fmt.Errorf("anybar can only be started on the localhost (client is set to %s)", c.Hostname)
	}

	// Quit, if already running
	c.Quit()

	open := exec.Command("/usr/bin/open", "-n", "-a", "AnyBar")
	env := fmt.Sprintf("ANYBAR_PORT=%d", c.Port)
	open.Env = append(os.Environ(), env)
	return open.Start()
}
