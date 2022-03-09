package delivery_service

import (
	"github.com/sirupsen/logrus"
	"sync"

	"github.com/capitanFlint129/architectural-patterns-in-go/command/pkg/command"
)

type restaurant = interface {
	GiveMenu() error
	CookOrder(dish string) error
}

// DeliveryService organizes delivery from several restaurants
type DeliveryService interface {
	RequestMenus() error
	MakeOrder(restaurantName string, dish string) error
}

type deliveryService struct {
	restaurants map[string]restaurant
}

// RequestMenus request menus from all restaurants
func (d *deliveryService) RequestMenus() error {
	logrus.Info("Delivery service: RequestMenus executes")

	wg := &sync.WaitGroup{}
	for _, restaurant := range d.restaurants {
		wg.Add(1)
		requestMenuCommand := command.NewRequestMenu(restaurant)

		go func() {
			defer wg.Done()
			requestMenuCommand.Execute()
		}()
	}
	wg.Wait()

	logrus.Info("Delivery service: menus given")
	return nil
}

// MakeOrder - orders the specified dish at the specified restaurant
func (d *deliveryService) MakeOrder(restaurantName string, dish string) error {
	logrus.Info("Delivery service: MakeOrder executes")
	err := d.restaurants[restaurantName].CookOrder(dish)
	return err
}

// NewDeliveryService - creates new delivery service
func NewDeliveryService(restaurantMap map[string]restaurant) DeliveryService {
	return &deliveryService{restaurants: restaurantMap}
}
