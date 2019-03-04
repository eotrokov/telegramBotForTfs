package requestHelper

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func Post(uri string, body interface{}, headers map[string][]string, response interface{}) {
	headers["Content-Type"] = []string{"application/json"}
	jsonBody, _ := json.Marshal(body)
	req := Request{
		Url:     uri,
		Headers: headers,
		Body:    bytes.NewReader(jsonBody),
		Type:    "POST",
	}
	res, err := BaseDo(req)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(res.Body.Bytes, &response)
	if err != nil {
		fmt.Println(err)
	}
}