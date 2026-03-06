package gt06

import (
	"encoding/binary"
	"fmt"

	"github.com/p-alvarenga/go_tcp-tracker-server/internal/protocol"
)

func BuildACK(packetType PacketType, serial int) ([]byte, error) {
	if !packetType.isValid() {
		return nil, fmt.Errorf("Invalid packet type")
	}

	buf := make([]byte, 0, 10)

	buf = append(buf, startBytes...)
	buf = append(buf, []byte{
		0x05,
		byte(packetType),
	}...)

	buf = binary.BigEndian.AppendUint16(buf, uint16(serial))
	buf = binary.BigEndian.AppendUint16(buf, protocol.CalculateCRC(buf[2:]))

	buf = append(buf, stopBytes...)

	return buf, nil
}
