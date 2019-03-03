package requestHelper

import (
	. "../config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func BasePost(req Request) (Response, error) {
	// Build the URL
	var url= req.Url
	// Create an HTTP client
	c := &http.Client{}
	// Create an HTTP requestHelper
	r, err := http.NewRequest("POST", url, req.Body)
	if err != nil {
		return Response{}, err
	}
	// Add any defined headers
	if req.Headers != nil {
		r.Header = http.Header(req.Headers)
	}
	// Send the requestHelper
	res, err := c.Do(r)
	// Check for error
	if err != nil {
		return Response{}, err
	}
	// Make sure to close after reading
	defer res.Body.Close()
	// Limit response body to 1mb
	lr := &io.LimitedReader{res.Body, 1000000}
	// Read all the response body
	rb, err := ioutil.ReadAll(lr)
	// Check for error
	if err != nil {
		return Response{}, err
	}
	// Build the output
	responseOutput := Response{
		Body: Body{
			Bytes:  rb,
			String: string(rb),
		},
		Headers:  res.Header,
		Status:   res.StatusCode,
		Protocol: res.Proto,
	}
	// Send it along
	return responseOutput, nil
}

func Post(uri string, body interface{}, headers map[string][]string, response interface{}) {
	configuration := GetConfig()
	headers["Content-Type"] = []string{"application/json"}
	jsonBody, _ := json.Marshal(body)
	req := Request{
		Url:     configuration.TfsUrl + uri,
		Headers: headers,
		Body:    bytes.NewReader(jsonBody),
	}
	res, err := BasePost(req)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(res.Body.Bytes, &response)
	if err != nil {
		fmt.Println(err)
	}
}