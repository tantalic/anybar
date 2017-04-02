// +build !darwin

package anybar // import "tantalic.com/anybar"

import (
	"fmt"
)

func (c *Client) Start() error {
	return fmt.Errorf("anybar can only be started on the localhost (client is set to %s)", c.Hostname)
}
