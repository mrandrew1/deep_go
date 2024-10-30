package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

const (
	OpenModeIn     = 1 // 0000 0001
	OpenModeOut    = 2 // 0000 0010
	OpenModeAppend = 4 // 0000 0100
	OpenModeBinary = 8 // 0000 1000

	// sugar for prepared masks
	OpenModeInAndOut = OpenModeIn | OpenModeOut // 0000 0001 + 0000 0010 = 0000 0011
)

func Open(filename string, mask int8) {
	if mask&OpenModeIn == 1 {
		fmt.Println("in mode")
	}
	if mask&OpenModeOut == 1 {
		fmt.Println("out mode")
	}
	if mask&OpenModeAppend == 1 {
		fmt.Println("append mode")
	}
	if mask&OpenModeBinary == 1 {
		fmt.Println("binary mode")
	}

	// implementation...
}
func main() {
	var number uint32 = 0xFFCCDDAA
	result := ToLittleEndian(number)
	fmt.Println(`result`)
	fmt.Println(strconv.FormatInt(int64(result), 16))
}

func ToLittleEndian(number uint32) uint32 {
	pointer := unsafe.Pointer(&number)
	// pointer2 := unsafe.Add(pointer, 1)
	// pointer3 := unsafe.Add(pointer, 2)
	// pointer4 := unsafe.Add(pointer, 3)
	fmt.Println(strconv.FormatInt(int64(uint32(*(*uint8)(unsafe.Add(pointer, 2)))<<8), 16))
	return uint32(*(*uint8)(unsafe.Add(pointer, 3))) | uint32(*(*uint8)(unsafe.Add(pointer, 2)))<<8 | uint32(*(*uint8)(unsafe.Add(pointer, 1)))<<16 | uint32(*(*uint8)(pointer))<<24
}
