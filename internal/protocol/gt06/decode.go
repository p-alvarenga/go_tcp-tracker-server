package gt06

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func Decode(raw []byte) (Packet, error) {
	err := validateFrame(raw)
	if err != nil {
		return nil, err
	}

	payloadEnd := int(raw[2]) - 1 // same as length + 4
	payload := raw[4:payloadEnd]

	serial := binary.BigEndian.Uint16(raw[payloadEnd+1 : payloadEnd+3]) // [ n+1, n+3 ) => { n+1, n+2 }

	switch PacketType(raw[3]) {
	case LoginType:
		pkt, err := decodeLogin(payload)
		if err != nil {
			return nil, err
		}

		pkt.Serial = int(serial) // validation

		return pkt, nil

	default: // same as IsValid()
		return nil, fmt.Errorf("packet type invalid or not supported yet (%X)", raw[3])
	}
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
