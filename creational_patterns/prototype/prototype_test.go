package prototype

import "testing"

func TestPayment_GetClone(t *testing.T) {
	// payment prototype
	payment := Payment{}

	// test stripe gateway
	testStripeGateway(payment, t)

	// test paypal gateway
	testPaypalGateway(payment, t)
}

////////////////////////////////////////////
// Test stripe gateway
////////////////////////////////////////////
func testStripeGateway(payment Payment, t *testing.T) {

	// cloning stripe gateway instance
	stripeInstance1 := payment.GetClone("stripe")
	stripeInstance2 := payment.GetClone("stripe")

	// asset stripe gateway clone copies
	if stripeInstance1 == stripeInstance2 {
		t.Error("Stripe gateway instance can't be same instance")
	}

	// assert default api key
	if stripeInstance1.GetApiKey() != "123" {
		t.Error("Stripe api key should be 123")
	}

	if stripeInstance2.GetApiKey() != "123" {
		t.Error("Stripe api key should be 123")
	}

	// change api key
	stripeInstance1.SetApiKey("456")

	// assert stripe api key value after changing api key value
	if stripeInstance1.GetApiKey() == "123" {
		t.Error("Stripe gateway api key should be 456")
	}
}

////////////////////////////////////////////
// Test paypal gateway
////////////////////////////////////////////
func testPaypalGateway(payment Payment, t *testing.T) {

	// cloning paypal gateway instance
	paypalInstance1 := payment.GetClone("paypal")
	paypalInstance2 := payment.GetClone("paypal")

	// asset paypal gateway clone copies
	if paypalInstance1 == paypalInstance2 {
		t.Error("Paypal gateway instance can't be same instance")
	}

	// assert default api key
	if paypalInstance1.GetApiKey() != "123" {
		t.Error("Paypal api key should be 123")
	}

	if paypalInstance2.GetApiKey() != "123" {
		t.Error("Paypal api key should be 123")
	}

	// change api key
	paypalInstance1.SetApiKey("456")

	// assert paypal api key value after changing api key value
	if paypalInstance1.GetApiKey() == "123" {
		t.Error("Paypal gateway api key should be 456")
	}
}
