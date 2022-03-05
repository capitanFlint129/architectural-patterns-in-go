package command

type restaurant interface {
	GiveMenu() error
	CookOrder(orderData string) error
}

// Command - some task for receiver
type Command interface {
	// Execute - executes command via receiver's method
	Execute() error
}
