package main

import (
	"flag"
	"fmt"
	"os"
)

type config struct {
	port int
}

type application struct {
	config config
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Number of port to start server")

	flag.Parse()

	app := &application{
		config: cfg,
	}

	err := app.serve()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
