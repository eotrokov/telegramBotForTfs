package requestHelper

import "io"

// Define Request struct
type Request struct {
	Url     string
	Headers map[string][]string
	Body    io.Reader
	Type string
}
