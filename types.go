package main

type EndpointResponse struct {
	StatusCode int    // StatusCode is the status code of the response
	RawData    any    // RawData is the raw data of the response
	Url        string // Url is the full url of the request
}

type ValidationItem struct {
	Response EndpointResponse
}
