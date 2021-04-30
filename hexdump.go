package hexdump

import (
	"fmt"
)

func Hexdump(content string, n int) {
	var printOut string
	if n == 0 {
		n = 16
	}
	mod := len(content)/n + 1
	iv := len(content)%n
	for i := 0 ; i < mod; i ++  {
		str0 := fmt.Sprintf("%.8x",i*n)
		str1 := ""
		str2 := ""
		jNum := 0
		if i == mod - 1 {
			jNum = iv
		} else {
			jNum = n
		}
		for j := 0; j < jNum; j ++ {
			str1 += fmt.Sprintf("%.2x ", content[i*n+j])
			if content[i*n+j] < 31 || content[i*n+j]> 127{
				str2 += ". "
			} else {
				str2 += fmt.Sprintf("%c ",content[i*n+j])
			}
		}
		for j := 0; j < (n-jNum)%n; j ++{
			str1 += "   "
		}
		if len(str1) > 0 {
			printOut += fmt.Sprintf("%s %s %s\n", str0, str1, str2)
		}
	}
	fmt.Print(printOut)
}
