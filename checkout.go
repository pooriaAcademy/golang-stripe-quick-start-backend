package main

import (
	"github.com/stripe/stripe-go/v72/customer"
	"log"
	"net/http"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

var PriceId = ""

func checkout(email string) (*stripe.CheckoutSession, error) {
	var discounts []*stripe.CheckoutSessionDiscountParams

	discounts = []*stripe.CheckoutSessionDiscountParams{
		&stripe.CheckoutSessionDiscountParams{
			Coupon: stripe.String("FMARC"),
		},
	}

	customerParams := &stripe.CustomerParams{
		Email: stripe.String(email),
	}
	customerParams.AddMetadata("FinalEmail", email)
	newCustomer, err := customer.New(customerParams)

	if err != nil {
		return nil, err
	}

	meta := map[string]string{
		"FinalEmail" : email,
	}

	log.Println("Creating meta for user: ", meta)

	params := &stripe.CheckoutSessionParams{
		Customer: &newCustomer.ID,
		SuccessURL: stripe.String("https://scalperfighter.com/success.html"),
		CancelURL: stripe.String("https://scalperfighter.com/"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Discounts: discounts,
		Mode: stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				Price:    stripe.String(PriceId),
				// For metered billing, do not pass quantity
				Quantity: stripe.Int64(1),
			},
		},
		SubscriptionData: &stripe.CheckoutSessionSubscriptionDataParams{
			TrialPeriodDays: stripe.Int64(7),
			Metadata: meta,
		},
	}
	// TODO in the future look if using this meta data is better
	//params.AddMetadata(domain.FinalEmailTag, email)
	return session.New(params)
}

func CheckoutCreator(w http.ResponseWriter, req * http.Request){

}












