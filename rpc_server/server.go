package main

import (
	`github.com/valyala/gorpc`
	`log`
)

func main()  {
	s := &gorpc.Server{
		// Accept clients on this TCP address.
		Addr: ":12346",

		// Echo handler - just return back the message we received from the client
		Handler: func(clientAddr string, request interface{}) interface{} {
			log.Printf("Obtained request %+v from the client %s\n", request, clientAddr)
			return request
		},
	}
	if err := s.Serve(); err != nil {
		log.Fatalf("Cannot start rpc_server server: %s", err)
	}
}