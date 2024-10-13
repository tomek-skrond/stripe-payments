package lib

import (
	"net/http"

	httpwrap "github.com/tomek-skrond/stripe-tests/lib/httpwrap"
)

var InternalServerErrorResponse = httpwrap.NewJSONResponse(http.StatusInternalServerError, "internal server error", nil)
