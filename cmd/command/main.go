package main

import (
	"fmt"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/delivery_service"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/receiver"
)

var (
	mcdonaldsName = "mcdonalds"
	mcdonaldsMenu = map[string]bool{
		"Big Mac":   true,
		"Coca cola": true,
	}
	kfcName = "kfc"
	kfcMenu = map[string]bool{
		"Chicken":   true,
		"Coca cola": true,
	}
	schoolCanteenName = "school canteen"
	schoolCanteenMenu = map[string]bool{
		"Mashed potatoes": true,
		"Cutlets":         true,
		"Сompote":         true,
	}
)

func main() {
	mcdonalds := receiver.NewRestaurant(mcdonaldsName, mcdonaldsMenu)
	kfc := receiver.NewRestaurant(kfcName, kfcMenu)
	schoolCanteen := receiver.NewRestaurant(schoolCanteenName, schoolCanteenMenu)

	yandexEda := delivery_service.NewDeliveryService(map[string]receiver.Receiver{
		mcdonaldsName:     mcdonalds,
		kfcName:           kfc,
		schoolCanteenName: schoolCanteen,
	})

	// Выполняем команды
	err := yandexEda.RequestMenus()
	if err != nil {
		fmt.Println("Can't request menus")
	}
	fmt.Println()
	err = yandexEda.MakeOrder(schoolCanteenName, "Сompote")
	if err != nil {
		fmt.Println("Can't make order")
	}
}
