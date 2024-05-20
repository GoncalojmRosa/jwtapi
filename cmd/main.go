package main

import (
	"goncalojrmosa/jwtapi/cmd/api"
	"log"
)

func main() {
	sv := api.NewAPIServer(":8080", nil)
	if err := sv.Run(); err != nil{
		log.Fatal(err)
	}
}