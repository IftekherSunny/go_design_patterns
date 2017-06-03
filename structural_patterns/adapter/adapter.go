package adapter

////////////////////////////////////////////
// Stripe gateway implementation
////////////////////////////////////////////
type Stripe struct{}

// accepting payment
func (s *Stripe) Payment() string {
	return "Payment has been accepted by the stripe gateway"
}

////////////////////////////////////////////
// Paypal gateway implementation
////////////////////////////////////////////
type Paypal struct{}

// accepting payment
func (p *Paypal) AcceptPayment() string {
	return "Payment has been accepted by the paypal gateway"
}

////////////////////////////////////////////
// Payment contract
////////////////////////////////////////////
type Payment interface {
	Pay() string
}

////////////////////////////////////////////
// Stripe adapter implementation
////////////////////////////////////////////
type StripeAdapter struct {
	stripe Stripe
}

// accepting payment ( implementing pay method of the payment struct )
func (s *StripeAdapter) Pay() string {
	return s.stripe.Payment()
}

////////////////////////////////////////////
// Paypal adapter implementation
////////////////////////////////////////////
type PaypalAdapter struct {
	paypal Paypal
}

// accepting payment ( implementing pay method of the payment struct )
func (p *PaypalAdapter) Pay() string {
	return p.paypal.AcceptPayment()
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
