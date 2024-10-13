package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/paymentintent"
	"github.com/tomek-skrond/stripe-tests/lib"
	"github.com/tomek-skrond/stripe-tests/lib/httpwrap"
)

type PaymentIntentJSON struct {
	Amount int64 `json:"amount"`
}

func paymentIntents(w http.ResponseWriter, r *http.Request) error {
	stripe.Key = os.Getenv("SK_TEST_KEY")

	var jsonPayload PaymentIntentJSON

	if err := json.NewDecoder(r.Body).Decode(&jsonPayload); err != nil {
		return lib.InternalServerErrorResponse
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(jsonPayload.Amount),
		Currency: stripe.String(string(stripe.CurrencyPLN)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	result, err := paymentintent.New(params)
	if err != nil {
		return lib.InternalServerErrorResponse
	}

	// log.Printf("%v\n", chargeResponse)
	responseContent := httpwrap.NewJSONResponse(http.StatusOK, "successfully created intent", result)
	return httpwrap.WriteJSON(w, responseContent)

}

func hello(w http.ResponseWriter, r *http.Request) error {
	resp := map[string]string{
		"message": "hello",
	}
	responseContent := httpwrap.NewJSONResponse(http.StatusOK, "hello success", resp)
	return httpwrap.WriteJSON(w, responseContent)
}
