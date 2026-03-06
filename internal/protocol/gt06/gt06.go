package gt06

type PacketType byte

const (
	LoginType    PacketType = 0x01
	LocationType PacketType = 0x12
)

var (
	startBytes = []byte{0x78, 0x78}
	stopBytes  = []byte{0x0A, 0x0D}
)

type Packet interface {
	Type() PacketType
}
