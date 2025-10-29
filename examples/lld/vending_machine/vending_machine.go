package main

import (
	"fmt"
	"sync"
)

type Dispenser struct{}

func (d *Dispenser) DispenseItem(item *Item) {
	fmt.Println("Dispensing item:", item.code)
}

type VendingMachine struct {
	inventory       *Inventory
	paymentStrategy PaymentStrategy
	display         DisplayStrategy
	dispenser       Dispenser
	currItemCode    string
	sync.Mutex
}

func DefaultVendingMachine() *VendingMachine {
	return &VendingMachine{
		inventory:       NewInventory(),
		paymentStrategy: NewCoinPayment(),
		display:         NewDisplayStrategy(),
		dispenser:       Dispenser{},
	}
}
func NewVendingMachine(inventory *Inventory, paymentStrategy PaymentStrategy) *VendingMachine {
	return &VendingMachine{
		inventory:       inventory,
		paymentStrategy: paymentStrategy,
		dispenser:       Dispenser{},
	}
}

func (v *VendingMachine) DisplayItems() {
	for _, i := range v.inventory.item {
		fmt.Printf("Item: %s, Price: %d", i.name, i.price)
	}
}

func (v *VendingMachine) SelectItem(code string) {
	v.Lock()
	defer v.Unlock()

	if !v.inventory.HasStock(code) {
		v.display.Display("Item out of stock")
		return
	}

	v.currItemCode = code
}

func (v *VendingMachine) ProcessPayment(code string, quantity int) error {
	v.Lock()
	defer v.Unlock()
	item, isExist := v.inventory.GetItem(code)
	if !isExist {
		v.display.Display("Item not found")
	}
	if item.quantity < quantity {
		v.display.Display("Item out of stock")
	}
	v.currItemCode = code
	_, err := v.paymentStrategy.ProcessAmount(item.price)
	if err != nil {
		return err
	}
	return nil
}

func (v *VendingMachine) DispenseItem(item *Item) {
	v.dispenser.DispenseItem(item)
}
func (v *VendingMachine) DisplayMsg(message string) {
	v.display.Display(message)
}

func main() {
	vm := DefaultVendingMachine()

	// Add Items
	vm.inventory.AddItem(Item{
		code:     "A",
		name:     "item-1",
		quantity: 10,
		price:    100,
	})

	// display to choose item
	vm.DisplayMsg("Choose item")
	vm.DisplayItems()
	vm.paymentStrategy.AddAmount(300)

	vm.SelectItem("A")
	if err := vm.ProcessPayment("A", 10); err != nil {
		vm.DisplayMsg(fmt.Sprintf("Error processing payment %v", err))
		return
	}

	vm.DispenseItem(&Item{code: "A"})

}
