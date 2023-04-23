package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancela := context.WithTimeout(ctx, time.Second)
	defer cancela()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}
	println(string(body))
}
