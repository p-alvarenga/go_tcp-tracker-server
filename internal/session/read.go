package session

func (s *Session) readLoop() {
	s.buffer = make([]byte, 4096)

	for {
		select {
		case <-s.ctx.Done():
			return
		default:
		}

		_, err := s.conn.Read(s.buffer)
		if err != nil {
			s.shutdown()
			return
		}
	}
}
