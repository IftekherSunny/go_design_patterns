package decorator

////////////////////////////////////////////
// Payment method contract
////////////////////////////////////////////
type PaymentMethod interface {
	Pay(amount float32) string
}

////////////////////////////////////////////
// Paypal gateway decorator implementation
////////////////////////////////////////////
type PaypalDecorator struct{}

// pay method implementation for the paypal gateway
func (self *PaypalDecorator) Pay(amount float32) string {

	// call gateway api and process billing

	return "Payment has been accepted by the paypal gateway decorator"
}

////////////////////////////////////////////
// Paypal gateway implementation
////////////////////////////////////////////
type Paypal struct {
	PaypalDecorator PaymentMethod
}

// pay method implementation for the paypal gateway
func (p *Paypal) Pay(amount float32) string {

	if p.PaypalDecorator != nil {
		return p.PaypalDecorator.Pay(amount)
	}

	// call gateway api and process billing

	return "Payment has been accepted by the paypal gateway"
}
