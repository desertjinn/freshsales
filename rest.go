package freshsales

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// make a request to an endpoint with the desired payload
func request(method string, path string, data map[string]interface{}) (err error) {
	var (
		request *http.Request
		response *http.Response
	)
	switch strings.ToUpper(method) {
	case http.MethodGet:
		request, err = http.NewRequest(http.MethodGet, path, nil)
	case http.MethodPost:
		var payload []byte
		if data != nil {
			payload, err = json.Marshal(data)
		}
		if err == nil {
			request, err = http.NewRequest(http.MethodGet, path, bytes.NewBuffer(payload))
		}
	}
	if err == nil && request != nil {
		if data[TokenKey] != nil {
			request.Header.Set("Authorization", fmt.Sprintf("Token token=%v", data[TokenKey]))	
		}
		request.Header.Set("content-type", "application/json")
		request.Header.Set("accept", "application/json")
		client := http.Client{Timeout: 60 * time.Second}
		if response, err = client.Do(request); err == nil {
			if response == nil {
				err = errors.New("freshsales did not respond to request")
			} else {
				if response.StatusCode != 200 {
					err = errors.New(fmt.Sprintf("freshsales responded with the status code of %d",
						response.StatusCode))
				}
			}
		}
	}
 	return
}