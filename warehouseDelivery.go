package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type WarehouseDelivery struct {
	warehouseService *WarehouseService
}

func NewWarehouseDelivery(c *AppConfig) *WarehouseDelivery {
	repo := NewWarehouseRepository(c.getDbPath(), c.getDB())
	warehouseService := NewWarehouseService(repo)
	return &WarehouseDelivery{warehouseService}
}

func (bd *WarehouseDelivery) Create() {
	var isExist = false
	var userChoice string

	bd.mainMenuForm()
	for isExist == false {
		fmt.Printf("\n%s", "Your Choice: ")
		fmt.Scan(&userChoice)
		switch {
		case userChoice == "01":
			bd.registrationWarehouseForm()
		case userChoice == "02":
			bd.viewWarehouseCollectionForm()
		// case userChoice == "03":
		// 	bd.rentWarehouseCollectionForm()
		case userChoice == "q":
			isExist = true
		default:
			fmt.Println("Unknown Menu Code")

		}
	}
}

func (bd *WarehouseDelivery) menuChoiceOrdered(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (bd *WarehouseDelivery) mainMenuForm() {
	var appMenu = map[string]string{
		"01": "Add warehouse to collection",
		"02": "View warehouse collection",
		"03": "Rent warehouse collection",
		"q":  "Quit aplication",
	}
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%26s\n", "My 1st Warehouse Collection")
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for _, menuCode := range bd.menuChoiceOrdered(appMenu) {
		fmt.Printf("%s. %s\n", menuCode, appMenu[menuCode])
	}
}

func (bd *WarehouseDelivery) registrationWarehouseForm() {
	consoleClear()
	var name string
	var address string
	var large float64
	var information string
	var price float64
	var confirmation string
	fmt.Println()
	fmt.Printf("%s\n", "Warehouse Registration Form")
	fmt.Printf("%s\n", strings.Repeat("-", 30))
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Name : ")
	sName, _ := scanner.ReadString('\n')
	name = strings.TrimSpace(sName)
	fmt.Print("Address : ")
	sAddress, _ := scanner.ReadString('\n')
	address = strings.TrimSpace(sAddress)
	fmt.Print("Large (/m2): ")
	sLarge, _ := scanner.ReadString('\n')
	large, _ = strconv.ParseFloat(strings.TrimSpace(sLarge), 64)
	fmt.Print("Information : ")
	sInformation, _ := scanner.ReadString('\n')
	information = strings.TrimSpace(sInformation)
	fmt.Print("Price (Rp.): ")
	sPrice, _ := scanner.ReadString('\n')
	price, _ = strconv.ParseFloat(strings.TrimSpace(sPrice), 64)

	fmt.Println("Save to collection? :Y/N")
	fmt.Scanln(&confirmation)

	if confirmation == "Y" {
		newWarehouse := NewWarehouse(name, address, large, information, price)
		bd.warehouseService.registerNewWarehouse(&newWarehouse)
	}
	consoleClear()
	bd.mainMenuForm()

}

func (bd *WarehouseDelivery) rentWarehouseCollectionForm() {
	consoleClear()
	var name string
	var goods string
	var largeUsed float64
	var dateOfEntry int64
	var totalDayRent int64
	var totalPrice float64
	var confirmation string
	fmt.Println()
	fmt.Printf("%s\n", "Rent Registration Form")
	fmt.Printf("%s\n", strings.Repeat("-", 30))
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Name : ")
	sName, _ := scanner.ReadString('\n')
	name = strings.TrimSpace(sName)
	fmt.Print("Goods : ")
	sGoods, _ := scanner.ReadString('\n')
	goods = strings.TrimSpace(sGoods)
	fmt.Print("Large Used (/m2): ")
	sLargeUsed, _ := scanner.ReadString('\n')
	largeUsed, _ = strconv.ParseFloat(strings.TrimSpace(sLargeUsed), 64)
	fmt.Print("Date Of Entry : ")
	sDateOfEntry, _ := scanner.ReadString('\n')
	dateOfEntry, _ = strconv.ParseInt(sDateOfEntry, 10, 64)
	fmt.Print("Total Rent (Days) : ")
	sTotalDayRent, _ := scanner.ReadString('\n')
	totalDayRent, _ = strconv.ParseInt(sTotalDayRent, 10, 64)
	fmt.Print("Total Price : ")
	sTotalPrice := int64(2000) * totalDayRent * int64(largeUsed)

	fmt.Println("Save to collection? :Y/N")
	fmt.Scanln(&confirmation)

	if confirmation == "Y" {
		newCustomer := NewCustomer(name, goods, largeUsed, dateOfEntry, totalDayRent, totalPrice)
		bd.warehouseService.registerNewWarehouse(&newCustomer)
	}
	consoleClear()
	bd.mainMenuForm()

}

func (bd *WarehouseDelivery) viewWarehouseCollectionForm() {
	consoleClear()
	warehouseView := bd.warehouseService.getAllWarehouseCollection()
	fmt.Println("")
	fmt.Println("My Warehouse Collection")
	fmt.Printf("%s\n", strings.Repeat("=", 115))
	fmt.Printf("%-35s%-20s%-20s%-15s%-20s%-15s\n", "Kode", "Name", "Address", "Large", "Information", "Price (Rp.)")
	fmt.Printf("%s\n", strings.Repeat("-", 115))
	for _, b := range warehouseView {
		fmt.Printf("%-35s%-20s%-20s%-15.0f%-20s%-15.2f\n", b.Kode, b.Name, b.Address, b.Large, b.Information, b.Price)
	}
	var confirmation string
	fmt.Print("Back To Menu  : (Y)")
	fmt.Scanln(&confirmation)

	if confirmation == "Y" {
		bd.mainMenuForm()
	}
}
