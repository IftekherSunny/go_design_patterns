package bridge

////////////////////////////////////////////
// Payment contract
////////////////////////////////////////////
type Payment interface {
	Pay() string
}

////////////////////////////////////////////
// Stripe gateway implementation
////////////////////////////////////////////
type Stripe struct{}

// accepting payment
func (s *Stripe) Pay() string {
	return "Payment has been accepted by the stripe gateway"
}

////////////////////////////////////////////
// Paypal gateway implementation
////////////////////////////////////////////
type Paypal struct{}

// accepting payment
func (p *Paypal) Pay() string {
	return "Payment has been accepted by the paypal gateway"
}

////////////////////////////////////////////
// Shopping cart
////////////////////////////////////////////
type ShoppingCart struct {
	Payment Payment
}

// accepting shopping card payment
func (s *ShoppingCart) Checkout(gatewayName string) string {
	switch gatewayName {
	case "paypal":
		return s.Payment.Pay()
	case "stripe":
		return s.Payment.Pay()
	default:
		return "Select paypal or stripe gateway"
	}
}
