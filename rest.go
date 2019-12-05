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
		payload []byte
	)
	switch strings.ToUpper(method) {
	case http.MethodGet:
		request, err = http.NewRequest(http.MethodGet, path, nil)
	case http.MethodPost:
		if data != nil {
			payload, err = json.Marshal(data)
			fmt.Println(fmt.Sprintf("The payload: %v", bytes.NewBuffer(payload)))
		}
		fmt.Println(fmt.Sprintf("The path: %v", path))
		if err == nil {
			request, err = http.NewRequest(http.MethodGet, path, nil)
		}
	}
	if err == nil && request != nil {
		request.Header.Set("content-type", "application/json")
		request.Header.Set("accept", "application/json")
		client := http.Client{Timeout: 60 * time.Second}
		fmt.Println(fmt.Sprintf("The request body: %v", request.Body))
		fmt.Println(fmt.Sprintf("The request headers: %v", request.Header))
		if response, err = client.Do(request); err == nil {
			if response == nil {
				err = errors.New("freshsales did not respond to request")
			} else {
				if response.StatusCode != 200 {
					fmt.Println(fmt.Sprintf("The response: %v", response))
					err = errors.New(fmt.Sprintf("freshsales responded with the status code of %d",
						response.StatusCode))
				}
			}
		}
	}
 	return
}