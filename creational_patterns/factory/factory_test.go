package factory

import "testing"

func TestPayment_Pay(t *testing.T) {
	payment := Payment{}

	// using stripe as payment gateway
	payment.SetPaymentGateway("stripe")

	if payment.GetInstance().Pay(200) != "Payment has been accepted by the stripe gateway" {
		t.Error("Payment has been accepted by the paypal gateway'")
	}

	// using paypal as payment gateway
	payment.SetPaymentGateway("paypal")

	if payment.GetInstance().Pay(200) != "Payment has been accepted by the paypal gateway" {
		t.Error("Message should be 'Payment has been accepted by the paypal gateway'")
	}
}
