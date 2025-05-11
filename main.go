package main

import (
	"log"
	"os/exec"
)

func main() {
	e := exec.Command("ls", "-l")
	b, err1 := e.Output()
	if err1 != nil {
		log.Println(err1)
	}
	log.Println(string(b))
}
