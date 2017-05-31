package prototype

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
	SetApiKey(apiKey string)
	GetApiKey() string
}

////////////////////////////////////////////
// Stripe gateway implementation
////////////////////////////////////////////
type Stripe struct {
	apiKey string
}

// set stripe api key
func (s *Stripe) SetApiKey(apiKey string) {
	s.apiKey = apiKey
}

// get stripe api key
func (s *Stripe) GetApiKey() string {
	return s.apiKey
}

////////////////////////////////////////////
// Paypal gateway implementation
////////////////////////////////////////////
type Paypal struct {
	apiKey string
}

// set paypal api key
func (p *Paypal) SetApiKey(apiKey string) {
	p.apiKey = apiKey
}

// get paypal api key
func (p *Paypal) GetApiKey() string {
	return p.apiKey
}

////////////////////////////////////////////
// Payment gateway instance
////////////////////////////////////////////
var stripe Stripe = Stripe{apiKey: "123"}
var paypal Paypal = Paypal{apiKey: "123"}

////////////////////////////////////////////
// Payment prototype
////////////////////////////////////////////
type Payment struct{}

// get clone copy for the given payment gateway type's
func (p *Payment) GetClone(gatewayType string) PaymentMethod {
	switch gatewayType {
	case STRIPE:
		stripeInstance := stripe
		return &stripeInstance
	case PAYPAL:
		paypalInstance := paypal
		return &paypalInstance
	default:
		stripeInstance := stripe
		return &stripeInstance
	}
}
