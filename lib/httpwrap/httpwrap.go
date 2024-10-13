package httpwrap

import (
	"encoding/json"
	"net/http"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type JSONResponse struct {
	ResponseContent any    `json:"response"`
	HTTPStatus      int    `json:"http_status"`
	Message         string `json:"message"`
}

type APIError struct {
	JSONResponse
	Err string
}

func (e APIError) Error() string {
	return e.Err
}

// WriteJSON writes the JSON response to the http.ResponseWriter
func WriteJSON(w http.ResponseWriter, content APIError) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(content.JSONResponse.HTTPStatus)
	return json.NewEncoder(w).Encode(content.JSONResponse)
}

func NewJSONResponse(status int, msg string, content any) APIError {
	response := JSONResponse{
		ResponseContent: content,
		HTTPStatus:      status,
		Message:         msg,
	}
	return APIError{
		Err:          msg,
		JSONResponse: response,
	}
}

func MakeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if e, ok := err.(APIError); ok {
				WriteJSON(w, e)
			}
			resp := NewJSONResponse(http.StatusInternalServerError, "internal server error", nil)
			WriteJSON(w, resp)
		}
	}
}
