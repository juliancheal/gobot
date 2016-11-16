package ble

import (
	"fmt"

	"github.com/hybridgroup/gobot"
)

var _ gobot.Driver = (*YeelightDriver)(nil)

type YeelightDriver struct {
	name          string
	connection    gobot.Connection
	seq           uint8
	packetChannel chan *packet
	gobot.Eventer
}

const (
	// service IDs
	YeelightService = "fff0"

	// characteristic IDs
	ControlCharacteristic             = "fff1"
	DelayCharacteristic               = "fff2"
	DelayQueryCharacteristic          = "fff3"
	DelayResponseCharacteristic       = "fff4"
	StatusQueryCharacteristic         = "fff5"
	StatusQueryResponseCharacteristic = "fff6"
	ColorFlowCharacteristic           = "fff7"
	LEDNameCharacteristic             = "fff8"
	LEDNameResponseCharacteristic     = "fff9"
	EffectSettingCharacteristic       = "fffc"
)

// NewYeelightDriver creates a YeelightDriver by name
func NewYeelightDriver(a *BLEClientAdaptor, name string) *YeelightDriver {
	n := &YeelightDriver{
		name:          name,
		connection:    a,
		Eventer:       gobot.NewEventer(),
		packetChannel: make(chan *packet, 1024),
	}

	return n
}
func (b *YeelightDriver) Connection() gobot.Connection { return b.connection }
func (b *YeelightDriver) Name() string                 { return b.name }

// adaptor returns BLE adaptor
func (b *YeelightDriver) adaptor() *BLEClientAdaptor {
	return b.Connection().(*BLEClientAdaptor)
}

// Start tells driver to get ready to do work
func (s *YeelightDriver) Start() (errs []error) {
	s.Init()

	return
}

// Halt stops Ollie driver (void)
func (b *YeelightDriver) Halt() (errs []error) {
	return
}

// TODO: Add all Yeelight Responses
func (b *YeelightDriver) Init() (err error) {
	// subscribe to Yeelight response notifications
	b.adaptor().Subscribe(ControlCharacteristic, ResponseCharacteristic, b.HandleResponses)

	return
}

// Turns on Yeelight
func (b *YeelightDriver) On() (err error) {
	// TODO: Extract to craftPacket
	buf := []byte(",,,100,,,,,,,,,,,,")

	err = b.adaptor().WriteCharacteristic(YeelightService, ControlCharacteristic, buf)
	if err != nil {
		fmt.Println("On error:", err)
		return err
	}

	return
}

// Turns off Yeelight
func (b *YeelightDriver) Off() (err error) {
	// TODO: Extract to craftPacket
	buf := []byte(",,,0,,,,,,,,,,,,,,")

	err = b.adaptor().WriteCharacteristic(YeelightService, ControlCharacteristic, buf)
	if err != nil {
		fmt.Println("Off error:", err)
		return err
	}

	return
}

// Handle responses returned from Ollie
func (b *YeelightDriver) HandleResponses(data []byte, e error) {
	fmt.Println("response data:", data)

	return
}

// SetRGB sets the Yeelight to the given r, g, and b values
func (s *YeelightDriver) SetRGB(r uint8, g uint8, b uint8) (err error) {
	// s.packetChannel <- s.craftPacket([]uint8{r, g, b, 0x01}, 0x02, 0x20)
	// TODO: Extract to craftPacket
  colour := fmt.Sprintf("%d,%d,%d,%d", r, g, b, 100)

  for i := len(colour); i < 18; i++ {
    colour += ","
  }

  buf := []byte(colour)

  err = s.adaptor().WriteCharacteristic(YeelightService, ControlCharacteristic, buf)
  if err != nil {
    fmt.Println("SetRGB error:", err)
    return err
  }

  return
}

// Go to sleep
func (s *YeelightDriver) Sleep() {
	// s.packetChannel <- s.craftPacket([]uint8{0x00, 0x00, 0x00, 0x00, 0x00}, 0x00, 0x22)
}

// func (s *YeelightDriver) craftPacket(body []uint8, did byte, cid byte) *packet {
// 	packet := new(packet)
// 	packet.body = body
// 	dlen := len(packet.body) + 1
// 	packet.header = []uint8{0xFF, 0xFF, did, cid, s.seq, uint8(dlen)}
// 	packet.checksum = s.calculateChecksum(packet)
// 	return packet
// }
