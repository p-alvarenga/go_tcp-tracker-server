package device

import (
	"time"
)

type Imei string

type Device struct {
	Imei     Imei
	LastSeen time.Time
}

func newDevice(imei Imei) *Device {
	return &Device{
		Imei: imei,
	}
}
