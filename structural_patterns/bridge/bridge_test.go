package bridge

import "testing"

func TestStripe_Pay(t *testing.T) {
	stripe := new(Stripe)

	if stripe.Pay() != "Payment has been accepted by the stripe gateway" {
		t.Error("Payment message should be 'Payment has been accepted by the stripe gateway'")
	}
}

func TestPaypal_Pay(t *testing.T) {
	paypal := new(Paypal)

	if paypal.Pay() != "Payment has been accepted by the paypal gateway" {
		t.Error("Payment message should be 'Payment has been accepted by the paypal gateway'")
	}
}

func TestShoppingCart_Checkout(t *testing.T) {

	// accepting payment through stripe
	stripe := new(Stripe)
	shoppingCart := ShoppingCart{
		Payment: stripe,
	}

	// assert shopping Cart checkout message with stripe return value
	if shoppingCart.Checkout("stripe") != stripe.Pay() {
		t.Error("Payment message should be 'Payment has been accepted by the stripe gateway'")
	}

	// accepting payment through paypal
	paypal := new(Paypal)
	shoppingCart = ShoppingCart{
		Payment: paypal,
	}

	// assert shopping Cart checkout message with paypal return value
	if shoppingCart.Checkout("paypal") != paypal.Pay() {
		t.Error("Payment message should be 'Payment has been accepted by the paypal gateway'")
	}
}
