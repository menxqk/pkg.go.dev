package main

import (
	"fmt"
	"runtime/metrics"
)

func main() {
	// Name of the metric we want to read.
	const myMetric = "/memory/classes/heap/free:bytes"

	// Create a sample for the metric.
	sample := make([]metrics.Sample, 1)
	sample[0].Name = myMetric

	metrics.Read(sample)

	// Check if the metric is actually supported.
	// If it's not, the resulting value will always have
	// kind KindBad.
	if sample[0].Value.Kind() == metrics.KindBad {
		panic(fmt.Sprintf("metric %q no longer supported", myMetric))
	}

	// Handle the result.
	//
	// It's OK to assume a particular kind for a metric;
	// they're guaranteed not to change.
	freeBytes := sample[0].Value.Uint64()

	fmt.Printf("free but not releases memory: %d\n", freeBytes)
}
