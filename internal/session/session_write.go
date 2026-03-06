package session

import "fmt"

func (c *Session) writeLoop() {
	for {
		select {
		case msg, ok := <-c.writeCh:
			if !ok {
				c.logger.Error("Write channel closed")
				return
			}

			c.logger.Info("Writting into stream", "msg", fmt.Sprintf(" %X", msg))
			_, err := c.conn.Write(msg)
			if err != nil {
				c.logger.Error("Could not write into stream", "err", err)
			}

		case <-c.ctx.Done():
			return
		}
	}
}
