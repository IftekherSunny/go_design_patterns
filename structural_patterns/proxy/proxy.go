package proxy

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
// Stripe gateway proxy implementation
////////////////////////////////////////////
type StripeProxy struct {
	stripe     Payment
	cachedData string
	cached     bool
}

// accepting payment by the stripe gateway proxy
func (sp *StripeProxy) Pay() string {
	if sp.stripe == nil {
		sp.stripe = new(Stripe)
	}

	if sp.cached == false {
		sp.cachedData = sp.stripe.Pay()
		sp.cached = true
	}

	return sp.cachedData
}
