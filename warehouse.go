package main

type Warehouse struct {
	Kode        string
	Name        string
	Address     string
	Large       float64
	Information string
	Price       float64
}

type Customer struct {
	Id           string
	Name         string
	Goods        string
	LargeUsed    float64
	DateOfEntry  int64
	TotalDayRent int64
	TotalPrice   float64
}

func NewWarehouse(name string, address string, large float64, information string, price float64) Warehouse {
	return Warehouse{
		Name:        name,
		Address:     address,
		Large:       large,
		Information: information,
		Price:       price,
	}
}

func NewCustomer(name string, goods string, largeUsed float64, dateOfEntry int64, totalDayRent int64, totalPrice float64) Customer {
	return Customer{
		Name:         name,
		Goods:        goods,
		LargeUsed:    largeUsed,
		DateOfEntry:  dateOfEntry,
		TotalDayRent: totalDayRent,
		TotalPrice:   totalPrice,
	}
}
