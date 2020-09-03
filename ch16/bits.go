package main

import "fmt"

var baseStrings = []string{
	"code",
	"heap",
	"stack",
	"unused",
}

var base = []int{
	32 * 1024,
	34 * 1024,
	28 * 1024,
}

var bounds = []int{
	2 * 1024,
	2 * 1024,
	2 * 1024,
}

func main() {
	// keep only bits 12 & 13
	segmentMask := 0x3000
	// keep 12 lower bits: bits 0-11
	offsetMask := 0x0fff

	virtualAddress := 100
	segmentShift := 12
	segment := (virtualAddress & segmentMask) >> segmentShift
	fmt.Printf("%02b = %s\n", segment, baseStrings[segment])
	offset := (virtualAddress & offsetMask)
	fmt.Printf("offset: %012b\n", offset)
	if offset >= bounds[segment] {
		panic("protection fault")
	}
	physicalAddress := base[segment] + offset
	fmt.Printf("physical address: %d\n", physicalAddress)
}
