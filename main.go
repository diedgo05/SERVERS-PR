package main

import (
	"server/principal"
	"server/replica"
)

func main() {
	go principal.Run()
	go replica.Run()

	select {}
}