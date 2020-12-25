package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {

	err := syscall.Mount(os.Args[1], os.Args[2], os.Args[3], 0, "")
	if err != nil {
		fmt.Println(err)
	}

}
