package hexdump

import "testing"
import "fmt"

//这个功能的实现还有问题。
func TestHex2str(t *testing.T) {
	content := "88"
	str := Hex2b(content)
	fmt.Printf("%x",str)
}



