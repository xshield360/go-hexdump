package hexdump

import "testing"

func TestHexdump(t *testing.T) {
	content := "AAAAAAAABBBBBBBBCCCCCCCC\x90\x00\x88\xcc"
	Hexdump(content,8)
}

func TestHexdump2(t *testing.T) {
	content := "\x88"
	Hexdump(content,8)
}
