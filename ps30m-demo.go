//go:build !tinygo && !rpi

package ps30m

type transport struct {
	start uint16
	words uint16
}

func newTransport() *transport {
	return &transport{}
}

func (t *transport) Write(buf []byte) (n int, err error) {
	// get start and words from Modbus request
	t.start = (uint16(buf[2]) << 8) | uint16(buf[3])
	t.words = (uint16(buf[4]) << 8) | uint16(buf[5])
	return n, nil
}

func (t *transport) Read(buf []byte) (n int, err error) {
	// simluate a Modbus request read on the device
	res := buf[3:]
	switch t.start {
	case regVerSw:
	case regAdcIcFShadow:
		// TODO make this more dynamic using a little bit of random
		copy(res[2:4], unf16(1.3))   // solar.amps
		copy(res[4:6], unf16(14.1))  // battery.volts
		copy(res[6:8], unf16(11.3))  // solar.volts
		copy(res[8:10], unf16(0))    // load.volts
		copy(res[10:12], unf16(3.3)) // battery.amps
		copy(res[12:14], unf16(0))   // load.amps
	}
	return int(5 + t.words*2), nil
}
