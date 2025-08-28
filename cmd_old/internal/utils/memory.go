package utils

import (
	"log"

	"github.com/shirou/gopsutil/mem"
)

func PrintMemUsage() {
	v, _ := mem.VirtualMemory()

	log.Printf("RAM  | Total: %v MB | Used: %v MB | Free: %v MB | In Used: %.2f %%", byteToMByte(v.Total), byteToMByte(v.Used), byteToMByte(v.Free), v.UsedPercent)
}
