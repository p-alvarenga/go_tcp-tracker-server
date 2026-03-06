package protocol

func BCDToASCII(bcd []byte) string {
	var imei []byte

	for _, b := range bcd {
		h := b >> 4
		l := b & 0x0F

		imei = append(imei, h+'0', l+'0')
	}

	return string(imei)
}
