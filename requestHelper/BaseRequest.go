package requestHelper

import (
	. "../config"
	"io"
	"io/ioutil"
	"net/http"
)

func BaseDo(req Request) (Response, error) {
	configuration := GetConfig()
	var url = configuration.TfsUrl + req.Url
	c := &http.Client{}
	r, err := http.NewRequest(req.Type, url, req.Body)
	if err != nil {
		return Response{}, err
	}
	if req.Headers != nil {
		r.Header = http.Header(req.Headers)
	}
	res, err := c.Do(r)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()
	lr := &io.LimitedReader{res.Body, 1000000}
	rb, err := ioutil.ReadAll(lr)
	if err != nil {
		return Response{}, err
	}
	responseOutput := Response{
		Body: Body{
			Bytes:  rb,
			String: string(rb),
		},
		Headers:  res.Header,
		Status:   res.StatusCode,
		Protocol: res.Proto,
	}
	return responseOutput, nil
}

