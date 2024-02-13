package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, res any) {
	decoder := json.NewDecoder(r.Body).Decode(res)
	PanicIfError(decoder)
}

func WriteToResponseBody(w http.ResponseWriter, res any) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w).Encode(res)
	PanicIfError(encoder)
}
