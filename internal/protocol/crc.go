package protocol

func CalculateCRC(raw []byte) uint16 {
	var crc uint16 = 0xffff

	for _, b := range raw {
		crc ^= uint16(b) << 8

		for range 8 {
			if crc&0x8000 != 0 {
				crc = (crc << 1) ^ 0x1021
			} else {
				crc <<= 1
			}
		}
	}

	return crc
}

func CheckCRC(raw []byte, crc uint16) bool {
	return CalculateCRC(raw) == crc
}
