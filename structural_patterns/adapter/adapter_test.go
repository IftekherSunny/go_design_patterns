package adapter

import "testing"

func TestStripe_Payment(t *testing.T) {
	stripe := new(Stripe)

	if stripe.Payment() != "Payment has been accepted by the stripe gateway" {
		t.Error("Payment message should be 'Payment has been accepted by the stripe gateway'")
	}
}

func TestStripeAdapter_Pay(t *testing.T) {
	stripeAdapter := new(StripeAdapter)

	if stripeAdapter.Pay() != "Payment has been accepted by the stripe gateway" {
		t.Error("Payment message should be 'Payment has been accepted by the stripe gateway'")
	}
}

func TestPaypal_AcceptPayment(t *testing.T) {
	paypal := new(Paypal)

	if paypal.AcceptPayment() != "Payment has been accepted by the paypal gateway" {
		t.Error("Payment message should be 'Payment has been accepted by the paypal gateway'")
	}
}

func TestPaypalAdapter_Pay(t *testing.T) {
	paypalAdapter := new(PaypalAdapter)

	if paypalAdapter.Pay() != "Payment has been accepted by the paypal gateway" {
		t.Error("Payment message should be 'Payment has been accepted by the paypal gateway'")
	}
}

func TestShoppingCart_Checkout(t *testing.T) {

	// accepting payment through stripe adapter
	stripeAdapter := new(StripeAdapter)
	shoppingCart := ShoppingCart{
		Payment: stripeAdapter,
	}

	// assert shopping card checkout message with stripe adapter return value
	if shoppingCart.Checkout("stripe") != stripeAdapter.Pay() {
		t.Error("Payment message should be 'Payment has been accepted by the stripe gateway'")
	}

	// accepting payment through paypal adapter
	paypalAdapter := new(PaypalAdapter)
	shoppingCart = ShoppingCart{
		Payment: paypalAdapter,
	}

	// assert shopping card checkout message with paypal adapter return value
	if shoppingCart.Checkout("paypal") != paypalAdapter.Pay() {
		t.Error("Payment message should be 'Payment has been accepted by the paypal gateway'")
	}
}
