package main

import (
	"strings"
)

type WarehouseService struct {
	r *WarehouseRepository
}

func NewWarehouseService(repo *WarehouseRepository) *WarehouseService {
	return &WarehouseService{r: repo}
}

func (bs *WarehouseService) registerNewWarehouse(b *Warehouse) {
	bs.r.AddNewWarehouse(b)
}
func (bs *WarehouseService) registerNewCustomer(b *Customer) {
	bs.r.AddNewWarehouse(b)
}

func (bs *WarehouseService) getAllWarehouseCollection() []*Warehouse {
	return bs.r.FindAllWarehouse()
}

func (bs *WarehouseService) searchWarehouseById(kode string) []*Warehouse {
	var funcFilter = func(b Warehouse) bool {
		return b.Kode == kode
	}
	return bs.r.FindWarehouseByField(funcFilter)
}

func (bs *WarehouseService) searchWarehouseByTitle(name string) []*Warehouse {
	var funcFilter = func(b Warehouse) bool {
		return strings.Contains(b.Name, name)
	}
	return bs.r.FindWarehouseByField(funcFilter)
}
