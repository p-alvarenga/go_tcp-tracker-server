package server

import "github.com/p-alvarenga/go_tcp-tracker-server/internal/session"

func (s *Server) loop() {
	for {
		select {
		case <-s.ctx.Done():
			return
		default:
		}

		conn, err := s.listener.Accept()
		if err != nil {
			s.logger.Error("Could not accept connection")
			return
		}

		session := session.New(conn, s.devices, s.ctx, s.rootLogger)
		session.Run()
	}
}
