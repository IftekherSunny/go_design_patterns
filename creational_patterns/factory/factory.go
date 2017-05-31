package factory

////////////////////////////////////////////
// Supported payment gateways
///////////////////////////////////////////
const (
	STRIPE = "stripe"
	PAYPAL = "paypal"
)

////////////////////////////////////////////
// Payment method contract
////////////////////////////////////////////
type PaymentMethod interface {
	Pay(amount float32) string
}

////////////////////////////////////////////
// Payment factory
////////////////////////////////////////////
type Payment struct {
	instance PaymentMethod
}

// set payment gateway
func (p *Payment) SetPaymentGateway(gatewayName string) *Payment {
	switch gatewayName {
	case STRIPE:
		p.instance = new(Stripe)
	case PAYPAL:
		p.instance = new(Paypal)
	default:
		p.instance = new(Stripe)
	}

	return p
}

// get instance of the registered payment gateway
func (p *Payment) GetInstance() PaymentMethod {
	return p.instance
}

////////////////////////////////////////////
// Stripe gateway implementation
////////////////////////////////////////////
type Stripe struct{}

// pay method implementation for the stripe gateway
func (s *Stripe) Pay(amount float32) string {

	// call gateway api and process billing

	return "Payment has been accepted by the stripe gateway"
}

////////////////////////////////////////////
// Paypal gateway implementation
////////////////////////////////////////////
type Paypal struct{}

// pay method implementation for the paypal gateway
func (s *Paypal) Pay(amount float32) string {

	// call gateway api and process billing

	return "Payment has been accepted by the paypal gateway"
}
