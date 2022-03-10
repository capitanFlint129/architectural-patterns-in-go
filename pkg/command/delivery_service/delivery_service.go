package delivery_service

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type command = interface {
	Execute() error
}

type restaurant = interface {
	GiveMenu() error
	CookOrder(dish string) error
}

type requestMenuCommandCreator func(restaurant restaurant) command
type cookOrderCommandCreator func(restaurant restaurant, dish string) command

// DeliveryService organizes delivery from several restaurants
type DeliveryService interface {
	RequestMenus() error
	MakeOrder(restaurantName string, dish string) error
}

type deliveryService struct {
	restaurants               map[string]restaurant
	requestMenuCommandCreator requestMenuCommandCreator
	cookOrderCommandCreator   cookOrderCommandCreator
}

// RequestMenus request menus from all restaurants
func (d *deliveryService) RequestMenus() error {
	logrus.Info("Delivery service: RequestMenus executes")

	wg := &sync.WaitGroup{}
	for _, restaurant := range d.restaurants {
		requestMenuCommand := d.requestMenuCommandCreator(restaurant)

		wg.Add(1)
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
	makeOrderCommand := d.cookOrderCommandCreator(d.restaurants[restaurantName], dish)
	err := makeOrderCommand.Execute()
	return err
}

// NewDeliveryService - creates new delivery service
func NewDeliveryService(
	restaurantMap map[string]restaurant,
	requestMenuCommandCreator requestMenuCommandCreator,
	cookOrderCommandCreator cookOrderCommandCreator,
) DeliveryService {
	return &deliveryService{
		restaurants:               restaurantMap,
		requestMenuCommandCreator: requestMenuCommandCreator,
		cookOrderCommandCreator:   cookOrderCommandCreator,
	}
}
