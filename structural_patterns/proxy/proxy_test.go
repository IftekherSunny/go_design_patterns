package proxy

import "testing"

func TestStripe_Pay(t *testing.T) {
	stripe := new(Stripe)

	if stripe.Pay() != "Payment has been accepted by the stripe gateway" {
		t.Error("Payment message should be 'Payment has been accepted by the stripe gateway'")
	}
}

func TestStripeProxy_Pay(t *testing.T) {
	stripeProxy := new(StripeProxy)

	if stripeProxy.cached == true {
		t.Error("Cached field of the stripe gateway proxy should be false")
	}

	if stripeProxy.Pay() != "Payment has been accepted by the stripe gateway" {
		t.Error("Payment method return should be 'Payment has been accepted by the stripe gateway'")
	}

	if stripeProxy.cached == false {
		t.Error("Cached field of the stripe gateway proxy should be true")
	}
}
