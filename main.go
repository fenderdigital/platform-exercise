package main

import (
	"github.com/mrsmuneton/platform-test/src/server"
	"github.com/mrsmuneton/platform-test/src/session"
)

func main() {
	session.Session()
	server.Serve()
}
