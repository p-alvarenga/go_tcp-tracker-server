package gt06

import (
	"bytes"
	"fmt"
)

func Decode(raw []byte) (Packet, error) {
	err := validateFrame(raw)
	if err != nil {
		return nil, err
	}

	payloadLength := int(raw[2]) - 5
	packetType := raw[3]
	payload := raw[4 : payloadLength+4]

	switch PacketType(packetType) {
	case LoginType:

	}

	return nil, nil
}

func validateFrame(raw []byte) error {
	if len(raw) < 10 {
		return fmt.Errorf("gt06: packet too small (%d bytes)", len(raw))
	}

	if !bytes.HasPrefix(raw, startBytes) {
		return fmt.Errorf("gt06: invalid header")
	}

	if !bytes.HasSuffix(raw, stopBytes) {
		return fmt.Errorf("gt06: invalid stop")
	}

	return nil
}
