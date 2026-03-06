package gt06

import (
	"fmt"

	"github.com/p-alvarenga/go_tcp-tracker-server/internal/domain/types"
	"github.com/p-alvarenga/go_tcp-tracker-server/internal/protocol"
)

type LoginPacket struct {
	IMEI   types.IMEI
	Serial int
}

func decodeLogin(payload []byte) (*LoginPacket, error) {
	if len(payload) != 8 { // Binary Coded Decimal (BCD)
		return nil, fmt.Errorf("gt06: invalid BCD IMEI. Length: %d, must be 8 bytes", len(payload))
	}

	return &LoginPacket{
		IMEI: types.IMEI(protocol.BCDToASCII(payload)),
	}, nil
}

func (l *LoginPacket) Type() PacketType {
	return LoginType
}

func (l *LoginPacket) ToACKPacket() {

}
