package main

import (
	"encoding/json"
	"io/ioutil"
)

type WarehouseRepositoryInfrastructure struct {
	dataPath string
}

func NewWarehouseRepoInfra(dataPath string) *WarehouseRepositoryInfrastructure {
	return &WarehouseRepositoryInfrastructure{dataPath}
}

func (bri *WarehouseRepositoryInfrastructure) saveToFile(warehouseCollection *[]*Warehouse) {
	file, _ := json.MarshalIndent(warehouseCollection, "", " ")
	_ = ioutil.WriteFile(bri.dataPath, file, 0644)
}

func (bri *WarehouseRepositoryInfrastructure) readFile(warehouseCollection *[]*Warehouse) {
	file, _ := ioutil.ReadFile(bri.dataPath)
	_ = json.Unmarshal(file, warehouseCollection)
}
