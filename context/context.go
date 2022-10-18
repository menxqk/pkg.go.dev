package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Ok         bool
	Msg        string
	StatusCode int
	Body       string
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	r, err := httpGet(ctx, "http://golang.org")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", r)
}

func httpGet(ctx context.Context, url string) (*Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	c := make(chan *Response)
	go func(r *http.Request, c chan *Response) {
		response := &Response{}

		result, err := http.DefaultClient.Do(r)
		if err != nil {
			response.Ok = false
			response.Msg = err.Error()
			response.StatusCode = http.StatusInternalServerError
			c <- response
			return
		}

		bytes, err := ioutil.ReadAll(result.Body)
		if err != nil {
			response.Ok = false
			response.Msg = err.Error()
			response.StatusCode = http.StatusInternalServerError
			c <- response
			return
		}
		defer result.Body.Close()

		response.Ok = true
		response.StatusCode = result.StatusCode
		response.Body = string(bytes[0:20])
		c <- response
	}(req, c)

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case response := <-c:
			if !response.Ok {
				return response, fmt.Errorf(response.Msg)
			}
			return response, nil
		}
	}
}
