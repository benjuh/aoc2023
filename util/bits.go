package util

import "bytes"

func StringToBinary(s string, true_bit rune, false_bit rune) int64 {
	var buffer bytes.Buffer
	for _, c := range s {
		if c == true_bit {
			buffer.WriteByte('1')
		} else if c == false_bit {
			buffer.WriteByte('0')
		}
	}
	binaryString := buffer.String()
	var num int64
	for i, c := range binaryString {
		if c == '1' {
			num += 1 << (len(binaryString) - 1 - i)
		}
	}
	return num
}
