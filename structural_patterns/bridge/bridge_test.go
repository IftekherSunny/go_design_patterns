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

func TestShoppingCard_Checkout(t *testing.T) {

	// accepting payment through stripe
	stripe := new(Stripe)
	shoppingCard := ShoppingCard{
		Payment: stripe,
	}

	// assert shopping card checkout message with stripe return value
	if shoppingCard.Checkout("stripe") != stripe.Pay() {
		t.Error("Payment message should be 'Payment has been accepted by the stripe gateway'")
	}

	// accepting payment through paypal
	paypal := new(Paypal)
	shoppingCard = ShoppingCard{
		Payment: paypal,
	}

	// assert shopping card checkout message with paypal return value
	if shoppingCard.Checkout("paypal") != paypal.Pay() {
		t.Error("Payment message should be 'Payment has been accepted by the paypal gateway'")
	}
}
