package product

// Product is common interface of products
type Product = interface {
	// Forward commands product to move forwards
	Forward()
	// Back commands product to move back
	Back()
}
