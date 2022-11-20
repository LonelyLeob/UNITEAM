package main

import (
	"github.com/L0nelyleob/messanger/internal"
)

func main() {
	cr, cm := internal.SetConfigs()
	srv := internal.Initialize(cr, cm)
	srv.Serve()
}
