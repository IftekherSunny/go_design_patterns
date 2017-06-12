package decorator

import "testing"

func TestPaypalDecorator_Pay(t *testing.T) {
	t.Run("Testing paypal decorator pay method", func(t *testing.T) {
		paypalDecorator := new(PaypalDecorator)

		if paypalDecorator.Pay(100) != "Payment has been accepted by the paypal gateway decorator" {
			t.Error("Payment message should be 'Payment has been accepted by the paypal gateway decorator'")
		}
	})
}

func TestPaypal_Pay(t *testing.T) {
	t.Run("Testing paypal pay method without decorator", func(t *testing.T) {
		paypal := new(Paypal)

		if paypal.Pay(100) != "Payment has been accepted by the paypal gateway" {
			t.Error("Payment message should be 'Payment has been accepted by the paypal gateway'")
		}
	})

	t.Run("Testing paypal pay method with decorator", func(t *testing.T) {
		paypal := Paypal{
			PaypalDecorator: new(PaypalDecorator),
		}

		if paypal.Pay(100) != "Payment has been accepted by the paypal gateway decorator" {
			t.Error("Payment message should be 'Payment has been accepted by the paypal gateway decorator'")
		}
	})
}
