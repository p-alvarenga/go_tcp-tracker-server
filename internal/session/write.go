package session

import "fmt"

func (s *Session) writeLoop() {
	for {
		select {
		case msg, ok := <-s.WriteChan:
			if !ok {
				s.logger.Error("Write channel closed")
				return
			}

			s.logger.Info("Writting into stream", "msg", fmt.Sprintf(" %X", msg))
			_, err := s.conn.Write(msg)
			if err != nil {
				s.logger.Error("Could not write into stream", "err", err)
			}

		case <-s.ctx.Done():
			return
		}
	}
}
