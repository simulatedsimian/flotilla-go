package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type Matrix struct {
	ModuleCommon

	buffer     [8]byte
	brightness byte
	dirty      bool
}

func (m *Matrix) Construct() {
	m.brightness = 64
	m.dirty = true
}

func (m *Matrix) Set(d *dock.Dock) error {
	if m.dirty && m.address != nil {
		m.dirty = false
		return d.SetModuleData(m.address.channel, m.Type(),
			int(m.buffer[0]), int(m.buffer[1]), int(m.buffer[2]), int(m.buffer[3]),
			int(m.buffer[4]), int(m.buffer[5]), int(m.buffer[6]), int(m.buffer[7]),
			int(m.brightness))
	}
	return nil
}

func (m *Matrix) Type() dock.ModuleType {
	return dock.Matrix
}

func (m *Matrix) SetBrightness(b int) {
	m.brightness = byte(b)
	m.dirty = true
}

func (m *Matrix) Plot(x, y, v int) {
	x = 7 - x&7
	y = y & 7

	if v == 0 {
		m.buffer[x] = m.buffer[x] & ^(1 << uint(y))
	} else {
		m.buffer[x] = m.buffer[x] | (1 << uint(y))
	}

	m.dirty = true
}

func (m *Matrix) DrawBarGraph(values []int, min, max int) {
	m.Clear()

	bars := MinInt(8, len(values))

	for i := 0; i < bars; i++ {
		y := Map(values[i], min, max, 0, 7)
		for n := 0; n <= y; n++ {
			m.Plot(i, 7-n, 1)
		}
	}
}

func (m *Matrix) Clear() {
	m.buffer = [8]byte{}
}

func (m *Matrix) Scroll(dir Direction, fill int) {
	if dir&DirLeft != 0 {
		m.ScrollLeft(fill)
	}
	if dir&DirRight != 0 {
		m.ScrollRight(fill)
	}
	if dir&DirUp != 0 {
		m.ScrollUp(fill)
	}
	if dir&DirDown != 0 {
		m.ScrollDown(fill)
	}
}

func (m *Matrix) ScrollRight(fill int) {
	copy(m.buffer[:], m.buffer[1:])
	m.buffer[7] = byte(fill)
	m.dirty = true
}

func (m *Matrix) ScrollLeft(fill int) {
	copy(m.buffer[1:], m.buffer[:])
	m.buffer[0] = byte(fill)
	m.dirty = true
}

func (m *Matrix) ScrollDown(fill int) {
	for i := range m.buffer {
		m.buffer[i] = (m.buffer[i] << 1) | (byte(fill)>>byte(7-i))&1
	}
	m.dirty = true
}

func (m *Matrix) ScrollUp(fill int) {
	for i := range m.buffer {
		m.buffer[i] = (m.buffer[i] >> 1) | ((byte(fill)>>byte(7-i))&1)<<7
	}
	m.dirty = true
}

func (m *Matrix) RollUp() {
	m.dirty = true
}

func (m *Matrix) RollDown() {
	m.dirty = true
}

func (m *Matrix) RollLeft() {
	m.dirty = true
}

func (m *Matrix) RollRight() {
	m.dirty = true
}
