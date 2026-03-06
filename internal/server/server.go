package server

import (
	"context"
	"log/slog"
	"net"
	"strconv"
	"sync"

	"github.com/p-alvarenga/go_tcp-tracker-server/internal/config"
	"github.com/p-alvarenga/go_tcp-tracker-server/internal/device"
	"github.com/p-alvarenga/go_tcp-tracker-server/internal/domain/types"
	"github.com/p-alvarenga/go_tcp-tracker-server/internal/session"
)

type Server struct {
	cfg *config.ServerConfig

	listener net.Listener

	sessions map[types.IMEI]*session.Session
	devices  *device.DeviceManager

	ctx    context.Context
	cancel context.CancelFunc

	mu sync.Mutex
	wg sync.WaitGroup

	logger     *slog.Logger
	rootLogger *slog.Logger
}

func NewServer(cfg *config.ServerConfig, rootLogger *slog.Logger) *Server {
	ctx, cancel := context.WithCancel(context.Background())

	return &Server{
		cfg: cfg,

		sessions: make(map[types.IMEI]*session.Session),
		devices:  device.NewDeviceManager(),

		ctx:    ctx,
		cancel: cancel,

		logger:     rootLogger.With("lyr", "Server"),
		rootLogger: rootLogger,
	}
}

func (s *Server) Boot() error {
	var err error
	addr := net.JoinHostPort(s.cfg.ServerHost, strconv.Itoa(s.cfg.ServerPort))

	s.listener, err = net.Listen("tcp", addr)
	if err != nil {
		s.logger.Error("Could not initialize server", "err", err)
	}

	s.logger.Info("Server booted correctly", "addr", addr)
	go s.loop()

	<-s.ctx.Done()

	return nil
}
