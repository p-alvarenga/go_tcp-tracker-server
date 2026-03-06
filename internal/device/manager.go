package device

import (
	"sync"

	"github.com/p-alvarenga/go_tcp-tracker-server/internal/domain/types"
)

type DeviceManager struct {
	mu      sync.RWMutex
	devices map[types.IMEI]*Device
}

func NewDeviceManager() *DeviceManager {
	return &DeviceManager{
		devices: make(map[types.IMEI]*Device),
	}
}

func (dm *DeviceManager) GetOrCreate(imei types.IMEI) *Device {
	dm.mu.RLock()
	d, ok := dm.devices[imei]
	dm.mu.RUnlock()

	if ok {
		return d
	}

	dm.mu.Lock()
	defer dm.mu.Unlock()

	d, ok = dm.devices[imei] // double checking
	if ok {
		return d
	}

	d = newDevice(imei)
	dm.devices[imei] = d

	return d
}
