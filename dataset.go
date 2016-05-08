package main

import (
	//"fmt"
	"math"
)

type datasetBase struct {
	head []string
	rows [][]string
	// because of making just one casting
	rowTotal float32
	//	      column key,value key,  type key, value
	counts    map[string]map[string]map[string]float32
	entropies map[string]float32
	gains     gainBase
}

func (d *datasetBase) Columns(no int) (columns []string) {
	columns = make([]string, 0)
	for _, row := range d.rows {
		columns = append(columns, row[no])
	}
	return
}

func (d *datasetBase) columnIdByName(name string) (id int) {
	for key, value := range d.head {
		if value == name {
			return key
		}
	}
	return -1
}

func (d *datasetBase) columnIdOfClass() (id int) {
	return len(d.head) - 1
}

func (d *datasetBase) columnNameOfClass() (name string) {
	return d.head[d.columnIdOfClass()]
}

func (d *datasetBase) classOfRow(id int) (class string) {
	return d.rows[id][d.columnIdOfClass()]
}

func (d *datasetBase) ResultClass() (column, value string) {
	return d.columnNameOfClass(), d.classOfRow(0)
}

func (d *datasetBase) entropy(count float32, sum float32) (value float32) {
	p := count / sum
	value = p * float32(math.Log2(float64(p)))
	return
}

func (d *datasetBase) isHaveOneResult() bool {
	return len(d.counts[d.columnNameOfClass()]) < 2
}

func (d *datasetBase) CalculateCounts() {
	d.rowTotal = float32(len(d.rows))
	// create var for key-value
	d.counts = make(map[string]map[string]map[string]float32)
	// get column id and name
	for columnNo, column := range d.head {
		// create var for column uniqe value
		d.counts[column] = make(map[string]map[string]float32)
		// for every row of column
		for id, row := range d.Columns(columnNo) {
			// if value detail is null then create it
			if d.counts[column][row] == nil {
				d.counts[column][row] = make(map[string]float32)
			}
			// calculate repeat count
			d.counts[column][row]["Total"]++
			// calculate count of result class
			d.counts[column][row][d.classOfRow(id)]++
		}
	}
}

func (d *datasetBase) CalculateEntropies() {
	// main class entropy
	d.entropies = make(map[string]float32)
	for _ /*key*/, count := range d.counts[d.columnNameOfClass()] {
		d.entropies["Class"] -= d.entropy(float32(count["Total"]), d.rowTotal)
	}

	// sub entropies
	for _ /*column*/, row := range d.counts {
		for _ /*class*/, cell := range row {
			for key, value := range cell {
				if key == "Total" || key == "Entropy" {
					continue
				}
				// each column entropy
				cell["Entropy"] -= d.entropy(value, cell["Total"])
			}
		}
	}

	// column-class entropies
	for key, column := range d.counts {
		if key == d.columnNameOfClass() {
			continue
		}

		for _ /*class*/, cell := range column {
			d.entropies[key] += cell["Total"] / d.rowTotal * cell["Entropy"]
		}
	}
}

func (d *datasetBase) CalculateGains() {
	d.gains = make(map[string]float32)
	// column-class gains
	for key, entropy := range d.entropies {
		if key == "Class" {
			continue
		}
		// calculate gain
		d.gains[key] = d.entropies["Class"] - entropy
	}
}

func (d *datasetBase) SubDataset(value string) (subDataset *datasetBase) {
	subDataset = new(datasetBase)
	// copy header to sub dataset
	subDataset.head = d.head

	// id of the column which has max gain
	maxGainKey := d.columnIdByName(d.gains.Max())

	for _, row := range d.rows {
		if row[maxGainKey] == value {
			// append new sub dataset
			subDataset.rows = append(subDataset.rows, row)
		}
	}
	return
}
