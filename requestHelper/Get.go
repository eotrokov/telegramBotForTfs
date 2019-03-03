package requestHelper

import (
	. "../config"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)


func BaseGet(req Request) (Response, error) {
	// Build the URL
	var url= req.Url
	// Create an HTTP client
	c := &http.Client{}
	// Create an HTTP requestHelper
	r, err := http.NewRequest("GET", url, nil)
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

func Get(uri string, headers map[string][]string, response interface{}) {
	configuration := GetConfig()
	req := Request{
		Url:     configuration.TfsUrl + uri,
		Headers: headers,
	}
	res, err := BaseGet(req)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(res.Body.Bytes, &response)
	if err != nil {
		fmt.Println(err)
	}
}
