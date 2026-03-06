package device

import (
	"time"

	"github.com/p-alvarenga/go_tcp-tracker-server/internal/domain/types"
)

type Device struct {
	IMEI     types.IMEI
	LastSeen time.Time
}

func newDevice(imei types.IMEI) *Device {
	return &Device{
		IMEI: imei,
	}
}
