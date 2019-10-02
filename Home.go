package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	out, err := exec.Command("mount").Output()
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Index(line, "on / type") != -1 {
			line = strings.Replace(line, " on / type ntfs (binary,noacl,auto)", "", -1)
			fmt.Println(line)
		}
	}
}
