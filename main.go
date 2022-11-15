package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("hi")
	out, err := exec.Command("python3.8", "--version").Output()
	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Println("SUCCESS: ", string(out))
	}
}
