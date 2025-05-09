package main

import (
	"HTTP_SERVER/server"
)

func main() {
	srv := server.New(":8080")

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
