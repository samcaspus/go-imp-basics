package main

import (
	"github.com/samcaspus/go-imp-basics/constants/server_constants"
	"github.com/samcaspus/go-imp-basics/router/default_router"
	"github.com/samcaspus/go-imp-basics/server"
	"os"
)

func main() {
	os.Setenv("GIN_MODE", "release")
	engine := server.InitServer()
	default_router.AttachRoutes(engine)
	engine.Run(server_constants.Port)
}
