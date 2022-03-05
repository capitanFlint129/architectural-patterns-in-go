package main

import (
	"fmt"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/command"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/receiver"
)

var orderData = "Big Mac"

// Отправитель
func main() {
	// Создаем получателя
	mcdonalds := receiver.NewRestaurant()

	// Создаем команды
	requestMenu := command.NewRequestMenu(mcdonalds)
	makeOrder := command.NewMakeOrder(mcdonalds, orderData)

	// Выполняем команды
	err := requestMenu.Execute()
	if err != nil {
		fmt.Println("Can't request menu")
	}
	err = makeOrder.Execute()
	if err != nil {
		fmt.Println("Can't make order")
	}
}
