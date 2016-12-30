package bitutil

func GetTwoComplement8(num int8) string {
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
