package main

import (
	"spin-it/api"
)

func main() {
	a := api.API{}
	a.Initialise()

	a.Run(":8080")
}
