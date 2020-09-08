package main

import (
	"crypto/md5"
	"fmt"
)

type WarehouseRepository struct {
	warehouseCollection      *[]*Warehouse
	customerCollection       *[]*Customer
	warehouseCollectionInfra *WarehouseRepositoryInfrastructure
}

func NewWarehouseRepository(dataPath string, warehouseCollection []*Warehouse) *WarehouseRepository {
	warehouseRepoInfra := NewWarehouseRepoInfra(dataPath)
	return &WarehouseRepository{warehouseCollection: &warehouseCollection, warehouseCollectionInfra: warehouseRepoInfra}
}
func NewCustomerRepository(dataPath string, customerCollection []*Customer) *WarehouseRepository {
	warehouseRepoInfra := NewWarehouseRepoInfra(dataPath)
	return &WarehouseRepository{warehouseCollection: &warehouseCollection, warehouseCollectionInfra: warehouseRepoInfra}
}

func (br *WarehouseRepository) AddNewWarehouse(warehouse *Warehouse) {
	data := []byte(warehouse.Name)
	warehouse.Kode = fmt.Sprintf("%x", md5.Sum(data))
	*br.warehouseCollection = append(*br.warehouseCollection, warehouse)
	br.warehouseCollectionInfra.saveToFile(br.warehouseCollection)
}

func (br *WarehouseRepository) AddNewCustomer(customer *Customer) {
	data := []byte(customer.Name)
	customer.Id = fmt.Sprintf("%x", md5.Sum(data))
	*br.CustomerCollection = append(*br.customerCollection, customer)
	br.warehouseCollectionInfra.saveToFile(br.warehouseCollection)
}

func (br *WarehouseRepository) FindAllWarehouse() []*Warehouse {
	br.warehouseCollectionInfra.readFile(br.warehouseCollection)
	return *br.warehouseCollection
}

func (br *WarehouseRepository) FindWarehouseByField(fnFilter func(Warehouse) bool) []*Warehouse {
	br.warehouseCollectionInfra.readFile(br.warehouseCollection)
	var searchResult = make([]*Warehouse, 0)
	for _, b := range *br.warehouseCollection {
		if fnFilter(*b) {
			searchResult = append(searchResult, b)
		}
	}
	return searchResult
}
