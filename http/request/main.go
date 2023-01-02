package main

import (
	"context"
	"net/http"
)

func main() {
	c := http.Request{}
	c.WithContext(context.Background())
}
