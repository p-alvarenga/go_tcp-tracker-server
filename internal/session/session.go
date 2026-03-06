package session

import (
	"context"
	"log/slog"
	"net"

	"github.com/p-alvarenga/go_tcp-tracker-server/internal/device"
	"github.com/p-alvarenga/go_tcp-tracker-server/internal/domain/types"
)

type Session struct {
	conn     net.Conn
	deviceId types.IMEI

	devicesManager *device.DeviceManager

	WriteChan chan []byte
	ReadChan  chan []byte
	buffer    []byte

	ctx    context.Context
	cancel context.CancelFunc

	logger *slog.Logger
}

func New(conn net.Conn, devManager *device.DeviceManager, parentCtx context.Context, rootLogger *slog.Logger) *Session {
	ctx, cancel := context.WithCancel(parentCtx)

	return &Session{
		conn:     conn,
		deviceId: "null",

		devicesManager: devManager,

		WriteChan: make(chan []byte), // buffered?
		ReadChan:  make(chan []byte), // buffered?

		ctx:    ctx,
		cancel: cancel,

		logger: rootLogger.With("lyr", "Session"),
	}
}

func (c *Session) Run() {
	go c.readLoop()
	go c.writeLoop()

	<-c.ctx.Done()
}

func (c *Session) shutdown() {
	c.logger.Warn("shutting down connection")
	c.cancel()
}
