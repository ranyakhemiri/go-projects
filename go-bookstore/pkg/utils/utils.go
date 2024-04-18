package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody parses the body of a request and stores it in the given interface
// x is an empty interface type interface{} which allows any type to be passed in.
func ParseBody(r *http.Request, x interface{}) {
	// Read the body of the request
	if body, err := io.ReadAll(r.Body); err == nil {
		// convert body (which is a byte slice) into JSON format and then unmarshalling it into x
		// if there is an error (err != nil) during the JSON parsing, the function does nothing further and returns
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
