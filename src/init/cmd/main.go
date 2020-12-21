package main

import (
	"fmt"
	"os"
	"os/exec"

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

	fmt.Println("Creating /dev/null")
	os.Create("/dev/null")
	os.Chmod("/dev/null", 666)

	os.Setenv("LD_LIBRARY_PATH", "/lib:/lib64")

	fmt.Println("Setting user path")
	os.Setenv("PATH", "/bin:/sbin:/usr/bin:/usr/sbin:/go/bin/")

	fmt.Println("Creating /proc/")
	os.Mkdir("/proc", 777)
	err := exec.Command("mount", "-t", "proc", "p", "/proc").Run()
	if err != nil {
		fmt.Println("Error: Could not mount /proc")
	}

	cmd := exec.Command("mount", "-t", "devtmpfs", "devtmpfs", "/dev")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second * 1)
	fmt.Println("--- BRUNIX V0.1 ---")
	fmt.Println(welcome)

	cmd = exec.Command("gsh")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	for {

	} // never stop!
}
