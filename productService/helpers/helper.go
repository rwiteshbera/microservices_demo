package helpers

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(res http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			res.Header()[key] = value
		}
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	_, err = res.Write(out)
	if err != nil {
		return err
	}

	return nil
}

// Read JSON Data
func ReadJSON(res http.ResponseWriter, req *http.Request, data interface{}) error {
	maxBytes := 1048576 // bytes = 1 MB // Maximum size of the body

	// Limit the size of the body
	req.Body = http.MaxBytesReader(res, req.Body, int64(maxBytes))

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(data)
	if err != nil {
		return err
	}
	return nil
}
