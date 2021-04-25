package unicodemisc

const hextable string = "0123456789abcdef"

// UnicodeEscapeSequence converts a single byte into bytes sequence of Unicode Escape Sequence.
func UnicodeEscapeSequence(singleByte byte) []byte {
	s := make([]byte, 6)
	// https://github.com/golang/go/blob/70deaa33ebd91944484526ab368fa19c499ff29f/src/encoding/hex/hex.go#L28-L29
	s[0], s[1], s[2], s[3], s[4], s[5] = '\\', 'u', '0', '0', hextable[singleByte>>4], hextable[singleByte&0x0f]
	return s
}

// UnicodeEscapeSequenceString converts a single byte into string of Unicode Escape Sequence.
func UnicodeEscapeSequenceString(singleByte byte) string {
	return string(UnicodeEscapeSequence(singleByte))
}

// ApppendUnicodeEscapeSequence converts a single byte into bytes sequence of Unicode Escape Sequence and appends to dst.
func ApppendUnicodeEscapeSequence(dst []byte, singleByte byte) []byte {
	if dst == nil {
		dst = make([]byte, 0, 6)
	}
	// https://github.com/golang/go/blob/70deaa33ebd91944484526ab368fa19c499ff29f/src/encoding/hex/hex.go#L28-L29
	dst = append(dst, '\\', 'u', '0', '0', hextable[singleByte>>4], hextable[singleByte&0x0f])
	return dst
}
