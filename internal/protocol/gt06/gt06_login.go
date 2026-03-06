package gt06

import (
	"fmt"

	"github.com/p-alvarenga/go_tcp-tracker-server/internal/device"
)

type LoginPacket struct {
	Imei   device.Imei
	Serial uint
}

func (l *LoginPacket) Type() PacketType {
	return LoginType
}

func decodeLogin(payload []byte) (*LoginPacket, error) {
	if len(payload) != 8 { // Binary Coded Decimal (BCD)
		return nil, fmt.Errorf("gt06: invalid BCD IMEI. Length: %d, must be 8 bytes", len(payload))
	}

	return &LoginPacket{
		Imei: device.Imei(bcdToASCII(payload)),
	}, nil
}
