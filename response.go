package main

// response format
type (
	Response struct {
		Response  interface{} `json:"response"`
		Exception interface{} `json:"exception"`
		Error     error       `json:"error"`
	}
)
