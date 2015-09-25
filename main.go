package main

import (
	"github.com/calebthompson/pgf/web"
	"github.com/guregu/kami"
)

func main() {
	kami.Get("/", web.Form)
	kami.Post("/keyring", web.KeychainHandler)
	kami.Serve()
}
