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

	for value, _ /* entropy */ := range dataset.counts[dataset.gains.Max()] {
		fmt.Print("|", strings.Repeat("	", depth*3), "|-")
		fmt.Println(dataset.gains.Max(), " => ", value)
		id3(dataset.SubDataset(value), depth+1)
	}
	return
}
