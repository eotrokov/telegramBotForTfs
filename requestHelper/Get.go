package requestHelper

import (
	"encoding/json"
	"fmt"
)


func Get(uri string, headers map[string][]string, response interface{}) {
	req := Request{
		Url:     uri,
		Headers: headers,
		Type:    "GET",
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
