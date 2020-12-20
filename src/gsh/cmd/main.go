package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {

		wd, err := os.Getwd()

		fmt.Print(wd, "> ")

		input, err := reader.ReadString('\n')

		// fmt.Println("DEBUG_INPUT 1: ", input)

		input = strings.TrimSuffix(input, "\n")

		// fmt.Println("DEBUG_INPUT 2: ", input)

		args := strings.Split(input, " ")

		// fmt.Println(args)
		switch args[0] {
		case "cd":
			if len(args) < 2 {
				err = os.Chdir("/home/root/")
				if err != nil {
					fmt.Println(err)
				}
			} else {
				err = os.Chdir(args[1])
				if err != nil {
					fmt.Println(err)
				}
			}

		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			err = cmd.Run()
			if err != nil {
				fmt.Println(err)
			}

		}

	} // never stop!
}
