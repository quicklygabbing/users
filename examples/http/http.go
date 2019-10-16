package main

import (
	"log"

	"github.com/quicklygabbing/users/cmd"
)

func main() {
	address := `:80`
	err := cmd.Starter(&address)
	if err != nil {
		log.Fatalln(err)
	}
}
