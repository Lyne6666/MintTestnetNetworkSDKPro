// cmd/minttestnetnetworksdkpro/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"minttestnetnetworksdkpro/internal/minttestnetnetworksdkpro"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := minttestnetnetworksdkpro.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
