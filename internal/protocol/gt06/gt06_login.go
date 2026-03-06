package gt06

import (
	"fmt"

	"github.com/p-alvarenga/go_tcp-tracker-server/internal/domain/types"
)

type LoginPacket struct {
	IMEI   types.IMEI
	Serial int
}

func (l *LoginPacket) Type() PacketType {
	return LoginType
}

func decodeLogin(payload []byte) (*LoginPacket, error) {
	if len(payload) != 8 { // Binary Coded Decimal (BCD)
		return nil, fmt.Errorf("gt06: invalid BCD IMEI. Length: %d, must be 8 bytes", len(payload))
	}

	return &LoginPacket{
		IMEI: types.IMEI(bcdToASCII(payload)),
	}, nil
}
