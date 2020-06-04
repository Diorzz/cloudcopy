package main

import (
	"server/action"
	"server/server"
)

func main() {
	// server start
	s := server.NewServer("127.0.0.1:8888")

	// user regist and login
	s.RegisterRouter("GET", "/", action.HandleHome)

	// start server
	s.Start()
}
