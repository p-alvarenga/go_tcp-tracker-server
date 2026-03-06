package session

import (
	"fmt"
)

func (c *Session) readLoop() {
	buf := make([]byte, 4096)

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		n, err := c.conn.Read(buf) // use bytes read
		if err != nil {
			c.shutdown()
			return
		}

		c.logger.Info(fmt.Sprintf("%d bytes read", n), "raw", fmt.Sprintf("% s", buf[:n]))
	}
}
