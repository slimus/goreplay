package main

import (
	"fmt"
)

const (
	outputDummyStatRequestCount = "output_dummy.count"
)

// DummyOutput used for debugging, prints all incoming requests
type DummyOutput struct {
	statistic statisticCollector
}

// NewDummyOutput constructor for DummyOutput
func NewDummyOutput(_ string, statistic statisticCollector) *DummyOutput {
	return &DummyOutput{
		statistic: statistic,
	}
}

func (do *DummyOutput) Write(data []byte) (int, error) {
	fmt.Println(string(data))

	do.statistic.Incr(outputDummyStatRequestCount)

	return len(data), nil
}

func (do *DummyOutput) String() string {
	return "Dummy Output"
}
