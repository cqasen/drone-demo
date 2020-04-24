package main

import (
	"github.com/cqasen/drone-demo/http/router"
	"github.com/ebar-go/ego/http"
	"github.com/ebar-go/ego/utils"
)

func main() {
	server := http.NewServer()
	router.InitRouter(server.Router)
	utils.FatalError("StartServer", server.Start())
}
