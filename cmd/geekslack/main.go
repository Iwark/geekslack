package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Iwark/geekslack"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

// Request is the data from the client
type Request struct {
	Body string `json:"body"`
}

// Response is the data return to the client
type Response struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

func handler(c context.Context, req geekslack.Request) (resp Response, err error) {
	resp.Headers = map[string]string{
		"Content-Type": "application/json",
	}

	text, err := geekslack.Handle(&req)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		bytes, _ := json.Marshal(map[string]string{"code": "E2", "error": err.Error()})
		resp.Body = string(bytes)
		return
	}
	resp.StatusCode = http.StatusOK
	resp.Body = text
	return
}
