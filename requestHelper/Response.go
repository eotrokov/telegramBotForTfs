package requestHelper

// Define Response struct
type Response struct {
	RawBody  []byte
	Body     Body
	Headers  map[string][]string
	Status   int
	Protocol string
}