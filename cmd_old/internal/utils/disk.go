package utils

import (
	"log"

	"github.com/shirou/gopsutil/disk"
)

func PrintDiskUsage(path string) {
	v, _ := disk.Usage(path)

	log.Printf("DISK | Total: %v MB | Used: %v MB | Free: %v MB | In Used: %.2f %%", byteToMByte(v.Total), byteToMByte(v.Used), byteToMByte(v.Free), v.UsedPercent)
}
