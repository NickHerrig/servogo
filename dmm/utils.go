package dmm 

// checksumByte returns returns a checksum byte in dmm format
func checksumByte(p []byte) byte {
	var cs byte
	for _, v := range p {
		cs += v
	}
	return (cs & 0x7f) | 0x80
}
