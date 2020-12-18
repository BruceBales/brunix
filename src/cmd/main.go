package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"time"
)

var welcome = `
 /$$$$$$$  /$$$$$$$  /$$   /$$ /$$   /$$ /$$$$$$ /$$   /$$
| $$__  $$| $$__  $$| $$  | $$| $$$ | $$|_  $$_/| $$  / $$
| $$  \ $$| $$  \ $$| $$  | $$| $$$$| $$  | $$  |  $$/ $$/
| $$$$$$$ | $$$$$$$/| $$  | $$| $$ $$ $$  | $$   \  $$$$/ 
| $$__  $$| $$__  $$| $$  | $$| $$  $$$$  | $$    >$$  $$ 
| $$  \ $$| $$  \ $$| $$  | $$| $$\  $$$  | $$   /$$/\  $$
| $$$$$$$/| $$  | $$|  $$$$$$/| $$ \  $$ /$$$$$$| $$  \ $$
|_______/ |__/  |__/ \______/ |__/  \__/|______/|__/  |__/
                                                          
                                                          
`

func main() {

	fmt.Println("Setting path to /bin")
	os.Setenv("PATH", "/bin")
	fmt.Println("Creating /dev/null")
	os.Create("/dev/null")
	os.Chmod("/dev/null", 666)

	time.Sleep(time.Second * 1)
	fmt.Println("--- BRUNIX V0.0.0.0.0.0.1 ---")
	fmt.Println(welcome)
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
			err = cmd.Run()
			if err != nil {
				fmt.Println(err)
			}

		}

	} // never stop!
}
