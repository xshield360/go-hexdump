package hexdump

import (
	"strconv"
)

func Hex2b(content string) []byte {
	length := len(content)
	retBytes := make([]byte, length/2)
	j := 0
	for i := 0; i < length; i = i +2 {
		ss := string(content[i]) + string(content[i+1])
		b, _ := strconv.ParseInt(ss, 16, 32) //16进制转成int32
		retBytes[j] = byte(b)
		j ++
	}
	return retBytes
}