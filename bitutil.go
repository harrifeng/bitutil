package bitutil

import (
	"fmt"
	"strconv"
	"strings"
)

func TwoComplementInt8ToRaw(num int8) string {
	var ret = []byte("00000000")

	for i := range ret {
		// we need a higher precision for the highest bit
		if ((1 << uint16(i)) & uint16(num)) > 0 {
			ret[i] = '1'
		}
	}
	// reverse the byte array
	beg := 0
	end := len(ret) - 1
	for beg <= end {
		ret[beg], ret[end] = ret[end], ret[beg]
		beg++
		end--
	}
	return string(ret)
}

func HexStringToBinaryString(x string) string {
	var ret []string

	for i := 0; i < len(x); i++ {
		curStr := x[i : i+1]
		ret = append(ret, hexToBinary(curStr))
	}
	return strings.Join(ret, "|")
}

func hexToBinary(x string) string {
	base, _ := strconv.ParseInt(x, 16, 10)
	return fmt.Sprintf("%04s", strconv.FormatInt(base, 2))
}
