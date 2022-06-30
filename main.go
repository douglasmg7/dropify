package main

import (
	"fmt"
	"os"
)

var dropify_client_id string
var dropify_client_secret string

func main() {
	// App id.
	dropify_client_id = os.Getenv("DROPIFY_CLIENT_ID")
	if dropify_client_id == "" {
		panic("DROPIFY_CLIENT_ID env not defined.")
	}
	// Secret key.
	dropify_client_secret = os.Getenv("DROPIFY_CLIENT_SECRET")
	if dropify_client_secret == "" {
		panic("DROPIFY_CLIENT_SECRET env not defined.")
	}

	fmt.Println("dropify_client_id: ", dropify_client_id)
}
