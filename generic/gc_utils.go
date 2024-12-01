package generic

import (
	"fmt"
	"runtime"
	"time"
)

const (
	MB = 1024 * 1024
	GB = 1024 * MB
)

func GetMemoryStats() *runtime.MemStats {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	return &memStats
}

func FormatMemorySize(bytes uint64) string {
	if bytes >= GB {
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
	}
	return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
}

func CheckMemoryAndRunGCDuration(sleepTime time.Duration, memoryInGb uint64) {
	if ms := GetMemoryStats(); ms.HeapInuse > memoryInGb*GB {
		runtime.GC()
		time.Sleep(sleepTime)
	}
}

func CheckMemoryAndRunGC(sleepTimeSec int, memoryInGb uint64) {
	CheckMemoryAndRunGCDuration(time.Duration(sleepTimeSec)*time.Second, memoryInGb)
}

func CheckMemoryAndRunGCInMs(sleepTimeMilli int, memoryInGb uint64) {
	CheckMemoryAndRunGCDuration(time.Duration(sleepTimeMilli)*time.Millisecond, memoryInGb)
}
