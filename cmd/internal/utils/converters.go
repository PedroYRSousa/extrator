package utils

func byteToGByte(b uint64) uint64 {
	return b / 1024 / 1024 / 1024
}

func byteToMByte(b uint64) uint64 {
	return b / 1024 / 1024
}

func byteToKByte(b uint64) uint64 {
	return b / 1024
}
