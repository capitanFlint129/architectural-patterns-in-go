package command

type restaurant interface {
	GiveMenu()
	CookOrder(orderData string)
}

// Command - some task for receiver
type Command interface {
	// Execute - executes command via receiver's method
	Execute()
}
