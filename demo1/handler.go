package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
	"golang.org/x/net/context/ctxhttp"
	"net/http"
)
func Handler(ctx context.Context) (int, error) {
	fmt.Print(ctx)
	httpClient := xray.Client(http.DefaultClient)
	resp, err := ctxhttp.Get(ctx, httpClient, "http://example.com/")
	return resp.StatusCode, err
}

func main() {
	lambda.Start(Handler)
}