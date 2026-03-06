package device

import "sync"

type DeviceManager struct {
	mu      sync.RWMutex
	devices map[Imei]*Device
}

func NewDeviceManager() *DeviceManager {
	return &DeviceManager{
		devices: make(map[Imei]*Device),
	}
}

func (dm *DeviceManager) GetOrCreate(imei Imei) *Device {
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
