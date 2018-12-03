package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	//   /usr/bin/ruby /home/way/ruby/analyzer.rb
	cmd := exec.Command("bash", "/home/way/ruby/test.bash")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("failed.")
	}
	fmt.Printf("%q\n", out.String())
}
