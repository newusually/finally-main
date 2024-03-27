package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("python", "savecsv.py", "15m")
	res, _ := cmd.Output()
	fmt.Println(string(res))
}
