package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonValue := bytes.NewBuffer([]byte(`{"name":"jorge"}`))
	resp, err := c.Post("https://www.google.com", "application/json", jsonValue)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
