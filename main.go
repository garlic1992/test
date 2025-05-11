package main

import (
	"log"
	"os/exec"
)

func main() {
	e := exec.Command("ls", "-l")
	b, err := e.Output()
	if err != nil {
		log.Println(err)
	}
	log.Println(string(b))
}
