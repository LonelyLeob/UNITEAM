package main

import "courses/internal"

func main() {
	cr, cd := internal.SetConfigs()
	srv := internal.NewServer(cd, cr)
	srv.Start()
}
