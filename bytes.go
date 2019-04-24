package api

func (m *Bytes) Write(p []byte) (n int, err error) {
	before := len(m.Bits)
	m.Bits = append(m.Bits, p...)
	after := len(m.Bits)
	return after - before, nil
}
