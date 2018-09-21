package main

import (
	`github.com/valyala/gorpc`
	`log`
)

func main()  {
	c := &gorpc.Client{
		// TCP address of the server.
		Addr: "127.0.0.1:12346",
	}
	c.Start()

	// All client methods issuing RPCs are thread-safe and goroutine-safe,
	// i.e. it is safe to call them from multiple concurrently running goroutines.
	resp, err := c.Call("foobar")
	if err != nil {
		log.Fatalf("Error when sending request to server: %s", err)
	}
	if resp.(string) != "foobar" {
		log.Fatalf("Unexpected response from the server: %+v", resp)
	}
}
