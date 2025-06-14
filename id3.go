package main

import (
	"fmt"
	"strings"
)

func main() {
	dataset := Load("data/baseball.csv")

	id3(dataset, 0)
}

func id3(dataset *datasetBase, depth int) (subdataset *datasetBase) {
	dataset.CalculateCounts()
	if dataset.isHaveOneResult() {
		fmt.Print("|", strings.Repeat("	", depth*3), "|->> ")

		_, value := dataset.ResultClass()
		fmt.Println(value)
	}

	dataset.CalculateEntropies()
	dataset.CalculateGains()

	maxGain:=dataset.gains.Max()
	for value, _ /* entropy */ := range dataset.counts[maxGain] {
		fmt.Print("|", strings.Repeat("	", depth*3), "|-")

		fmt.Println(maxGain, " => ", value)
		id3(dataset.SubDataset(maxGain,value), depth+1)
	}
	return
}
