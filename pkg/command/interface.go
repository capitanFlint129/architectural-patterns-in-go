package command

type restaurant interface {
	GiveMenu()
	CookOrder(orderData string)
}

type Command interface {
	Execute()
}
