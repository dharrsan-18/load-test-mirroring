package metrics

import (
	"sync/atomic"
)

type Metrics struct {
	ReceivedCount  atomic.Int64
	ProcessedCount atomic.Int64
	ExportedCount  atomic.Int64
	DroppedCount   atomic.Int64
}

var GlobalMetrics Metrics

func IncrementReceived() int64 {
	return GlobalMetrics.ReceivedCount.Add(1)
}

func IncrementProcessed() int64 {
	return GlobalMetrics.ProcessedCount.Add(1)
}

func IncrementExported() int64 {
	return GlobalMetrics.ExportedCount.Add(1)
}

func IncrementDropped() int64 {
	return GlobalMetrics.DroppedCount.Add(1)
}

func GetMetrics() map[string]int64 {
	return map[string]int64{
		"received":  GlobalMetrics.ReceivedCount.Load(),
		"processed": GlobalMetrics.ProcessedCount.Load(),
		"exported":  GlobalMetrics.ExportedCount.Load(),
		"dropped":   GlobalMetrics.DroppedCount.Load(),
	}
}
